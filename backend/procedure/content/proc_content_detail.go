package contentProcedure

import (
	"backend/generate/psql"
	"backend/helper/convert"
	"backend/type/payload"
	"context"

	"github.com/bsthun/gut"
)

func (r *Procedure) ContentDetail(ctx context.Context, contentId *uint64) (*payload.ContentDetailResponse, *gut.ErrorInstance) {
	// * query content
	content, err := r.database.P().ContentGet(ctx, contentId)
	if err != nil {
		return nil, gut.Err(false, "failed to get content", err)
	}

	// * get course from enroll
	enroll, err := r.database.P().EnrollGet(ctx, content.EnrollId)
	if err != nil {
		return nil, gut.Err(false, "failed to get enroll", err)
	}

	course, err := r.database.P().CourseGet(ctx, enroll.CourseId)
	if err != nil {
		return nil, gut.Err(false, "failed to get course", err)
	}

	// * get content sections
	contentSections, err := r.database.P().ContentSectionListDetail(ctx, contentId)
	if err != nil {
		return nil, gut.Err(false, "failed to get content sections", err)
	}

	// * convert content to payload
	contentPayload := convert.Content.ContentRowToPayload(content)

	// * convert course to payload
	coursePayload := convert.Course.CourseRowToPayload(course)

	// * convert content sections to payload
	sectionsPayload, _ := gut.Iterate(contentSections, func(section psql.ContentSection) (*payload.ContentSection, *gut.ErrorInstance) {
		return &payload.ContentSection{
			Id:        section.Id,
			ContentId: section.ContentId,
			SectionNo: section.SectionNo,
			Title:     section.Title,
			Subtitle:  section.Subtitle,
			Content:   section.Content,
			CreatedAt: section.CreatedAt,
			UpdatedAt: section.UpdatedAt,
		}, nil
	})

	// * return
	return &payload.ContentDetailResponse{
		Course:   coursePayload,
		Content:  contentPayload,
		Sections: sectionsPayload,
	}, nil
}
