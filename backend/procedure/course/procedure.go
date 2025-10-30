package courseProcedure

import (
	"backend/common/config"
	"backend/service/agent"
	"backend/service/compute"
	"backend/service/path"
	"backend/type/common"
	"backend/type/payload"
	"context"
	"io"

	"github.com/bsthun/gut"
	"github.com/minio/minio-go/v7"
	"github.com/qdrant/go-client/qdrant"
)

type Proc interface {
	CourseCreate(ctx context.Context, name *string, description *string, userId *uint64) (*payload.Course, *gut.ErrorInstance)
	CourseManageEdit(ctx context.Context, courseId *uint64, name *string, description *string, promptInstruction *string) (*payload.Course, *gut.ErrorInstance)
	CourseManageDelete(ctx context.Context, courseId *uint64) (*payload.Course, *gut.ErrorInstance)
	CourseManageDetail(ctx context.Context, courseId *uint64) (*payload.Course, []*payload.EnrollInfo, []*payload.ContentInfo, *gut.ErrorInstance)
	CourseListManager(ctx context.Context, userId *uint64, name *string, limit *uint64, offset *uint64) ([]*payload.CourseExtended, *uint64, *gut.ErrorInstance)
	CourseListEnroll(ctx context.Context, userId *uint64, name *string, limit *uint64, offset *uint64) ([]*payload.CourseExtended, *uint64, *gut.ErrorInstance)
	CourseListExplore(ctx context.Context, userId *uint64, name *string, limit *uint64, offset *uint64) ([]*payload.CourseExtended, *uint64, *gut.ErrorInstance)
	CoursePhotoList(ctx context.Context, courseId *uint64, title *string, sort *string, order *string, limit *uint64, offset *uint64) ([]*payload.CoursePhoto, *uint64, *gut.ErrorInstance)
	CoursePhotoUpload(ctx context.Context, courseId *uint64, imageReader io.Reader) (*payload.CoursePhoto, *gut.ErrorInstance)
	ContentList(ctx context.Context, courseId *uint64, userId *uint64, title *string, sort *string, order *string, limit *uint64, offset *uint64) ([]*payload.ContentInfo, *uint64, *gut.ErrorInstance)
}

type Procedure struct {
	config         *config.Config
	database       common.Database
	minio          *minio.Client
	qdrant         *qdrant.Client
	computeService compute.Server
	agentService   agent.Server
	pathService    path.Server
}

func Proceed(
	config *config.Config,
	database common.Database,
	minio *minio.Client,
	qdrant *qdrant.Client,
	computeService compute.Server,
	agentService agent.Server,
	pathService path.Server,
) Proc {
	return &Procedure{
		config:         config,
		database:       database,
		minio:          minio,
		qdrant:         qdrant,
		computeService: computeService,
		agentService:   agentService,
		pathService:    pathService,
	}
}
