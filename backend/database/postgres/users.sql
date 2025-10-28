-- name: UserGet :one
SELECT *
FROM users
WHERE id = $1;

-- name: UserGetByOid :one
SELECT *
FROM users
WHERE oid = $1;

-- name: UserList :many
SELECT users.id, users.oid, users.firstname, users.lastname, users.email, users.picture_url, users.is_admin, users.created_at, users.updated_at,
       (SELECT COALESCE(COUNT(*), 0)::BIGINT FROM course_managers WHERE course_managers.user_id = users.id) AS course_manager_count,
       (SELECT COALESCE(COUNT(*), 0)::BIGINT FROM enrolls WHERE enrolls.user_id = users.id) AS enroll_count
FROM users
WHERE (sqlc.narg(firstname)::TEXT IS NULL OR LOWER(users.firstname) LIKE LOWER('%' || sqlc.narg(firstname) || '%'))
  AND (sqlc.narg(lastname)::TEXT IS NULL OR LOWER(users.lastname) LIKE LOWER('%' || sqlc.narg(lastname) || '%'))
  AND (sqlc.narg(email)::TEXT IS NULL OR LOWER(users.email) LIKE LOWER('%' || sqlc.narg(email) || '%'))
GROUP BY users.id
ORDER BY
  CASE WHEN sqlc.narg('sort') = 'firstname' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN users.firstname END,
  CASE WHEN sqlc.narg('sort') = 'firstname' AND sqlc.narg('order') = 'desc' THEN users.firstname END DESC,
  CASE WHEN sqlc.narg('sort') = 'lastname' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN users.lastname END,
  CASE WHEN sqlc.narg('sort') = 'lastname' AND sqlc.narg('order') = 'desc' THEN users.lastname END DESC,
  CASE WHEN sqlc.narg('sort') = 'email' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN users.email END,
  CASE WHEN sqlc.narg('sort') = 'email' AND sqlc.narg('order') = 'desc' THEN users.email END DESC,
  CASE WHEN sqlc.narg('sort') = 'createdAt' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN users.created_at END,
  CASE WHEN sqlc.narg('sort') = 'createdAt' AND sqlc.narg('order') = 'desc' THEN users.created_at END DESC,
  CASE WHEN sqlc.narg('sort') = 'updatedAt' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN users.updated_at END,
  CASE WHEN sqlc.narg('sort') = 'updatedAt' AND sqlc.narg('order') = 'desc' THEN users.updated_at END DESC
LIMIT sqlc.narg('limit')::BIGINT
OFFSET COALESCE(sqlc.narg('offset')::BIGINT, 0);

-- name: UserGetByMetadataUsername :one
SELECT *
FROM users
WHERE metadata->'credential'->>'username' = sqlc.narg(username)::text;

-- name: UserCreate :one
INSERT INTO users (oid, firstname, lastname, email, picture_url, is_admin, metadata)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UserUpdateMetadata :one
UPDATE users
SET metadata = $2
WHERE id = $1
RETURNING *;
