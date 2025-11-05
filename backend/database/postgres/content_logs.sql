-- name: ContentLogCreate :one
INSERT INTO content_logs (content_id, action, prompt, call)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: ContentLogGet :one
SELECT *
FROM content_logs
WHERE id = $1;

-- name: ContentLogUpdate :one
UPDATE content_logs
SET
    content_id = COALESCE(sqlc.narg(content_id), content_id),
    prompt = COALESCE(sqlc.narg(prompt), prompt),
    call = COALESCE(sqlc.narg(call), call)
WHERE id = $1
RETURNING *;

-- name: ContentLogDelete :one
DELETE FROM content_logs
WHERE id = $1
RETURNING *;

-- name: ContentLogList :many
SELECT content_logs.id, content_logs.content_id, content_logs.created_at, content_logs.updated_at,
       contents.id, contents.enroll_id, contents.title, contents.created_at, contents.updated_at
FROM content_logs
LEFT JOIN contents ON content_logs.content_id = contents.id
WHERE (sqlc.narg(content_id)::BIGINT IS NULL OR content_logs.content_id = sqlc.narg(content_id)::BIGINT)
GROUP BY content_logs.id, contents.id
ORDER BY
  CASE WHEN sqlc.narg('sort') = 'createdAt' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN content_logs.created_at END,
  CASE WHEN sqlc.narg('sort') = 'createdAt' AND sqlc.narg('order') = 'desc' THEN content_logs.created_at END DESC,
  CASE WHEN sqlc.narg('sort') = 'updatedAt' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN content_logs.updated_at END,
  CASE WHEN sqlc.narg('sort') = 'updatedAt' AND sqlc.narg('order') = 'desc' THEN content_logs.updated_at END DESC
LIMIT sqlc.narg('limit')::BIGINT
OFFSET COALESCE(sqlc.narg('offset')::BIGINT, 0);

-- name: ContentLogCount :one
SELECT COALESCE(COUNT(*), 0)::BIGINT AS content_log_count
FROM content_logs
WHERE (sqlc.narg(content_id)::BIGINT IS NULL OR content_logs.content_id = sqlc.narg(content_id)::BIGINT);

