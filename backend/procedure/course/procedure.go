package courseProcedure

import (
	"backend/type/common"
	"backend/type/payload"
	"context"

	"github.com/bsthun/gut"
)

type Proc interface {
	CourseCreate(ctx context.Context, name *string, description *string, userId *uint64) (*payload.Course, *gut.ErrorInstance)
	CourseEdit(ctx context.Context, courseId *uint64, name *string, description *string, promptInstruction *string) (*payload.Course, *gut.ErrorInstance)
	CourseDelete(ctx context.Context, courseId *uint64) (*payload.Course, *gut.ErrorInstance)
	CourseManageDetail(ctx context.Context, courseId *uint64) (*payload.Course, []*payload.EnrollInfo, []*payload.ContentInfo, *gut.ErrorInstance)
	CourseListManager(ctx context.Context, userId *uint64, name *string, limit *uint64, offset *uint64) ([]*payload.CourseExtended, *uint64, *gut.ErrorInstance)
	CourseListEnroll(ctx context.Context, userId *uint64, name *string, limit *uint64, offset *uint64) ([]*payload.CourseExtended, *uint64, *gut.ErrorInstance)
	CourseListExplore(ctx context.Context, userId *uint64, name *string, limit *uint64, offset *uint64) ([]*payload.CourseExtended, *uint64, *gut.ErrorInstance)
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
