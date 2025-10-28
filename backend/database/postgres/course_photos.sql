-- name: CoursePhotoCreate :one
INSERT INTO course_photos (course_id, title, description)
VALUES ($1, $2, $3)
RETURNING *;

-- name: CoursePhotoGetById :one
SELECT *
FROM course_photos
WHERE id = $1;

-- name: CoursePhotoUpdateById :one
UPDATE course_photos
SET
    course_id = COALESCE(sqlc.narg(course_id), course_id),
    title = COALESCE(sqlc.narg(title), title),
    description = COALESCE(sqlc.narg(description), description)
WHERE id = $1
RETURNING *;

-- name: CoursePhotoDeleteById :exec
DELETE FROM course_photos
WHERE id = $1;

-- name: CoursePhotoList :many
SELECT *
FROM course_photos
ORDER BY created_at DESC;

-- name: CoursePhotoCount :one
SELECT COALESCE(COUNT(*), 0)::BIGINT AS course_photo_count
FROM course_photos;

-- name: CoursePhotoListByCourseId :many
SELECT *
FROM course_photos
WHERE course_id = $1
ORDER BY created_at DESC;