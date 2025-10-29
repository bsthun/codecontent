package config

import (
	"backend/type/enum"
	"os"

	"github.com/bsthun/gut"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Environment       *enum.Environment `yaml:"environment" validate:"required"`
	WebListen         [2]*string        `yaml:"webListen" validate:"required"`
	FrontendUrl       *string           `yaml:"frontendUrl" validate:"required"`
	Secret            *string           `yaml:"secret" validate:"required"`
	PostgresDsn       *string           `yaml:"postgresDsn" validate:"required"`
	QdrantDsn         *string           `yaml:"qdrantDsn" validate:"required"`
	QdrantCollection  *string           `yaml:"qdrantCollection" validate:"required"`
	MinioEndpoint     *string           `yaml:"minioEndpoint" validate:"required"`
	MinioBucket       *string           `yaml:"minioBucket" validate:"required"`
	MinioAccessKey    *string           `yaml:"minioAccessKey" validate:"required"`
	MinioSecretKey    *string           `yaml:"minioSecretKey" validate:"required"`
	OauthClientId     *string           `yaml:"oauthClientId" validate:"required"`
	OauthClientSecret *string           `yaml:"oauthClientSecret" validate:"required"`
	OauthEndpoint     *string           `yaml:"oauthEndpoint" validate:"required"`
	OpenaiBaseUrl     *string           `yaml:"openaiBaseUrl" validate:"required"`
	OpenaiApiKey      *string           `yaml:"openaiApiKey" validate:"required"`
	OpenaiModel       *string           `yaml:"openaiModel" validate:"required"`
	OpenaiVisionModel *string           `yaml:"openaiVisionModel" validate:"required"`
	DockerUri         *string           `yaml:"dockerUri" validate:"required"`
}

func Init() *Config {
	// * parse arguments
	path := os.Getenv("BACKEND_CONFIG_PATH")
	if path == "" {
		path = "config.yml"
	}

	// * declare struct
	config := new(Config)

	// * read config
	yml, err := os.ReadFile(path)
	if err != nil {
		gut.Fatal("Unable to read configuration file", err)
	}

	// * parse config
	if err := yaml.Unmarshal(yml, config); err != nil {
		gut.Fatal("Unable to parse configuration file", err)
	}

	// * validate config
	if err := gut.Validate(config); err != nil {
		gut.Fatal("Invalid configuration", err)
	}

	// * apply secret key
	var bytes = []byte(*config.Secret)
	if len(bytes) < 16 {
		for i := len(bytes); i < 16; i++ {
			bytes = append(bytes, 0)
		}
	}
	if err := gut.SetIdEncoderKey(bytes[:16]); err != nil {
		gut.Fatal("unable to set secret key", err)
	}

	return config
}
