package courseProcedure

import (
	"backend/type/common"
	"backend/type/payload"
	"context"
	"github.com/bsthun/gut"
)

type Proc interface {
	CourseCreate(ctx context.Context, params *payload.CourseCreateParams) (*payload.Course, *gut.ErrorInstance)
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
