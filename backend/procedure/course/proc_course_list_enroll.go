package courseProcedure

import (
	"backend/generate/psql"
	"backend/helper/convert"
	"backend/type/payload"
	"context"

	"github.com/bsthun/gut"
)

func (r *Procedure) CourseListByEnroll(ctx context.Context, userId *uint64, name *string, limit *uint64, offset *uint64) ([]*payload.CourseExtended, *uint64, *gut.ErrorInstance) {
	// * query course list by enroll
	courseRows, err := r.database.P().CourseListByEnroll(ctx, &psql.CourseListByEnrollParams{
		UserId: userId,
		Name:   name,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, nil, gut.Err(false, "failed to list courses by enroll", err)
	}

	// * query count
	count, err := r.database.P().CourseCountByEnroll(ctx, &psql.CourseCountByEnrollParams{
		UserId: userId,
		Name:   name,
	})
	if err != nil {
		return nil, nil, gut.Err(false, "failed to count courses by enroll", err)
	}

	// * map course rows to extended courses
	items, er := gut.Iterate(courseRows, func(course psql.CourseListByEnrollRow) (*payload.CourseExtended, *gut.ErrorInstance) {
		coursePayload := convert.Course.CourseRowToPayload(course.Course)
		return &payload.CourseExtended{
			Course:             *coursePayload,
			CourseManagerCount: course.CourseManagerCount,
			EnrollCount:        course.EnrollCount,
			CoursePhotoCount:   course.CoursePhotoCount,
		}, nil
	})
	if er != nil {
		return nil, nil, er
	}

	// * return
	return items, count, nil
}