package courseEndpoint

import (
	"backend/helper/image"
	"backend/type/common"
	"backend/type/payload"
	"backend/type/response"
	"bytes"

	"github.com/bsthun/gut"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (r *Handler) HandlePhotoUpload(c *fiber.Ctx) error {
	// * get user claims
	l := c.Locals("l").(*jwt.Token).Claims.(*common.LoginClaims)

	// * parse course id
	courseId, err := gut.IdDecode(c.FormValue("courseId"))
	if err != nil {
		return gut.Err(false, "invalid course id", err)
	}

	// * permission act
	er := r.permissionProcedure.CourseManage(c.Context(), l.UserId, &courseId, gut.Ptr("manage"))
	if er != nil {
		return er
	}

	// * get file from multipart form
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return gut.Err(false, "failed to get file", err)
	}

	// * open file
	file, err := fileHeader.Open()
	if err != nil {
		return gut.Err(false, "failed to open file", err)
	}
	defer file.Close()

	// * read file data
	fileData := make([]byte, fileHeader.Size)
	_, err = file.Read(fileData)
	if err != nil {
		return gut.Err(false, "failed to read file data", err)
	}

	// * encode photo to png
	encodedData, er := image.EncodePhotoToPng(fileData)
	if er != nil {
		return er
	}

	// * call procedure
	photo, er := r.courseProcedure.CoursePhotoUpload(c.Context(), &courseId, bytes.NewReader(encodedData))
	if er != nil {
		return er
	}

	// * return
	return c.JSON(response.Success(c, &payload.CoursePhotoUploadResponse{
		Photo: photo,
	}))
}
