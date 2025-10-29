package courseProcedure

import (
	"backend/helper/convert"
	"backend/type/payload"
	"context"

	"github.com/bsthun/gut"
)

func (r *Procedure) CourseManageDelete(ctx context.Context, courseId *uint64) (*payload.Course, *gut.ErrorInstance) {
	// * query course delete
	course, err := r.database.P().CourseDelete(ctx, courseId)
	if err != nil {
		return nil, gut.Err(false, "failed to delete course", err)
	}

	// * map course to payload course
	result := convert.Course.CourseRowToPayload(course)

	// * return
	return result, nil
}
