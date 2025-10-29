package courseProcedure

import (
	"backend/generate/psql"
	"backend/helper/convert"
	"backend/type/payload"
	"context"

	"github.com/bsthun/gut"
)

func (r *Procedure) CourseManageDetail(ctx context.Context, courseId *uint64) (*payload.Course, []*payload.EnrollInfo, []*payload.ContentInfo, *gut.ErrorInstance) {
	// * get course
	course, err := r.database.P().CourseGet(ctx, courseId)
	if err != nil {
		return nil, nil, nil, gut.Err(false, "failed to get course", err)
	}

	// * get enroll list with user info
	enrollRows, err := r.database.P().EnrollList(ctx, &psql.EnrollListParams{
		CourseId: courseId,
		Limit:    nil,
		Offset:   nil,
	})
	if err != nil {
		return nil, nil, nil, gut.Err(false, "failed to get enroll list", err)
	}

	// * convert enroll rows to enroll info
	enrollList, _ := gut.Iterate(enrollRows, func(enroll psql.EnrollListRow) (*payload.EnrollInfo, *gut.ErrorInstance) {
		return &payload.EnrollInfo{
			Id:             enroll.Enroll.Id,
			CourseId:       enroll.Enroll.CourseId,
			UserId:         enroll.Enroll.UserId,
			UserOid:        enroll.User.Oid,
			UserFirstname:  enroll.User.Firstname,
			UserLastname:   enroll.User.Lastname,
			UserEmail:      enroll.User.Email,
			UserPictureUrl: enroll.User.PictureUrl,
			ContentCount:   enroll.ContentCount,
			CreatedAt:      enroll.Enroll.CreatedAt,
			UpdatedAt:      enroll.Enroll.UpdatedAt,
		}, nil
	})

	// * get all enroll ids for content query
	var enrollIds []*uint64
	_, _ = gut.Iterate(enrollList, func(enroll *payload.EnrollInfo) (struct{}, *gut.ErrorInstance) {
		enrollIds = append(enrollIds, enroll.Id)
		return struct{}{}, nil
	})

	// * get content list for all enrolls in this course
	var contentList []*payload.ContentInfo
	if len(enrollIds) > 0 {
		// * we need to query contents for each enroll since there's no direct course-to-contents query
		for _, enrollId := range enrollIds {
			contentRows, err := r.database.P().ContentList(ctx, &psql.ContentListParams{
				Enrollid: enrollId,
				Limit:    gut.Ptr(uint64(1000)), // large limit to get all
				Offset:   gut.Ptr(uint64(0)),
			})
			if err != nil {
				return nil, nil, nil, gut.Err(false, "failed to get content list", err)
			}

			// * convert content rows and append to list
			enrollContentList, _ := gut.Iterate(contentRows, func(content psql.ContentListRow) (*payload.ContentInfo, *gut.ErrorInstance) {
				return &payload.ContentInfo{
					Id:                  content.Content.Id,
					EnrollId:            content.Content.EnrollId,
					Title:               content.Content.Title,
					ContentSectionCount: content.ContentSectionCount,
					ContentLogCount:     content.ContentLogCount,
					CreatedAt:           content.Content.CreatedAt,
					UpdatedAt:           content.Content.UpdatedAt,
				}, nil
			})

			contentList = append(contentList, enrollContentList...)
		}
	}

	// * convert course to payload
	coursePayload := convert.Course.CourseRowToPayload(course)

	return coursePayload, enrollList, contentList, nil
}
