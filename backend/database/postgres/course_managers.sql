-- name: CourseManagerCreate :one
INSERT INTO course_managers (course_id, user_id)
VALUES ($1, $2)
RETURNING *;

-- name: CourseManagerGet :one
SELECT *
FROM course_managers
WHERE id = $1;

-- name: CourseManagerUpdate :one
UPDATE course_managers
SET
    course_id = COALESCE(sqlc.narg(course_id), course_id),
    user_id = COALESCE(sqlc.narg(user_id), user_id)
WHERE id = $1
RETURNING *;

-- name: CourseManagerDelete :one
DELETE FROM course_managers
WHERE id = $1
RETURNING *;

-- name: CourseManagerList :many
SELECT sqlc.embed(course_managers),
       sqlc.embed(courses),
       sqlc.embed(users)
FROM course_managers
LEFT JOIN courses ON course_managers.course_id = courses.id
LEFT JOIN users ON course_managers.user_id = users.id
WHERE (sqlc.narg(courseId)::BIGINT IS NULL OR course_managers.course_id = sqlc.narg(courseId)::BIGINT)
  AND (sqlc.narg(userId)::BIGINT IS NULL OR course_managers.user_id = sqlc.narg(userId)::BIGINT)
GROUP BY course_managers.id, courses.id, users.id
ORDER BY
  CASE WHEN sqlc.narg('sort') = 'createdAt' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN course_managers.created_at END,
  CASE WHEN sqlc.narg('sort') = 'createdAt' AND sqlc.narg('order') = 'desc' THEN course_managers.created_at END DESC,
  CASE WHEN sqlc.narg('sort') = 'updatedAt' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN course_managers.updated_at END,
  CASE WHEN sqlc.narg('sort') = 'updatedAt' AND sqlc.narg('order') = 'desc' THEN course_managers.updated_at END DESC
LIMIT sqlc.narg('limit')::BIGINT
OFFSET COALESCE(sqlc.narg('offset')::BIGINT, 0);

-- name: CourseManagerCount :one
SELECT COALESCE(COUNT(*), 0)::BIGINT AS course_manager_count
FROM course_managers
WHERE (sqlc.narg(courseId)::BIGINT IS NULL OR course_managers.course_id = sqlc.narg(courseId)::BIGINT)
  AND (sqlc.narg(userId)::BIGINT IS NULL OR course_managers.user_id = sqlc.narg(userId)::BIGINT);