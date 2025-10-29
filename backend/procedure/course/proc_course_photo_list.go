package courseProcedure

import (
	"backend/generate/psql"
	"backend/type/payload"
	"context"

	"github.com/bsthun/gut"
)

func (r *Procedure) CoursePhotoList(ctx context.Context, courseId *uint64, title *string, sort *string, order *string, limit *uint64, offset *uint64) ([]*payload.CoursePhoto, *uint64, *gut.ErrorInstance) {
	// * query: course photo count
	count, err := r.database.P().CoursePhotoCount(ctx, &psql.CoursePhotoCountParams{
		CourseId: courseId,
		Title:    title,
	})
	if err != nil {
		return nil, nil, gut.Err(false, "failed to count course photos", err)
	}

	// * query: course photo list
	coursePhotoRows, err := r.database.P().CoursePhotoList(ctx, &psql.CoursePhotoListParams{
		CourseId: courseId,
		Title:    title,
		Sort:     sort,
		Order:    order,
		Offset:   offset,
		Limit:    limit,
	})
	if err != nil {
		return nil, nil, gut.Err(false, "failed to list course photos", err)
	}

	// * map course photos to items
	coursePhotoItems, _ := gut.Iterate(coursePhotoRows, func(row psql.CoursePhotoListRow) (*payload.CoursePhoto, *gut.ErrorInstance) {
		return &payload.CoursePhoto{
			Id:          row.CoursePhoto.Id,
			CourseId:    row.CoursePhoto.CourseId,
			Title:       row.CoursePhoto.Title,
			Description: row.CoursePhoto.Description,
			PhotoUrl:    r.pathService.CoursePhotoMinioUrl(row.CoursePhoto.CourseId, row.CoursePhoto.Id),
			CreatedAt:   row.CoursePhoto.CreatedAt,
			UpdatedAt:   row.CoursePhoto.UpdatedAt,
		}, nil
	})

	// * return
	return coursePhotoItems, count, nil
}
