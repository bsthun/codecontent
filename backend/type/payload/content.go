package payload

import (
	"time"
)

type ContentCreateRequest struct {
	CourseId *uint64 `json:"courseId" validate:"required"`
	Prompt   *string `json:"prompt" validate:"required"`
}

type Content struct {
	Id        *uint64    `json:"id"`
	EnrollId  *uint64    `json:"enrollId"`
	Title     *string    `json:"title"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type ContentWrapper struct {
	Content *Content `json:"content"`
}
