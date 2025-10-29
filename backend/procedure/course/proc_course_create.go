package courseProcedure

import (
	"backend/generate/psql"
	"backend/helper/convert"
	"backend/type/payload"
	"context"

	"github.com/bsthun/gut"
)

func (r *Procedure) CourseCreate(ctx context.Context, name *string, description *string, userId *uint64) (*payload.Course, *gut.ErrorInstance) {
	// * begin transaction
	tx, querier := r.database.Ptx(ctx, nil)
	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
		}
	}()

	// * query course create
	course, err := querier.CourseCreate(ctx, &psql.CourseCreateParams{
		Name:              name,
		Description:       description,
		PromptInstruction: gut.Ptr(""),
		Token:             gut.Random(gut.RandomSet.MixedAlphaNum, 16),
	})
	if err != nil {
		_ = tx.Rollback()
		return nil, gut.Err(false, "failed to create course", err)
	}

	// * query course manager create
	_, err = querier.CourseManagerCreate(ctx, &psql.CourseManagerCreateParams{
		CourseId: course.Id,
		UserId:   userId,
	})
	if err != nil {
		_ = tx.Rollback()
		return nil, gut.Err(false, "failed to create course manager", err)
	}

	// * commit transaction
	if err := tx.Commit(); err != nil {
		return nil, gut.Err(false, "failed to commit transaction", err)
	}

	// * map course to payload course
	result := convert.Course.CourseRowToPayload(course)

	// * return
	return result, nil
}
