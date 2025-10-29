package convert

import (
	"backend/generate/psql"
	"backend/type/payload"
)

// goverter:converter
// goverter:output:file ../../generate/convert/course.go
type CourseConverter interface {
	CourseRowToPayload(source psql.Course) *payload.Course
	CourseRowsToPayload(source []psql.Course) []*payload.Course
	CourseListByManagerRowToPayload(source psql.CourseListByManagerRow) *payload.CourseExtended
	CourseListByManagerRowsToPayload(source []psql.CourseListByManagerRow) []*payload.CourseExtended
	CourseListExploreRowToPayload(source psql.CourseListExploreRow) *payload.CourseExtended
	CourseListExploreRowsToPayload(source []psql.CourseListExploreRow) []*payload.CourseExtended
	CourseListByEnrollRowToPayload(source psql.CourseListByEnrollRow) *payload.CourseExtended
	CourseListByEnrollRowsToPayload(source []psql.CourseListByEnrollRow) []*payload.CourseExtended
}
