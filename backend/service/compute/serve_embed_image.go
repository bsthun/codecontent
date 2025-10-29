package compute

import (
	"backend/type/external"
	"io"

	"github.com/bsthun/gut"
	"github.com/go-resty/resty/v2"
)

func (r *Service) EmbedImage(imageReader io.Reader) (*external.EmbeddingResponse, *gut.ErrorInstance) {
	// * prepare request
	client := resty.New()
	embedding := new(external.EmbeddingResponse)

	_, err := client.R().
		SetHeader("Content-Type", "multipart/form-data").
		SetMultipartField("image", "image.png", "image/png", imageReader).
		SetResult(embedding).
		Post("http://10.2.1.179:8001/embed")

	if err != nil {
		return nil, gut.Err(false, "failed to call embedding service", err)
	}

	return embedding, nil
}
