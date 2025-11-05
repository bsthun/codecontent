package permissionProcedure

import (
	"backend/generate/psql"
	"context"

	"github.com/bsthun/gut"
)

func (r *Procedure) ContentAccess(ctx context.Context, userId *uint64, contentId *uint64) *gut.ErrorInstance {
	// * query: user get
	user, er := r.database.P().UserGet(ctx, userId)
	if er != nil {
		return gut.Err(false, "failed to get user", er)
	}

	// * check if user is admin
	if *user.IsAdmin {
		return nil
	}

	// * query: content get
	content, er := r.database.P().ContentGet(ctx, contentId)
	if er != nil {
		return gut.Err(false, "failed to get content", er)
	}

	// * query: enroll get
	enroll, er := r.database.P().EnrollGet(ctx, content.EnrollId)
	if er != nil {
		return gut.Err(false, "failed to get enroll", er)
	}

	// * check if user is enrolled in the course
	if *enroll.UserId == *userId {
		return nil
	}

	// * check if user is a course manager
	_, er = r.database.P().CourseManagerGetByUserAndCourse(ctx, &psql.CourseManagerGetByUserAndCourseParams{
		UserId:   userId,
		CourseId: enroll.CourseId,
	})
	if er != nil {
		return gut.Err(false, "user does not have access to this content", nil)
	}

	return nil
}
