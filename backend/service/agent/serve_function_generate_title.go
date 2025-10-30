package agent

import (
	"backend/type/external"

	"github.com/bsthun/gut"
	"go.scnd.dev/open/model/agentic/package/call"
)

func (r *Service) FunctionGenerateTitle(prompt string) (*string, *gut.ErrorInstance) {
	// * prepare request
	request := &call.Request{
		Model:     r.config.OpenaiModel,
		MaxTokens: gut.Ptr(64),
		Messages: []*call.Message{
			{
				Role: gut.Ptr("user"),
				Content: gut.Ptr("Generate a short, descriptive title (max 32 characters) for content based on this prompt: " + prompt +
					"The title should be concise but informative. Response in json `{ \"title\": string }`"),
			},
		},
	}

	// * prepare option for structured output
	option := &call.Option{
		SchemaName:        gut.Ptr("TitleGenerationResult"),
		SchemaDescription: gut.Ptr("Content title generation result"),
	}

	// * prepare output struct
	output := new(external.TitleGenerationResult)

	// * call agentic service
	_, er := r.caller.Call(request, option, output)
	if er != nil {
		return nil, gut.Err(false, "failed to call agentic service", er)
	}

	return output.Title, nil
}
