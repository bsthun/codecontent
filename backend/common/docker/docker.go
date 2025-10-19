package docker

import (
	"backend/common/config"

	"github.com/bsthun/gut"
	"github.com/moby/moby/client"
)

func Init(config config.Config) *client.Client {
	cli, err := client.NewClientWithOpts(client.WithHost(*config.DockerUri), client.WithAPIVersionNegotiation())
	if err != nil {
		gut.Fatal("failed to create docker client", err)
	}

	return cli
}
