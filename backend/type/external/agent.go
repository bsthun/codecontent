package external

type PhotoDescriptionResult struct {
	Title       *string `json:"title" description:"A descriptive title for the photo" validate:"required"`
	Description *string `json:"description" description:"A detailed description of what's in the photo" validate:"required"`
}

type TitleGenerationResult struct {
	Title *string `json:"title" description:"A short, descriptive title for the content" validate:"required"`
}
