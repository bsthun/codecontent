package agent

import (
	"github.com/bsthun/gut"
)

type Server interface {
	FunctionPhotoDescription(imageUrl string) (*string, *string, *gut.ErrorInstance)
}

type Service struct {
}

func Serve() Server {
	return &Service{}
}
