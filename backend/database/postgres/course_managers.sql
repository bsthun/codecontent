-- name: CourseManagerCreate :one
INSERT INTO course_managers (course_id, user_id)
VALUES ($1, $2)
RETURNING *;

-- name: CourseManagerGetById :one
SELECT *
FROM course_managers
WHERE id = $1;

-- name: CourseManagerUpdateById :one
UPDATE course_managers
SET
    course_id = COALESCE(sqlc.narg(course_id), course_id),
    user_id = COALESCE(sqlc.narg(user_id), user_id)
WHERE id = $1
RETURNING *;

-- name: CourseManagerDeleteById :exec
DELETE FROM course_managers
WHERE id = $1;

-- name: CourseManagerList :many
SELECT *
FROM course_managers
ORDER BY created_at DESC;

-- name: CourseManagerCount :one
SELECT COALESCE(COUNT(*), 0)::BIGINT AS course_manager_count
FROM course_managers;

-- name: CourseManagerListByCourseId :many
SELECT *
FROM course_managers
WHERE course_id = $1
ORDER BY created_at DESC;

-- name: CourseManagerListByUserId :many
SELECT *
FROM course_managers
WHERE user_id = $1
ORDER BY created_at DESC;