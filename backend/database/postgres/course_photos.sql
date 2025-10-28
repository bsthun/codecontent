-- name: CoursePhotoCreate :one
INSERT INTO course_photos (course_id, title, description)
VALUES ($1, $2, $3)
RETURNING *;

-- name: CoursePhotoGet :one
SELECT *
FROM course_photos
WHERE id = $1;

-- name: CoursePhotoUpdate :one
UPDATE course_photos
SET
    course_id = COALESCE(sqlc.narg(course_id), course_id),
    title = COALESCE(sqlc.narg(title), title),
    description = COALESCE(sqlc.narg(description), description)
WHERE id = $1
RETURNING *;

-- name: CoursePhotoDelete :one
DELETE FROM course_photos
WHERE id = $1
RETURNING *;

-- name: CoursePhotoList :many
SELECT course_photos.id, course_photos.course_id, course_photos.title, course_photos.created_at, course_photos.updated_at,
       courses.id, courses.name, courses.description, courses.prompt_instruction, courses.created_at, courses.updated_at,
       (SELECT COALESCE(COUNT(*), 0)::BIGINT FROM content_section_photos WHERE content_section_photos.course_photo_id = course_photos.id) AS content_section_photo_count
FROM course_photos
LEFT JOIN courses ON course_photos.course_id = courses.id
WHERE (sqlc.narg(course_id)::BIGINT IS NULL OR course_photos.course_id = sqlc.narg(course_id)::BIGINT)
  AND (sqlc.narg(title)::TEXT IS NULL OR LOWER(course_photos.title) LIKE LOWER('%' || sqlc.narg(title) || '%'))
GROUP BY course_photos.id, courses.id
ORDER BY
  CASE WHEN sqlc.narg('sort') = 'title' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN course_photos.title END,
  CASE WHEN sqlc.narg('sort') = 'title' AND sqlc.narg('order') = 'desc' THEN course_photos.title END DESC,
  CASE WHEN sqlc.narg('sort') = 'createdAt' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN course_photos.created_at END,
  CASE WHEN sqlc.narg('sort') = 'createdAt' AND sqlc.narg('order') = 'desc' THEN course_photos.created_at END DESC,
  CASE WHEN sqlc.narg('sort') = 'updatedAt' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN course_photos.updated_at END,
  CASE WHEN sqlc.narg('sort') = 'updatedAt' AND sqlc.narg('order') = 'desc' THEN course_photos.updated_at END DESC
LIMIT sqlc.narg('limit')::BIGINT
OFFSET COALESCE(sqlc.narg('offset')::BIGINT, 0);

-- name: CoursePhotoCount :one
SELECT COALESCE(COUNT(*), 0)::BIGINT AS course_photo_count
FROM course_photos
WHERE (sqlc.narg(course_id)::BIGINT IS NULL OR course_photos.course_id = sqlc.narg(course_id)::BIGINT)
  AND (sqlc.narg(title)::TEXT IS NULL OR LOWER(course_photos.title) LIKE LOWER('%' || sqlc.narg(title) || '%'));

