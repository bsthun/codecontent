package courseProcedure

import (
	"backend/generate/psql"
	"backend/type/payload"
	"bytes"
	"context"
	"io"

	"github.com/bsthun/gut"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/qdrant/go-client/qdrant"
)

func (r *Procedure) CoursePhotoUpload(ctx context.Context, courseId *uint64, imageReader io.Reader) (*payload.CoursePhoto, *gut.ErrorInstance) {
	// * begin transaction
	tx, querier := r.database.Ptx(ctx, nil)
	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
		}
	}()

	// * create course photo
	coursePhoto, er := querier.CoursePhotoCreate(ctx, &psql.CoursePhotoCreateParams{
		CourseId:    courseId,
		Title:       gut.Ptr(""),
		Description: gut.Ptr(""),
	})
	if er != nil {
		_ = tx.Rollback()
		return nil, gut.Err(false, "failed to create course photo record", er)
	}

	// * create buffer for tee reader
	var buf bytes.Buffer
	tee := io.TeeReader(imageReader, &buf)

	// * upload to minio
	_, er = r.minio.PutObject(
		ctx,
		*r.config.MinioBucket,
		*r.pathService.CoursePhotoMinioPath(courseId, coursePhoto.Id),
		tee,
		-1,
		minio.PutObjectOptions{
			ContentType: "image/png",
		},
	)
	if er != nil {
		_ = tx.Rollback()
		return nil, gut.Err(false, "failed to upload to minio", er)
	}

	// * compute embedding
	embeddingResp, er := r.computeService.EmbedImage(&buf)
	if er != nil {
		_ = tx.Rollback()
		return nil, gut.Err(false, "failed to get embedding", er)
	}

	// * insert embedding to qdrant
	pointId := uuid.New().String()
	point := &qdrant.PointStruct{
		Id: &qdrant.PointId{
			PointIdOptions: &qdrant.PointId_Uuid{
				Uuid: pointId,
			},
		},
		Vectors: &qdrant.Vectors{
			VectorsOptions: &qdrant.Vectors_Vector{
				Vector: &qdrant.Vector{
					Data: embeddingResp.Embeddings,
				},
			},
		},
		Payload: map[string]*qdrant.Value{
			"courseId": {
				Kind: &qdrant.Value_IntegerValue{
					IntegerValue: int64(*courseId),
				},
			},
			"coursePhotoId": {
				Kind: &qdrant.Value_IntegerValue{
					IntegerValue: int64(*coursePhoto.Id),
				},
			},
			"type": {
				Kind: &qdrant.Value_StringValue{
					StringValue: "course_photo",
				},
			},
		},
	}

	_, er = r.qdrant.Upsert(ctx, &qdrant.UpsertPoints{
		CollectionName: *r.config.QdrantCollection,
		Points: []*qdrant.PointStruct{
			point,
		},
	})
	if er != nil {
		_ = tx.Rollback()
		return nil, gut.Err(false, "failed to upsert to qdrant", er)
	}

	// * construct photo url
	photoUrl := r.pathService.CoursePhotoMinioUrl(courseId, coursePhoto.Id)

	// * call agent service for photo description
	title, description, er := r.agentService.FunctionPhotoDescription(*photoUrl)
	if er != nil {
		_ = tx.Rollback()
		return nil, gut.Err(false, "failed to get photo description", er)
	}

	// * update course photo with generated title and description
	updatedPhoto, er := querier.CoursePhotoUpdate(ctx, &psql.CoursePhotoUpdateParams{
		Id:          coursePhoto.Id,
		Title:       title,
		Description: description,
	})
	if er != nil {
		_ = tx.Rollback()
		return nil, gut.Err(false, "failed to update course photo", er)
	}

	// * commit transaction
	if er := tx.Commit(); er != nil {
		return nil, gut.Err(false, "failed to commit transaction", er)
	}

	// * construct final photo object
	finalPhoto := &payload.CoursePhoto{
		Id:          updatedPhoto.Id,
		CourseId:    updatedPhoto.CourseId,
		Title:       updatedPhoto.Title,
		Description: updatedPhoto.Description,
		PhotoUrl:    photoUrl,
		CreatedAt:   updatedPhoto.CreatedAt,
		UpdatedAt:   updatedPhoto.UpdatedAt,
	}

	return finalPhoto, nil
}
