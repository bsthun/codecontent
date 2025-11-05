-- name: ContentSectionCreate :one
INSERT INTO content_sections (content_id, section_no, title, subtitle, content)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: ContentSectionGet :one
SELECT *
FROM content_sections
WHERE id = $1;

-- name: ContentSectionUpdate :one
UPDATE content_sections
SET
    content_id = COALESCE(sqlc.narg(content_id), content_id),
    section_no = COALESCE(sqlc.narg(section_no), section_no),
    title = COALESCE(sqlc.narg(title), title),
    subtitle = COALESCE(sqlc.narg(subtitle), subtitle),
    content = COALESCE(sqlc.narg(content), content)
WHERE id = $1
RETURNING *;

-- name: ContentSectionDelete :one
DELETE FROM content_sections
WHERE id = $1
RETURNING *;

-- name: ContentSectionCount :one
SELECT COALESCE(COUNT(*), 0)::BIGINT AS content_section_count
FROM content_sections
WHERE (sqlc.narg(contentId)::BIGINT IS NULL OR content_sections.content_id = sqlc.narg(contentId)::BIGINT)
  AND (sqlc.narg(title)::TEXT IS NULL OR LOWER(content_sections.title) LIKE LOWER('%' || sqlc.narg(title) || '%'));

-- name: ContentSectionListDetail :many
SELECT id, content_id, section_no, title, subtitle, content, created_at, updated_at
FROM content_sections
WHERE content_id = $1
ORDER BY section_no;
