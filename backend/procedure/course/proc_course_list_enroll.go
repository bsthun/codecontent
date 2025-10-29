package courseProcedure

import (
	"backend/generate/psql"
	"backend/helper/convert"
	"backend/type/payload"
	"context"

	"github.com/bsthun/gut"
)

func (r *Procedure) CourseListEnroll(ctx context.Context, userId *uint64, name *string, limit *uint64, offset *uint64) ([]*payload.CourseExtended, *uint64, *gut.ErrorInstance) {
	// * query course list enroll
	courseRows, err := r.database.P().CourseListByEnroll(ctx, &psql.CourseListByEnrollParams{
		UserId: userId,
		Name:   name,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, nil, gut.Err(false, "failed to list enroll courses", err)
	}

	// * query count
	count, err := r.database.P().CourseCountByEnroll(ctx, &psql.CourseCountByEnrollParams{
		UserId: userId,
		Name:   name,
	})
	if err != nil {
		return nil, nil, gut.Err(false, "failed to count enroll courses", err)
	}

	// * map course rows to payload courses
	items := convert.Course.CourseListByEnrollRowsToPayload(courseRows)

	// * return
	return items, count, nil
}
