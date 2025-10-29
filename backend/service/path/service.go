package path

import "backend/common/config"

type Server interface {
	CoursePhotoMinioPath(courseId *uint64, photoId *uint64) *string
	CoursePhotoMinioUrl(courseId *uint64, photoId *uint64) *string
}

type Service struct {
	config *config.Config
}

func Serve(
	config *config.Config,
) Server {
	return &Service{
		config: config,
	}
}
