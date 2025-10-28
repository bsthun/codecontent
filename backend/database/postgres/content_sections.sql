-- name: ContentSectionCreate :one
INSERT INTO content_sections (content_id, section_no, title, subtitle, content)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: ContentSectionGetById :one
SELECT *
FROM content_sections
WHERE id = $1;

-- name: ContentSectionUpdateById :one
UPDATE content_sections
SET
    content_id = COALESCE(sqlc.narg(content_id), content_id),
    section_no = COALESCE(sqlc.narg(section_no), section_no),
    title = COALESCE(sqlc.narg(title), title),
    subtitle = COALESCE(sqlc.narg(subtitle), subtitle),
    content = COALESCE(sqlc.narg(content), content)
WHERE id = $1
RETURNING *;

-- name: ContentSectionDeleteById :exec
DELETE FROM content_sections
WHERE id = $1;

-- name: ContentSectionList :many
SELECT *
FROM content_sections
ORDER BY created_at DESC;

-- name: ContentSectionCount :one
SELECT COALESCE(COUNT(*), 0)::BIGINT AS content_section_count
FROM content_sections;

-- name: ContentSectionListByContentId :many
SELECT *
FROM content_sections
WHERE content_id = $1
ORDER BY section_no ASC;