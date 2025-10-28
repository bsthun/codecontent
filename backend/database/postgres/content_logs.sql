-- name: ContentLogCreate :one
INSERT INTO content_logs (content_id, prompt, call)
VALUES ($1, $2, $3)
RETURNING *;

-- name: ContentLogGetById :one
SELECT *
FROM content_logs
WHERE id = $1;

-- name: ContentLogUpdateById :one
UPDATE content_logs
SET
    content_id = COALESCE(sqlc.narg(content_id), content_id),
    prompt = COALESCE(sqlc.narg(prompt), prompt),
    call = COALESCE(sqlc.narg(call), call)
WHERE id = $1
RETURNING *;

-- name: ContentLogDeleteById :exec
DELETE FROM content_logs
WHERE id = $1;

-- name: ContentLogList :many
SELECT *
FROM content_logs
ORDER BY created_at DESC;

-- name: ContentLogCount :one
SELECT COALESCE(COUNT(*), 0)::BIGINT AS content_log_count
FROM content_logs;

-- name: ContentLogListByContentId :many
SELECT *
FROM content_logs
WHERE content_id = $1
ORDER BY created_at DESC;