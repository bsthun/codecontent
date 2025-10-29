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

	// * map course rows to payload courses
	items := convert.Course.CourseListExploreRowsToPayload(courseRows)

	// * return
	return items, count, nil
}
