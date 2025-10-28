package courseProcedure

import (
	"backend/generate/psql"
	"backend/helper/convert"
	"backend/type/payload"
	"context"

	"github.com/bsthun/gut"
)

func (r *Procedure) CourseListExplore(ctx context.Context, userId *uint64, name *string, limit *uint64, offset *uint64) ([]*payload.CourseExtended, *uint64, *gut.ErrorInstance) {
	// * query course list explore
	courseRows, err := r.database.P().CourseListExplore(ctx, &psql.CourseListExploreParams{
		UserId: userId,
		Name:   name,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, nil, gut.Err(false, "failed to list explore courses", err)
	}

	// * query count
	count, err := r.database.P().CourseCountExplore(ctx, &psql.CourseCountExploreParams{
		UserId: userId,
		Name:   name,
	})
	if err != nil {
		return nil, nil, gut.Err(false, "failed to count explore courses", err)
	}

	// * map course rows to extended courses
	items, er := gut.Iterate(courseRows, func(course psql.CourseListExploreRow) (*payload.CourseExtended, *gut.ErrorInstance) {
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