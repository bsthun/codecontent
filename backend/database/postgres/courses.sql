-- name: CourseCreate :one
INSERT INTO courses (name, description, prompt_instruction)
VALUES ($1, $2, $3)
RETURNING *;

-- name: CourseGetById :one
SELECT *
FROM courses
WHERE id = $1;

-- name: CourseUpdateById :one
UPDATE courses
SET
    name = COALESCE(sqlc.narg(name), name),
    description = COALESCE(sqlc.narg(description), description),
    prompt_instruction = COALESCE(sqlc.narg(prompt_instruction), prompt_instruction)
WHERE id = $1
RETURNING *;

-- name: CourseDeleteById :exec
DELETE FROM courses
WHERE id = $1;

-- name: CourseList :many
SELECT *
FROM courses
ORDER BY created_at DESC;

-- name: CourseCount :one
SELECT COALESCE(COUNT(*), 0)::BIGINT AS course_count
FROM courses;