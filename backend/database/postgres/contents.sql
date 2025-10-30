-- name: ContentCreate :one
INSERT INTO contents (enroll_id, title)
VALUES ($1, $2)
RETURNING *;

-- name: ContentGet :one
SELECT *
FROM contents
WHERE id = $1;

-- name: ContentUpdate :one
UPDATE contents
SET
    enroll_id = COALESCE(sqlc.narg(enroll_id), enroll_id),
    title = COALESCE(sqlc.narg(title), title)
WHERE id = $1
RETURNING *;

-- name: ContentDelete :one
DELETE FROM contents
WHERE id = $1
RETURNING *;

-- name: ContentList :many
SELECT sqlc.embed(contents),
       sqlc.embed(enrolls),
       (SELECT COALESCE(COUNT(*), 0)::BIGINT FROM content_sections WHERE content_sections.content_id = contents.id) AS content_section_count,
       (SELECT COALESCE(COUNT(*), 0)::BIGINT FROM content_logs WHERE content_logs.content_id = contents.id) AS content_log_count
FROM contents
LEFT JOIN enrolls ON contents.enroll_id = enrolls.id
WHERE (sqlc.narg(title)::TEXT IS NULL OR LOWER(contents.title) LIKE LOWER('%' || sqlc.narg(title) || '%'))
  AND (sqlc.narg(enrollId)::BIGINT IS NULL OR contents.enroll_id = sqlc.narg(enrollId)::BIGINT)
GROUP BY contents.id, enrolls.id
ORDER BY
  CASE WHEN sqlc.narg('sort') = 'title' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN contents.title END,
  CASE WHEN sqlc.narg('sort') = 'title' AND sqlc.narg('order') = 'desc' THEN contents.title END DESC,
  CASE WHEN sqlc.narg('sort') = 'createdAt' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN contents.created_at END,
  CASE WHEN sqlc.narg('sort') = 'createdAt' AND sqlc.narg('order') = 'desc' THEN contents.created_at END DESC,
  CASE WHEN sqlc.narg('sort') = 'updatedAt' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN contents.updated_at END,
  CASE WHEN sqlc.narg('sort') = 'updatedAt' AND sqlc.narg('order') = 'desc' THEN contents.updated_at END DESC
LIMIT sqlc.narg('limit')::BIGINT
OFFSET COALESCE(sqlc.narg('offset')::BIGINT, 0);

-- name: ContentCount :one
SELECT COALESCE(COUNT(*), 0)::BIGINT AS content_count
FROM contents
WHERE (sqlc.narg(title)::TEXT IS NULL OR LOWER(contents.title) LIKE LOWER('%' || sqlc.narg(title) || '%'))
  AND (sqlc.narg(enrollId)::BIGINT IS NULL OR contents.enroll_id = sqlc.narg(enrollId)::BIGINT);

-- name: ContentListByCourse :many
SELECT sqlc.embed(contents),
       sqlc.embed(enrolls),
       (SELECT COALESCE(COUNT(*), 0)::BIGINT FROM content_sections WHERE content_sections.content_id = contents.id) AS content_section_count,
       (SELECT COALESCE(COUNT(*), 0)::BIGINT FROM content_logs WHERE content_logs.content_id = contents.id) AS content_log_count
FROM contents
LEFT JOIN enrolls ON contents.enroll_id = enrolls.id
WHERE (sqlc.narg(courseId)::BIGINT IS NULL OR enrolls.course_id = sqlc.narg(courseId)::BIGINT)
  AND (sqlc.narg(userId)::BIGINT IS NULL OR enrolls.user_id = sqlc.narg(userId)::BIGINT)
  AND (sqlc.narg(title)::TEXT IS NULL OR LOWER(contents.title) LIKE LOWER('%' || sqlc.narg(title) || '%'))
GROUP BY contents.id, enrolls.id
ORDER BY
  CASE WHEN sqlc.narg('sort') = 'title' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN contents.title END,
  CASE WHEN sqlc.narg('sort') = 'title' AND sqlc.narg('order') = 'desc' THEN contents.title END DESC,
  CASE WHEN sqlc.narg('sort') = 'createdAt' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN contents.created_at END,
  CASE WHEN sqlc.narg('sort') = 'createdAt' AND sqlc.narg('order') = 'desc' THEN contents.created_at END DESC,
  CASE WHEN sqlc.narg('sort') = 'updatedAt' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN contents.updated_at END,
  CASE WHEN sqlc.narg('sort') = 'updatedAt' AND sqlc.narg('order') = 'desc' THEN contents.updated_at END DESC
LIMIT sqlc.narg('limit')::BIGINT
OFFSET COALESCE(sqlc.narg('offset')::BIGINT, 0);

-- name: ContentCountByCourse :one
SELECT COALESCE(COUNT(*), 0)::BIGINT AS content_count
FROM contents
LEFT JOIN enrolls ON contents.enroll_id = enrolls.id
WHERE (sqlc.narg(courseId)::BIGINT IS NULL OR enrolls.course_id = sqlc.narg(courseId)::BIGINT)
  AND (sqlc.narg(userId)::BIGINT IS NULL OR enrolls.user_id = sqlc.narg(userId)::BIGINT)
  AND (sqlc.narg(title)::TEXT IS NULL OR LOWER(contents.title) LIKE LOWER('%' || sqlc.narg(title) || '%'));

