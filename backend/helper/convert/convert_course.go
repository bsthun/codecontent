package convert

import (
	"backend/generate/psql"
	"backend/type/payload"
)

// goverter:converter
// goverter:output:file ../../generate/convert/course.go
type CourseConverter interface {
	CoursePayloadsFromCourseRows(source []psql.Course) []*payload.Course
	CoursePayloadFromCourseRow(source psql.Course) *payload.Course
}
