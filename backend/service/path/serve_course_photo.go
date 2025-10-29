package path

import (
	"fmt"

	"github.com/bsthun/gut"
)

func (r *Service) CoursePhotoMinioPath(courseId *uint64, coursePhotoId *uint64) *string {
	path := fmt.Sprintf("/course/%s/photo/%s.png", gut.Base62(*courseId), gut.Base62(*coursePhotoId))
	return &path
}

func (r *Service) CoursePhotoMinioUrl(courseId *uint64, coursePhotoId *uint64) *string {
	url := fmt.Sprintf("%s/%s/%s", *r.config.MinioEndpoint, *r.config.MinioBucket, *r.CoursePhotoMinioPath(courseId, coursePhotoId))
	return &url
}
