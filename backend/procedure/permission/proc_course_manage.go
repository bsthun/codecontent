package permissionProcedure

import (
	"backend/generate/psql"
	"context"

	"github.com/bsthun/gut"
)

func (r *Procedure) CourseManage(ctx context.Context, userId *uint64, courseId *uint64, level *string) *gut.ErrorInstance {
	// * query: user get
	user, er := r.database.P().UserGet(ctx, userId)
	if er != nil {
		return gut.Err(false, "failed to get user", er)
	}

	// * check if user is admin
	if *user.IsAdmin {
		return nil
	}

	// * query: course manager get by user and course
	_, er = r.database.P().CourseManagerGetByUserAndCourse(ctx, &psql.CourseManagerGetByUserAndCourseParams{
		UserId:   userId,
		CourseId: courseId,
	})
	if er != nil {
		return gut.Err(false, "user is not a course manager", nil)
	}

	return nil
}
