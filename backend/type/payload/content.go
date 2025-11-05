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

type ContentSection struct {
	Id        *uint64    `json:"id"`
	ContentId *uint64    `json:"contentId"`
	SectionNo *int32     `json:"sectionNo"`
	Title     *string    `json:"title"`
	Subtitle  *string    `json:"subtitle"`
	Content   *string    `json:"content"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type ContentIdRequest struct {
	ContentId *uint64 `json:"contentId" validate:"required"`
}

type ContentDetailResponse struct {
	Course   *Course           `json:"course"`
	Content  *Content          `json:"content"`
	Sections []*ContentSection `json:"sections"`
}
