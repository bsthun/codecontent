package path

import (
	"context"
	"fmt"
	"time"

	"github.com/bsthun/gut"
)

func (r *Service) CoursePhotoMinioPath(courseId *uint64, coursePhotoId *uint64) *string {
	path := fmt.Sprintf("/course/%s/photo/%s.png", gut.Base62(*courseId), gut.Base62(*coursePhotoId))
	return &path
}

func (r *Service) CoursePhotoMinioUrl(courseId *uint64, coursePhotoId *uint64) *string {
	// * generate presigned url
	presignedUrl, err := r.minio.PresignedGetObject(
		context.TODO(),
		*r.config.MinioBucket,
		*r.CoursePhotoMinioPath(courseId, coursePhotoId),
		time.Hour,
		nil,
	)
	if err != nil {
		return nil
	}

	url := presignedUrl.String()

	return &url
}
