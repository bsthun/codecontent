package main

import (
	"backend/common/agentic"
	"backend/common/config"
	"backend/common/database"
	"backend/common/fiber"
	"backend/common/fiber/middleware"
	"backend/endpoint"
	courseEndpoint "backend/endpoint/course"
	publicEndpoint "backend/endpoint/public"
	stateEndpoint "backend/endpoint/state"
	courseProcedure "backend/procedure/course"
	"backend/type/common"
	"embed"

	"go.uber.org/fx"
)

//go:embed database/postgres/migration/*.sql
var migration embed.FS

//go:embed .local/dist/*
var frontend embed.FS

func main() {
	fx.New(
		fx.Provide(
			func() common.MigrationFS {
				return migration
			},
			func() common.FrontendFS {
				return frontend
			},
			config.Init,
			agentic.Init,
			database.Init,
			fiber.Init,
			middleware.Init,
			courseProcedure.Proceed,
			publicEndpoint.Handle,
			stateEndpoint.Handle,
			courseEndpoint.Handle,
		),
		fx.Invoke(
			endpoint.Bind,
		),
	).Run()
}
