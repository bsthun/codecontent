package agent

import (
	"backend/common/config"

	"github.com/bsthun/gut"
	"go.scnd.dev/open/model/agentic/package/call"
)

type Server interface {
	FunctionPhotoDescription(imageUrl string) (*string, *string, *gut.ErrorInstance)
	FunctionGenerateTitle(prompt string) (*string, *gut.ErrorInstance)
}

type Service struct {
	config    *config.Config
	caller    call.Caller
	ExtraArgs map[string]any
}

func Serve(
	config *config.Config,
	caller call.Caller,
) Server {
	return &Service{
		config: config,
		caller: caller,
		ExtraArgs: map[string]any{
			"thinking": map[string]any{
				"type": "disabled",
			},
		},
	}
}
