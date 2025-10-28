package courseProcedure

import (
	"backend/generate/psql"
	"backend/helper/convert"
	"backend/type/payload"
	"context"

	"github.com/bsthun/gut"
)

func (r *Procedure) CourseEdit(ctx context.Context, courseId *uint64, name *string, description *string, promptInstruction *string) (*payload.Course, *gut.ErrorInstance) {
	// * query course update
	course, err := r.database.P().CourseUpdate(ctx, &psql.CourseUpdateParams{
		Id:                courseId,
		Name:              name,
		Description:       description,
		PromptInstruction: promptInstruction,
	})
	if err != nil {
		return nil, gut.Err(false, "failed to update course", err)
	}

	// * map course to payload course
	result := convert.Course.CourseRowToPayload(course)

	// * return
	return result, nil
}
