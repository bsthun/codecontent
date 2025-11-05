package contentEndpoint

import (
	contentProcedure "backend/procedure/content"
	permissionProcedure "backend/procedure/permission"
	"backend/type/common"
)

type Handler struct {
	database            common.Database
	contentProcedure    contentProcedure.Proc
	permissionProcedure permissionProcedure.Proc
}

func Handle(
	database common.Database,
	contentProcedure contentProcedure.Proc,
	permissionProcedure permissionProcedure.Proc,
) *Handler {
	return &Handler{
		database:            database,
		contentProcedure:    contentProcedure,
		permissionProcedure: permissionProcedure,
	}
}
