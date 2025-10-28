-- name: ContentSectionPhotoCreate :one
INSERT INTO content_section_photos (content_section_id, course_photo_id)
VALUES ($1, $2)
RETURNING *;

-- name: ContentSectionPhotoGetById :one
SELECT *
FROM content_section_photos
WHERE id = $1;

-- name: ContentSectionPhotoUpdateById :one
UPDATE content_section_photos
SET
    content_section_id = COALESCE(sqlc.narg(content_section_id), content_section_id),
    course_photo_id = COALESCE(sqlc.narg(course_photo_id), course_photo_id)
WHERE id = $1
RETURNING *;

-- name: ContentSectionPhotoDeleteById :exec
DELETE FROM content_section_photos
WHERE id = $1;

-- name: ContentSectionPhotoList :many
SELECT *
FROM content_section_photos
ORDER BY created_at DESC;

-- name: ContentSectionPhotoCount :one
SELECT COALESCE(COUNT(*), 0)::BIGINT AS content_section_photo_count
FROM content_section_photos;

-- name: ContentSectionPhotoListByContentSectionId :many
SELECT *
FROM content_section_photos
WHERE content_section_id = $1
ORDER BY created_at DESC;

-- name: ContentSectionPhotoListByCoursePhotoId :many
SELECT *
FROM content_section_photos
WHERE course_photo_id = $1
ORDER BY created_at DESC;