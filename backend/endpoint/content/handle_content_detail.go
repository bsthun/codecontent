package contentEndpoint

import (
	"backend/type/common"
	"backend/type/payload"
	"backend/type/response"

	"github.com/bsthun/gut"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (r *Handler) HandleContentDetail(c *fiber.Ctx) error {
	// * get user claims
	l := c.Locals("l").(*jwt.Token).Claims.(*common.LoginClaims)

	// * parse body
	body := new(payload.ContentIdRequest)
	if err := c.BodyParser(body); err != nil {
		return gut.Err(false, "invalid body", err)
	}

	// * validate body
	if err := gut.Validate(body); err != nil {
		return err
	}

	// * permission access
	er := r.permissionProcedure.ContentAccess(c.Context(), l.UserId, body.ContentId)
	if er != nil {
		return er
	}

	// * call procedure
	contentDetail, er := r.contentProcedure.ContentDetail(c.Context(), body.ContentId)
	if er != nil {
		return er
	}

	// * return
	return c.JSON(response.Success(c, contentDetail))
}
