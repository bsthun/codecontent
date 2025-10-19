package agentic

import (
	"backend/common/config"

	"go.scnd.dev/open/model-agentic/package/call"
)

func Init(config *config.Config) call.Caller {
	caller := call.NewOpenai(*config.OpenaiBaseUrl, *config.OpenaiApiKey)
	return caller
}
