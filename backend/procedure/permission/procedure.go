package permissionProcedure

import (
	"backend/type/common"
	"context"

	"github.com/bsthun/gut"
)

type Proc interface {
	Act(ctx context.Context, userId *uint64, requestedUserId *uint64) *gut.ErrorInstance
	CourseManage(ctx context.Context, userId *uint64, courseId *uint64, level *string) *gut.ErrorInstance
}

type Procedure struct {
	database common.Database
}

func Proceed(
	database common.Database,
) Proc {
	return &Procedure{
		database: database,
	}
}
