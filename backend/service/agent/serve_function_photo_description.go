package agent

import (
	"backend/type/external"

	"github.com/bsthun/gut"
	"go.scnd.dev/open/model/agentic/package/call"
)

func (r *Service) FunctionPhotoDescription(imageUrl string) (*string, *string, *gut.ErrorInstance) {
	// * prepare request
	request := &call.Request{
		Model:     r.config.OpenaiVisionModel,
		MaxTokens: gut.Ptr(1024),
		Messages: []*call.Message{
			{
				Role: gut.Ptr("user"),
				Content: gut.Ptr("Analyze this image and generate a descriptive title and detailed description." +
					"The title should be concise but descriptive, and the description should explain what's happening." +
					"Response in json `{ \"title\": string, \"description\": string }`"),
				ImageUrl: &imageUrl,
			},
		},
	}

	// * prepare option for structured output
	option := &call.Option{
		SchemaName:        gut.Ptr("PhotoDescriptionResult"),
		SchemaDescription: gut.Ptr("Photo analysis result with title and description"),
	}

	// * prepare output struct
	output := new(external.PhotoDescriptionResult)

	// * call agentic service
	_, er := r.caller.Call(request, option, output)
	if er != nil {
		return nil, nil, gut.Err(false, "failed to call agentic service", er)
	}

	return output.Title, output.Description, nil
}
