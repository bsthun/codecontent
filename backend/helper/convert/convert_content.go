package convert

import (
	"backend/generate/psql"
	"backend/type/payload"
)

// goverter:converter
// goverter:output:file ../../generate/convert/content.go
type ContentConverter interface {
	ContentRowToPayload(source psql.Content) *payload.Content
	ContentRowsToPayload(source []psql.Content) []*payload.Content
}