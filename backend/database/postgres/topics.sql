-- name: TopicCreate :one
INSERT INTO topics (name, label, embedding_no)
VALUES ($1, $2, $3)
RETURNING *;

-- name: TopicGetById :one
SELECT *
FROM topics
WHERE id = $1;

-- name: TopicUpdateById :one
UPDATE topics
SET
    name = COALESCE(sqlc.narg(name), name),
    label = COALESCE(sqlc.narg(label), label),
    embedding_no = COALESCE(sqlc.narg(embedding_no), embedding_no)
WHERE id = $1
RETURNING *;

-- name: TopicDeleteById :exec
DELETE FROM topics
WHERE id = $1;

-- name: TopicList :many
SELECT *
FROM topics
ORDER BY created_at DESC;

-- name: TopicCount :one
SELECT COALESCE(COUNT(*), 0)::BIGINT AS topic_count
FROM topics;