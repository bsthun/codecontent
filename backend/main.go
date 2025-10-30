package main

import (
	"backend/common/agentic"
	"backend/common/config"
	"backend/common/database"
	"backend/common/fiber"
	"backend/common/fiber/middleware"
	"backend/common/minio"
	"backend/common/qdrant"
	"backend/endpoint"
	contentEndpoint "backend/endpoint/content"
	courseEndpoint "backend/endpoint/course"
	publicEndpoint "backend/endpoint/public"
	stateEndpoint "backend/endpoint/state"
	contentProcedure "backend/procedure/content"
	courseProcedure "backend/procedure/course"
	permissionProcedure "backend/procedure/permission"
	"backend/service/agent"
	"backend/service/compute"
	"backend/service/path"
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
			database.Init,
			qdrant.Init,
			minio.Init,
			agentic.Init,
			fiber.Init,
			middleware.Init,
			compute.Serve,
			agent.Serve,
			path.Serve,
			permissionProcedure.Proceed,
			courseProcedure.Proceed,
			contentProcedure.Proceed,
			publicEndpoint.Handle,
			stateEndpoint.Handle,
			courseEndpoint.Handle,
			contentEndpoint.Handle,
		),
		fx.Invoke(
			endpoint.Bind,
		),
	).Run()
}
