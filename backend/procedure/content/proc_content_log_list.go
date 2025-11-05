package contentProcedure

import (
	"backend/generate/psql"
	"backend/type/payload"
	"context"

	"github.com/bsthun/gut"
)

func (r *Procedure) ContentLogList(ctx context.Context, contentId *uint64, limit *uint64, offset *uint64) ([]*payload.ContentLog, *uint64, *gut.ErrorInstance) {
	// * query: content log count
	count, err := r.database.P().ContentLogCount(ctx, contentId)
	if err != nil {
		return nil, nil, gut.Err(false, "failed to get content log count", err)
	}

	// * query: content log list
	logRows, err := r.database.P().ContentLogList(ctx, &psql.ContentLogListParams{
		ContentId: contentId,
		Limit:     limit,
		Offset:    offset,
		Sort:      gut.Ptr("createdAt"),
		Order:     gut.Ptr("desc"),
	})
	if err != nil {
		return nil, nil, gut.Err(false, "failed to get content log list", err)
	}

	// * map content logs to items
	logItems, er := gut.Iterate(logRows, func(row psql.ContentLogListRow) (*payload.ContentLog, *gut.ErrorInstance) {
		// * convert call json to string
		callString := string(row.ContentLog.Call)
		return &payload.ContentLog{
			Id:        row.ContentLog.Id,
			ContentId: row.ContentLog.ContentId,
			Action:    row.ContentLog.Action,
			Prompt:    row.ContentLog.Prompt,
			Call:      &callString,
			CreatedAt: row.ContentLog.CreatedAt,
			UpdatedAt: row.ContentLog.UpdatedAt,
		}, nil
	})
	if er != nil {
		return nil, nil, er
	}

	// * return
	return logItems, count, nil
}
