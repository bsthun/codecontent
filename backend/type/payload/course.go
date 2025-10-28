package payload

import (
	"time"
)

type CourseCreateRequest struct {
	Name        *string `json:"name" validate:"required"`
	Description *string `json:"description" validate:"required"`
}

type CourseCreateParams struct {
	Name        *string
	Description *string
	UserId      *uint64
}

type CourseCreateResponse struct {
	Course *Course `json:"course"`
}

type Course struct {
	Id                *uint64    `json:"id"`
	Name              *string    `json:"name"`
	Description       *string    `json:"description"`
	PromptInstruction *string    `json:"promptInstruction"`
	CreatedAt         *time.Time `json:"createdAt"`
	UpdatedAt         *time.Time `json:"updatedAt"`
}
