package contentEndpoint

import (
	"backend/type/common"
	"backend/type/payload"
	"backend/type/response"

	"github.com/bsthun/gut"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (r *Handler) HandleContentLogList(c *fiber.Ctx) error {
	// * get user claims
	l := c.Locals("l").(*jwt.Token).Claims.(*common.LoginClaims)

	// * parse body
	body := new(payload.ContentLogListRequest)
	if err := c.BodyParser(body); err != nil {
		return gut.Err(false, "invalid body", err)
	}

	// * validate body
	if err := gut.Validate(body); err != nil {
		return err
	}

	// * permission act
	er := r.permissionProcedure.Act(c.Context(), l.UserId, l.UserId)
	if er != nil {
		return er
	}

	// * set defaults
	limit := body.Limit
	if limit == nil {
		limit = gut.Ptr(uint64(50))
	}
	offset := body.Offset
	if offset == nil {
		offset = gut.Ptr(uint64(0))
	}

	// * call procedure
	items, count, er := r.contentProcedure.ContentLogList(c.Context(), body.ContentId, limit, offset)
	if er != nil {
		return er
	}

	// * return
	return c.JSON(response.Success(c, &payload.ContentLogListResponse{
		Items: items,
		Count: count,
	}))
}
