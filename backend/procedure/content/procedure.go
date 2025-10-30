package contentProcedure

import (
	"backend/common/config"
	"backend/type/common"
	"backend/type/payload"
	"context"

	"github.com/bsthun/gut"
)

type Proc interface {
	ContentCreate(ctx context.Context, enrollId *uint64, prompt *string) (*payload.Content, *gut.ErrorInstance)
}

type Procedure struct {
	config   *config.Config
	database common.Database
}

func Proceed(
	config *config.Config,
	database common.Database,
) Proc {
	return &Procedure{
		config:   config,
		database: database,
	}
}
