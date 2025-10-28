-- name: ContentSectionPhotoCreate :one
INSERT INTO content_section_photos (content_section_id, course_photo_id)
VALUES ($1, $2)
RETURNING *;

-- name: ContentSectionPhotoGet :one
SELECT *
FROM content_section_photos
WHERE id = $1;

-- name: ContentSectionPhotoUpdate :one
UPDATE content_section_photos
SET
    content_section_id = COALESCE(sqlc.narg(content_section_id), content_section_id),
    course_photo_id = COALESCE(sqlc.narg(course_photo_id), course_photo_id)
WHERE id = $1
RETURNING *;

-- name: ContentSectionPhotoDelete :one
DELETE FROM content_section_photos
WHERE id = $1
RETURNING *;

-- name: ContentSectionPhotoList :many
SELECT sqlc.embed(content_section_photos),
       sqlc.embed(content_sections),
       sqlc.embed(course_photos)
FROM content_section_photos
LEFT JOIN content_sections ON content_section_photos.content_section_id = content_sections.id
LEFT JOIN course_photos ON content_section_photos.course_photo_id = course_photos.id
WHERE (sqlc.narg(contentSectionId)::BIGINT IS NULL OR content_section_photos.content_section_id = sqlc.narg(contentSectionId)::BIGINT)
  AND (sqlc.narg(coursePhotoId)::BIGINT IS NULL OR content_section_photos.course_photo_id = sqlc.narg(coursePhotoId)::BIGINT)
GROUP BY content_section_photos.id, content_sections.id, course_photos.id
ORDER BY
  CASE WHEN sqlc.narg('sort') = 'createdAt' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN content_section_photos.created_at END,
  CASE WHEN sqlc.narg('sort') = 'createdAt' AND sqlc.narg('order') = 'desc' THEN content_section_photos.created_at END DESC,
  CASE WHEN sqlc.narg('sort') = 'updatedAt' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN content_section_photos.updated_at END,
  CASE WHEN sqlc.narg('sort') = 'updatedAt' AND sqlc.narg('order') = 'desc' THEN content_section_photos.updated_at END DESC
LIMIT sqlc.narg('limit')::BIGINT
OFFSET COALESCE(sqlc.narg('offset')::BIGINT, 0);

-- name: ContentSectionPhotoCount :one
SELECT COALESCE(COUNT(*), 0)::BIGINT AS content_section_photo_count
FROM content_section_photos
WHERE (sqlc.narg(contentSectionId)::BIGINT IS NULL OR content_section_photos.content_section_id = sqlc.narg(contentSectionId)::BIGINT)
  AND (sqlc.narg(coursePhotoId)::BIGINT IS NULL OR content_section_photos.course_photo_id = sqlc.narg(coursePhotoId)::BIGINT);

