package courseEndpoint

import (
	"backend/type/common"
	"backend/type/payload"
	"backend/type/response"

	"github.com/bsthun/gut"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (r *Handler) HandleContentList(c *fiber.Ctx) error {
	// * get user claims
	l := c.Locals("l").(*jwt.Token).Claims.(*common.LoginClaims)

	// * parse body
	body := new(payload.ContentListRequest)
	if err := c.BodyParser(body); err != nil {
		return gut.Err(false, "invalid body", err)
	}

	// * validate body
	if err := gut.Validate(body); err != nil {
		return err
	}

	// * permission check for course manage
	er := r.permissionProcedure.CourseManage(c.Context(), l.UserId, body.CourseId, gut.Ptr("manage"))
	if er != nil {
		return er
	}

	// * call procedure
	items, count, er := r.courseProcedure.ContentList(c.Context(), body.CourseId, body.UserId, body.Title, body.Sort, body.Order, body.Limit, body.Offset)
	if er != nil {
		return er
	}

	// * return
	return c.JSON(response.Success(c, &payload.ContentListResponse{
		Items: items,
		Count: count,
	}))
}