-- name: EnrollCreate :one
INSERT INTO enrolls (course_id, user_id)
VALUES ($1, $2)
RETURNING *;

-- name: EnrollGetById :one
SELECT *
FROM enrolls
WHERE id = $1;

-- name: EnrollUpdateById :one
UPDATE enrolls
SET
    course_id = COALESCE(sqlc.narg(course_id), course_id),
    user_id = COALESCE(sqlc.narg(user_id), user_id)
WHERE id = $1
RETURNING *;

-- name: EnrollDeleteById :exec
DELETE FROM enrolls
WHERE id = $1;

-- name: EnrollList :many
SELECT *
FROM enrolls
ORDER BY created_at DESC;

-- name: EnrollCount :one
SELECT COALESCE(COUNT(*), 0)::BIGINT AS enroll_count
FROM enrolls;

-- name: EnrollListByCourseId :many
SELECT *
FROM enrolls
WHERE course_id = $1
ORDER BY created_at DESC;

-- name: EnrollListByUserId :many
SELECT *
FROM enrolls
WHERE user_id = $1
ORDER BY created_at DESC;