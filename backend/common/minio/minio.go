package minio

import (
	"backend/common/config"
	"net/url"

	"github.com/bsthun/gut"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func Init(config *config.Config) *minio.Client {
	// * initialize minio client
	parsed, err := url.Parse(*config.MinioEndpoint)
	if err != nil {
		gut.Fatal("failed to parse minio endpoint", err)
	}

	minioClient, err := minio.New(parsed.Host, &minio.Options{
		Creds:  credentials.NewStaticV4(*config.MinioAccessKey, *config.MinioSecretKey, ""),
		Secure: parsed.Scheme == "https",
	})

	if err != nil {
		gut.Fatal("failed to initialize minio", err)
	}

	return minioClient
}
