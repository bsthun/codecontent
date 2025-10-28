package courseEndpoint

import (
	"backend/type/common"
	"backend/type/payload"
	"backend/type/response"
	"github.com/bsthun/gut"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (r *Handler) HandleCourseCreate(c *fiber.Ctx) error {
	// * get user claims
	l := c.Locals("l").(*jwt.Token).Claims.(*common.LoginClaims)

	// * parse body
	body := new(payload.CourseCreateRequest)
	if err := c.BodyParser(body); err != nil {
		return gut.Err(false, "invalid body", err)
	}

	// * validate body
	if err := gut.Validate(body); err != nil {
		return err
	}

	// * call procedure
	course, er := r.courseProcedure.CourseCreate(c.Context(), &payload.CourseCreateParams{
		Name:        body.Name,
		Description: body.Description,
		UserId:      l.UserId,
	})
	if er != nil {
		return er
	}

	// * return
	return c.JSON(response.Success(c, &payload.CourseCreateResponse{
		Course: course,
	}))
}
