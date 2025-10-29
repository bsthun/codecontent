package courseProcedure

import (
	"backend/generate/psql"
	"backend/type/payload"
	"context"

	"github.com/bsthun/gut"
)

func (r *Procedure) CoursePhotoList(ctx context.Context, courseId *uint64, limit *uint64, offset *uint64) ([]*payload.CoursePhoto, *uint64, *gut.ErrorInstance) {
	// * query: course photo count
	count, err := r.database.P().CoursePhotoCount(ctx, &psql.CoursePhotoCountParams{
		CourseId: courseId,
		Title:    nil,
	})
	if err != nil {
		return nil, nil, gut.Err(false, "failed to count course photos", err)
	}

	// * query: course photo list
	coursePhotoRows, err := r.database.P().CoursePhotoList(ctx, &psql.CoursePhotoListParams{
		CourseId: courseId,
		Limit:    limit,
		Offset:   offset,
	})
	if err != nil {
		return nil, nil, gut.Err(false, "failed to list course photos", err)
	}

	// * map course photos to items
	coursePhotoItems, _ := gut.Iterate(coursePhotoRows, func(row psql.CoursePhotoListRow) (*payload.CoursePhoto, *gut.ErrorInstance) {
		return &payload.CoursePhoto{
			Id:          row.Id,
			CourseId:    row.CourseId,
			Title:       row.Title,
			Description: row.Description,
			PhotoUrl:    r.pathService.CoursePhotoMinioUrl(row.CourseId, row.Id),
			CreatedAt:   row.CreatedAt,
			UpdatedAt:   row.UpdatedAt,
		}, nil
	})

	// * return
	return coursePhotoItems, count, nil
}
