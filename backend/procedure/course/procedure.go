package courseProcedure

import (
	"backend/type/common"
	"backend/type/payload"
	"context"

	"github.com/bsthun/gut"
)

type Proc interface {
	CourseCreate(ctx context.Context, name *string, description *string, userId *uint64) (*payload.Course, *gut.ErrorInstance)
	CourseListByManager(ctx context.Context, userId *uint64, name *string, limit *uint64, offset *uint64) ([]*payload.CourseExtended, *uint64, *gut.ErrorInstance)
}

type Procedure struct {
	database common.Database
}

func Proceed(
	database common.Database,
) *Procedure {
	return &Procedure{
		database: database,
	}
}
