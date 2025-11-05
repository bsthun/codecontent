package contentProcedure

import (
	"backend/generate/psql"
	"backend/helper/convert"
	"backend/type/payload"
	"context"

	"github.com/bsthun/gut"
)

func (r *Procedure) ContentCreate(ctx context.Context, enrollId *uint64, prompt *string) (*payload.Content, *gut.ErrorInstance) {
	// * generate title from prompt
	title, er := r.agentService.FunctionGenerateTitle(*prompt)
	if er != nil {
		return nil, er
	}

	// * begin transaction
	tx, querier := r.database.Ptx(ctx, nil)
	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
		}
	}()

	// * query content create
	content, err := querier.ContentCreate(ctx, &psql.ContentCreateParams{
		EnrollId: enrollId,
		Title:    title,
	})
	if err != nil {
		_ = tx.Rollback()
		return nil, gut.Err(false, "failed to create content", err)
	}

	// * query content log create
	_, err = querier.ContentLogCreate(ctx, &psql.ContentLogCreateParams{
		ContentId: content.Id,
		Action:    gut.Ptr("initial"),
		Prompt:    prompt,
		Call:      []byte("{}"),
	})
	if err != nil {
		_ = tx.Rollback()
		return nil, gut.Err(false, "failed to create content log", err)
	}

	// * commit transaction
	if err := tx.Commit(); err != nil {
		return nil, gut.Err(false, "failed to commit transaction", err)
	}

	// * convert content to payload
	result := convert.Content.ContentRowToPayload(content)

	// * return
	return result, nil
}
