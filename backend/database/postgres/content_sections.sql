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

-- name: ContentSectionList :many
SELECT content_sections.id, content_sections.content_id, content_sections.section_no, content_sections.title, content_sections.subtitle, content_sections.created_at, content_sections.updated_at,
       contents.id, contents.enroll_id, contents.title, contents.created_at, contents.updated_at,
       (SELECT COALESCE(COUNT(*), 0)::BIGINT FROM content_section_photos WHERE content_section_photos.content_section_id = content_sections.id) AS content_section_photo_count
FROM content_sections
LEFT JOIN contents ON content_sections.content_id = contents.id
WHERE (sqlc.narg(contentId)::BIGINT IS NULL OR content_sections.content_id = sqlc.narg(contentId)::BIGINT)
  AND (sqlc.narg(title)::TEXT IS NULL OR LOWER(content_sections.title) LIKE LOWER('%' || sqlc.narg(title) || '%'))
GROUP BY content_sections.id, contents.id
ORDER BY
  CASE WHEN sqlc.narg('sort') = 'sectionNo' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN content_sections.section_no END,
  CASE WHEN sqlc.narg('sort') = 'sectionNo' AND sqlc.narg('order') = 'desc' THEN content_sections.section_no END DESC,
  CASE WHEN sqlc.narg('sort') = 'title' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN content_sections.title END,
  CASE WHEN sqlc.narg('sort') = 'title' AND sqlc.narg('order') = 'desc' THEN content_sections.title END DESC,
  CASE WHEN sqlc.narg('sort') = 'createdAt' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN content_sections.created_at END,
  CASE WHEN sqlc.narg('sort') = 'createdAt' AND sqlc.narg('order') = 'desc' THEN content_sections.created_at END DESC,
  CASE WHEN sqlc.narg('sort') = 'updatedAt' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN content_sections.updated_at END,
  CASE WHEN sqlc.narg('sort') = 'updatedAt' AND sqlc.narg('order') = 'desc' THEN content_sections.updated_at END DESC
LIMIT sqlc.narg('limit')::BIGINT
OFFSET COALESCE(sqlc.narg('offset')::BIGINT, 0);

-- name: ContentSectionCount :one
SELECT COALESCE(COUNT(*), 0)::BIGINT AS content_section_count
FROM content_sections
WHERE (sqlc.narg(contentId)::BIGINT IS NULL OR content_sections.content_id = sqlc.narg(contentId)::BIGINT)
  AND (sqlc.narg(title)::TEXT IS NULL OR LOWER(content_sections.title) LIKE LOWER('%' || sqlc.narg(title) || '%'));

