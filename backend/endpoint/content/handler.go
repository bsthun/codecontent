package contentEndpoint

import (
	"backend/type/common"
	contentProcedure "backend/procedure/content"
)

type Handler struct {
	database         common.Database
	contentProcedure contentProcedure.Proc
}

func Handle(
	database common.Database,
	contentProcedure contentProcedure.Proc,
) *Handler {
	return &Handler{
		database:         database,
		contentProcedure: contentProcedure,
	}
}