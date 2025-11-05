package contentProcedure

import (
	"backend/common/config"
	"backend/service/agent"
	"backend/type/common"
	"backend/type/payload"
	"context"

	"github.com/bsthun/gut"
)

type Proc interface {
	ContentCreate(ctx context.Context, enrollId *uint64, prompt *string) (*payload.Content, *gut.ErrorInstance)
	ContentDetail(ctx context.Context, contentId *uint64) (*payload.ContentDetailResponse, *gut.ErrorInstance)
	ContentLogList(ctx context.Context, contentId *uint64, limit *uint64, offset *uint64) ([]*payload.ContentLog, *uint64, *gut.ErrorInstance)
}

type Procedure struct {
	config       *config.Config
	database     common.Database
	agentService agent.Server
}

func Proceed(
	config *config.Config,
	database common.Database,
	agentService agent.Server,
) Proc {
	return &Procedure{
		config:       config,
		database:     database,
		agentService: agentService,
	}
}
