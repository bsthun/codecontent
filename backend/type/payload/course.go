package payload

import (
	"time"
)

type CourseCreateRequest struct {
	Name        *string `json:"name" validate:"required"`
	Description *string `json:"description" validate:"required"`
}

type CourseCreateResponse struct {
	Course *Course `json:"course"`
}

type Course struct {
	Id                *uint64    `json:"id"`
	Name              *string    `json:"name"`
	Description       *string    `json:"description"`
	PromptInstruction *string    `json:"promptInstruction"`
	Token             *string    `json:"token"`
	CreatedAt         *time.Time `json:"createdAt"`
	UpdatedAt         *time.Time `json:"updatedAt"`
}

type CourseExtended struct {
	Course
	CourseManagerCount *uint64 `json:"courseManagerCount"`
	EnrollCount        *uint64 `json:"enrollCount"`
	CoursePhotoCount   *uint64 `json:"coursePhotoCount"`
}

type CourseListByManagerRequest struct {
	UserId *uint64 `json:"userId" validate:"required"`
	Name   *string `json:"name"`
	Limit  *uint64 `json:"limit"`
	Offset *uint64 `json:"offset"`
}

type CourseItem struct {
	Id                 *uint64    `json:"id"`
	Name               *string    `json:"name"`
	Description        *string    `json:"description"`
	PromptInstruction  *string    `json:"promptInstruction"`
	CreatedAt          *time.Time `json:"createdAt"`
	UpdatedAt          *time.Time `json:"updatedAt"`
	CourseManagerCount *uint64    `json:"courseManagerCount"`
	EnrollCount        *uint64    `json:"enrollCount"`
	CoursePhotoCount   *uint64    `json:"coursePhotoCount"`
}

type CourseListByManagerResponse struct {
	Items []*CourseExtended `json:"items"`
	Count *uint64           `json:"count"`
}

type CourseListByEnrollRequest struct {
	UserId *uint64 `json:"userId" validate:"required"`
	Name   *string `json:"name"`
	Limit  *uint64 `json:"limit"`
	Offset *uint64 `json:"offset"`
}

type CourseListByEnrollResponse struct {
	Items []*CourseExtended `json:"items"`
	Count *uint64           `json:"count"`
}

type CourseListExploreRequest struct {
	UserId *uint64 `json:"userId" validate:"required"`
	Name   *string `json:"name"`
	Limit  *uint64 `json:"limit"`
	Offset *uint64 `json:"offset"`
}

type CourseListExploreResponse struct {
	Items []*CourseExtended `json:"items"`
	Count *uint64           `json:"count"`
}

type CourseIdRequest struct {
	CourseId *uint64 `json:"courseId" validate:"required"`
}

type CourseWrapper struct {
	Course *Course `json:"course"`
}

type CourseManageDetailRequest struct {
	CourseId *uint64 `json:"courseId" validate:"required"`
}

type EnrollInfo struct {
	Id             *uint64    `json:"id"`
	CourseId       *uint64    `json:"courseId"`
	UserId         *uint64    `json:"userId"`
	UserOid        *string    `json:"userOid"`
	UserFirstname  *string    `json:"userFirstname"`
	UserLastname   *string    `json:"userLastname"`
	UserEmail      *string    `json:"userEmail"`
	UserPictureUrl *string    `json:"userPictureUrl"`
	ContentCount   *uint64    `json:"contentCount"`
	CreatedAt      *time.Time `json:"createdAt"`
	UpdatedAt      *time.Time `json:"updatedAt"`
}

type ContentInfo struct {
	Id                  *uint64    `json:"id"`
	EnrollId            *uint64    `json:"enrollId"`
	Title               *string    `json:"title"`
	ContentSectionCount *uint64    `json:"contentSectionCount"`
	ContentLogCount     *uint64    `json:"contentLogCount"`
	CreatedAt           *time.Time `json:"createdAt"`
	UpdatedAt           *time.Time `json:"updatedAt"`
}

type CourseManageDetailResponse struct {
	Course      *Course        `json:"course"`
	EnrollList  []*EnrollInfo  `json:"enrollList"`
	ContentList []*ContentInfo `json:"contentList"`
}
