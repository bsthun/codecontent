package courseEndpoint

import (
	"backend/procedure/course"
	permissionProcedure "backend/procedure/permission"
)

type Handler struct {
	courseProcedure     courseProcedure.Proc
	permissionProcedure permissionProcedure.Proc
}

func Handle(
	courseProcedure courseProcedure.Proc,
	permissionProcedure permissionProcedure.Proc,
) *Handler {
	return &Handler{
		courseProcedure:     courseProcedure,
		permissionProcedure: permissionProcedure,
	}
}
