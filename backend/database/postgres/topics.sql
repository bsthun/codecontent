-- name: TopicCreate :one
INSERT INTO topics (name, label, embedding_no)
VALUES ($1, $2, $3)
RETURNING *;

-- name: TopicGet :one
SELECT *
FROM topics
WHERE id = $1;

-- name: TopicUpdate :one
UPDATE topics
SET
    name = COALESCE(sqlc.narg(name), name),
    label = COALESCE(sqlc.narg(label), label),
    embedding_no = COALESCE(sqlc.narg(embedding_no), embedding_no)
WHERE id = $1
RETURNING *;

-- name: TopicDelete :one
DELETE FROM topics
WHERE id = $1
RETURNING *;

-- name: TopicList :many
SELECT *
FROM topics
WHERE (sqlc.narg(name)::TEXT IS NULL OR LOWER(topics.name) LIKE LOWER('%' || sqlc.narg(name) || '%'))
  AND (sqlc.narg(label)::TEXT IS NULL OR LOWER(topics.label) LIKE LOWER('%' || sqlc.narg(label) || '%'))
ORDER BY
  CASE WHEN sqlc.narg('sort') = 'name' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN topics.name END,
  CASE WHEN sqlc.narg('sort') = 'name' AND sqlc.narg('order') = 'desc' THEN topics.name END DESC,
  CASE WHEN sqlc.narg('sort') = 'createdAt' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN topics.created_at END,
  CASE WHEN sqlc.narg('sort') = 'createdAt' AND sqlc.narg('order') = 'desc' THEN topics.created_at END DESC,
  CASE WHEN sqlc.narg('sort') = 'updatedAt' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN topics.updated_at END,
  CASE WHEN sqlc.narg('sort') = 'updatedAt' AND sqlc.narg('order') = 'desc' THEN topics.updated_at END DESC
LIMIT sqlc.narg('limit')::BIGINT
OFFSET COALESCE(sqlc.narg('offset')::BIGINT, 0);

-- name: TopicCount :one
SELECT COALESCE(COUNT(*), 0)::BIGINT AS topic_count
FROM topics
WHERE (sqlc.narg(name)::TEXT IS NULL OR LOWER(topics.name) LIKE LOWER('%' || sqlc.narg(name) || '%'))
  AND (sqlc.narg(label)::TEXT IS NULL OR LOWER(topics.label) LIKE LOWER('%' || sqlc.narg(label) || '%'));