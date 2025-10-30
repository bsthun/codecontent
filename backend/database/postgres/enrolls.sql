-- name: EnrollCreate :one
INSERT INTO enrolls (course_id, user_id)
VALUES ($1, $2)
RETURNING *;

-- name: EnrollGet :one
SELECT *
FROM enrolls
WHERE id = $1;

-- name: EnrollUpdate :one
UPDATE enrolls
SET
    course_id = COALESCE(sqlc.narg(course_id), course_id),
    user_id = COALESCE(sqlc.narg(user_id), user_id)
WHERE id = $1
RETURNING *;

-- name: EnrollDelete :one
DELETE FROM enrolls
WHERE id = $1
RETURNING *;

-- name: EnrollList :many
SELECT sqlc.embed(enrolls),
       sqlc.embed(courses),
       sqlc.embed(users),
       (SELECT COALESCE(COUNT(*), 0)::BIGINT FROM contents WHERE contents.enroll_id = enrolls.id) AS content_count
FROM enrolls
LEFT JOIN courses ON enrolls.course_id = courses.id
LEFT JOIN users ON enrolls.user_id = users.id
WHERE (sqlc.narg(course_id)::BIGINT IS NULL OR enrolls.course_id = sqlc.narg(course_id)::BIGINT)
  AND (sqlc.narg(user_id)::BIGINT IS NULL OR enrolls.user_id = sqlc.narg(user_id)::BIGINT)
GROUP BY enrolls.id, courses.id, users.id
ORDER BY
  CASE WHEN sqlc.narg('sort') = 'createdAt' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN enrolls.created_at END,
  CASE WHEN sqlc.narg('sort') = 'createdAt' AND sqlc.narg('order') = 'desc' THEN enrolls.created_at END DESC,
  CASE WHEN sqlc.narg('sort') = 'updatedAt' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN enrolls.updated_at END,
  CASE WHEN sqlc.narg('sort') = 'updatedAt' AND sqlc.narg('order') = 'desc' THEN enrolls.updated_at END DESC
LIMIT sqlc.narg('limit')::BIGINT
OFFSET COALESCE(sqlc.narg('offset')::BIGINT, 0);

-- name: EnrollCount :one
SELECT COALESCE(COUNT(*), 0)::BIGINT AS enroll_count
FROM enrolls
WHERE (sqlc.narg(course_id)::BIGINT IS NULL OR enrolls.course_id = sqlc.narg(course_id)::BIGINT)
  AND (sqlc.narg(user_id)::BIGINT IS NULL OR enrolls.user_id = sqlc.narg(user_id)::BIGINT);

-- name: EnrollGetOrCreate :one
INSERT INTO enrolls (course_id, user_id)
VALUES ($1, $2)
ON CONFLICT (course_id, user_id) DO UPDATE SET
    updated_at = CURRENT_TIMESTAMP
RETURNING *;

