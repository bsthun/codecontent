-- name: ContentCreate :one
INSERT INTO contents (enroll_id, title)
VALUES ($1, $2)
RETURNING *;

-- name: ContentGetById :one
SELECT *
FROM contents
WHERE id = $1;

-- name: ContentUpdateById :one
UPDATE contents
SET
    enroll_id = COALESCE(sqlc.narg(enroll_id), enroll_id),
    title = COALESCE(sqlc.narg(title), title)
WHERE id = $1
RETURNING *;

-- name: ContentDeleteById :exec
DELETE FROM contents
WHERE id = $1;

-- name: ContentList :many
SELECT *
FROM contents
ORDER BY created_at DESC;

-- name: ContentCount :one
SELECT COALESCE(COUNT(*), 0)::BIGINT AS content_count
FROM contents;

-- name: ContentListByEnrollId :many
SELECT *
FROM contents
WHERE enroll_id = $1
ORDER BY created_at DESC;