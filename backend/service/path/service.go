package path

import (
	"backend/common/config"

	"github.com/minio/minio-go/v7"
)

type Server interface {
	CoursePhotoMinioPath(courseId *uint64, photoId *uint64) *string
	CoursePhotoMinioUrl(courseId *uint64, photoId *uint64) *string
}

type Service struct {
	config *config.Config
	minio  *minio.Client
}

func Serve(
	config *config.Config,
	minio *minio.Client,
) Server {
	return &Service{
		config: config,
		minio:  minio,
	}
}
