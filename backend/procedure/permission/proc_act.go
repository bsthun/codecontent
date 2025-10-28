package permissionProcedure

import (
	"context"

	"github.com/bsthun/gut"
)

func (r *Procedure) Act(ctx context.Context, userId *uint64, requestedUserId *uint64) *gut.ErrorInstance {
	// * query: user get
	user, err := r.database.P().UserGet(ctx, userId)
	if err != nil {
		return gut.Err(false, "failed to get user", err)
	}

	if !*user.IsAdmin && *user.Id != *requestedUserId {
		return gut.Err(false, "insufficient shadow permission", nil)
	}

	return nil
}
