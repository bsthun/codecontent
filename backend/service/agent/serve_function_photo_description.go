package agent

import (
	"github.com/bsthun/gut"
)

func (r *Service) FunctionPhotoDescription(imageUrl string) (*string, *string, *gut.ErrorInstance) {
	// * mockup implementation
	title := gut.Ptr("Generated Photo Title")
	description := gut.Ptr("Generated photo description based on the image content")

	return title, description, nil
}
