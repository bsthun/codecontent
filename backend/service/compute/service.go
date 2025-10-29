package compute

import (
	"backend/type/external"
	"io"

	"github.com/bsthun/gut"
)

type Server interface {
	EmbedImage(imageReader io.Reader) (*external.EmbeddingResponse, *gut.ErrorInstance)
}

type Service struct {
}

func Serve() Server {
	return &Service{}
}
