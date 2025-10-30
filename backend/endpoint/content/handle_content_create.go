package contentEndpoint

import (
	"backend/generate/psql"
	"backend/type/common"
	"backend/type/payload"
	"backend/type/response"

	"github.com/bsthun/gut"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (r *Handler) HandleContentCreate(c *fiber.Ctx) error {
	// * get user claims
	l := c.Locals("l").(*jwt.Token).Claims.(*common.LoginClaims)

	// * parse body
	body := new(payload.ContentCreateRequest)
	if err := c.BodyParser(body); err != nil {
		return gut.Err(false, "invalid body", err)
	}

	// * validate body
	if err := gut.Validate(body); err != nil {
		return err
	}

	// * get or create enroll
	enroll, err := r.database.P().EnrollGetOrCreate(c.Context(), &psql.EnrollGetOrCreateParams{
		CourseId: body.CourseId,
		UserId:   l.UserId,
	})
	if err != nil {
		return gut.Err(false, "failed to get or create enroll", err)
	}

	// * call procedure
	content, er := r.contentProcedure.ContentCreate(c.Context(), enroll.Id, body.Prompt)
	if er != nil {
		return er
	}

	// * return
	return c.JSON(response.Success(c, &payload.ContentWrapper{
		Content: content,
	}))
}
