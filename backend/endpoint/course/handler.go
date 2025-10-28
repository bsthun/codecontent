package courseEndpoint

import (
	"backend/procedure/course"
)

type Handler struct {
	courseProcedure courseProcedure.Proc
}

func Handle(
	courseProcedure courseProcedure.Proc,
) *Handler {
	return &Handler{
		courseProcedure: courseProcedure,
	}
}
