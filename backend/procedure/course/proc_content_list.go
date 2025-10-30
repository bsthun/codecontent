package courseProcedure

import (
	"backend/generate/psql"
	"backend/type/payload"
	"context"

	"github.com/bsthun/gut"
)

func (r *Procedure) ContentList(ctx context.Context, courseId *uint64, userId *uint64, title *string, sort *string, order *string, limit *uint64, offset *uint64) ([]*payload.ContentInfo, *uint64, *gut.ErrorInstance) {
	// * get content list by course
	contentRows, er := r.database.P().ContentListByCourse(ctx, &psql.ContentListByCourseParams{
		Courseid: courseId,
		Userid:   userId,
		Title:    title,
		Sort:     sort,
		Order:    order,
		Limit:    limit,
		Offset:   offset,
	})
	if er != nil {
		return nil, nil, gut.Err(false, "failed to get content list", er)
	}

	// * get total count
	count, er := r.database.P().ContentCountByCourse(ctx, &psql.ContentCountByCourseParams{
		Courseid: courseId,
		Userid:   userId,
		Title:    title,
	})
	if er != nil {
		return nil, nil, gut.Err(false, "failed to get content count", er)
	}

	// * convert to payload
	contentItems, _ := gut.Iterate(contentRows, func(content psql.ContentListByCourseRow) (*payload.ContentInfo, *gut.ErrorInstance) {
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

	return contentItems, count, nil
}