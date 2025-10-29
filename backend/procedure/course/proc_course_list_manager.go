package courseProcedure

import (
	"backend/generate/psql"
	"backend/helper/convert"
	"backend/type/payload"
	"context"

	"github.com/bsthun/gut"
)

func (r *Procedure) CourseListManager(ctx context.Context, userId *uint64, name *string, limit *uint64, offset *uint64) ([]*payload.CourseExtended, *uint64, *gut.ErrorInstance) {
	// * query course list by manager
	courseRows, err := r.database.P().CourseListByManager(ctx, &psql.CourseListByManagerParams{
		Userid: userId,
		Name:   name,
		Sort:   gut.Ptr("createdAt"),
		Order:  gut.Ptr("desc"),
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, nil, gut.Err(false, "failed to get courses by manager", err)
	}

	// * query course count by manager
	countRow, err := r.database.P().CourseCountByManager(ctx, &psql.CourseCountByManagerParams{
		Userid: userId,
		Name:   name,
	})
	if err != nil {
		return nil, nil, gut.Err(false, "failed to count courses by manager", err)
	}

	// * map courses to payload courses
	courses := convert.Course.CourseListByManagerRowsToPayload(courseRows)

	// * return
	return courses, countRow, nil
}
