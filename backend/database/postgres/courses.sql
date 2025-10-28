-- name: CourseCreate :one
INSERT INTO courses (name, description, prompt_instruction)
VALUES ($1, $2, $3)
RETURNING *;

-- name: CourseGet :one
SELECT *
FROM courses
WHERE id = $1;

-- name: CourseUpdate :one
UPDATE courses
SET
    name = COALESCE(sqlc.narg(name), name),
    description = COALESCE(sqlc.narg(description), description),
    prompt_instruction = COALESCE(sqlc.narg(prompt_instruction), prompt_instruction)
WHERE id = $1
RETURNING *;

-- name: CourseDelete :one
DELETE FROM courses
WHERE id = $1
RETURNING *;

-- name: CourseList :many
SELECT sqlc.embed(courses),
       (SELECT COALESCE(COUNT(*), 0)::BIGINT FROM course_managers WHERE course_managers.course_id = courses.id) AS course_manager_count,
       (SELECT COALESCE(COUNT(*), 0)::BIGINT FROM enrolls WHERE enrolls.course_id = courses.id) AS enroll_count,
       (SELECT COALESCE(COUNT(*), 0)::BIGINT FROM course_photos WHERE course_photos.course_id = courses.id) AS course_photo_count
FROM courses
WHERE (sqlc.narg(name)::TEXT IS NULL OR LOWER(courses.name) LIKE LOWER('%' || sqlc.narg(name) || '%'))
GROUP BY courses.id
ORDER BY
  CASE WHEN sqlc.narg('sort') = 'name' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN courses.name END,
  CASE WHEN sqlc.narg('sort') = 'name' AND sqlc.narg('order') = 'desc' THEN courses.name END DESC,
  CASE WHEN sqlc.narg('sort') = 'createdAt' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN courses.created_at END,
  CASE WHEN sqlc.narg('sort') = 'createdAt' AND sqlc.narg('order') = 'desc' THEN courses.created_at END DESC,
  CASE WHEN sqlc.narg('sort') = 'updatedAt' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN courses.updated_at END,
  CASE WHEN sqlc.narg('sort') = 'updatedAt' AND sqlc.narg('order') = 'desc' THEN courses.updated_at END DESC
LIMIT sqlc.narg('limit')::BIGINT
OFFSET COALESCE(sqlc.narg('offset')::BIGINT, 0);

-- name: CourseCount :one
SELECT COALESCE(COUNT(*), 0)::BIGINT AS course_count
FROM courses
WHERE (sqlc.narg(name)::TEXT IS NULL OR LOWER(courses.name) LIKE LOWER('%' || sqlc.narg(name) || '%'));

-- name: CourseListByManager :many
SELECT sqlc.embed(courses),
       (SELECT COALESCE(COUNT(*), 0)::BIGINT FROM course_managers WHERE course_managers.course_id = courses.id) AS course_manager_count,
       (SELECT COALESCE(COUNT(*), 0)::BIGINT FROM enrolls WHERE enrolls.course_id = courses.id) AS enroll_count,
       (SELECT COALESCE(COUNT(*), 0)::BIGINT FROM course_photos WHERE course_photos.course_id = courses.id) AS course_photo_count
FROM courses
JOIN course_managers ON courses.id = course_managers.course_id
WHERE (sqlc.narg(userId)::BIGINT IS NULL OR course_managers.user_id = sqlc.narg(userId)::BIGINT)
  AND (sqlc.narg(name)::TEXT IS NULL OR LOWER(courses.name) LIKE LOWER('%' || sqlc.narg(name) || '%'))
GROUP BY courses.id
ORDER BY
  CASE WHEN sqlc.narg('sort') = 'name' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN courses.name END,
  CASE WHEN sqlc.narg('sort') = 'name' AND sqlc.narg('order') = 'desc' THEN courses.name END DESC,
  CASE WHEN sqlc.narg('sort') = 'createdAt' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN courses.created_at END,
  CASE WHEN sqlc.narg('sort') = 'createdAt' AND sqlc.narg('order') = 'desc' THEN courses.created_at END DESC,
  CASE WHEN sqlc.narg('sort') = 'updatedAt' AND COALESCE(sqlc.narg('order'), 'asc') = 'asc' THEN courses.updated_at END,
  CASE WHEN sqlc.narg('sort') = 'updatedAt' AND sqlc.narg('order') = 'desc' THEN courses.updated_at END DESC
LIMIT sqlc.narg('limit')::BIGINT
OFFSET COALESCE(sqlc.narg('offset')::BIGINT, 0);

-- name: CourseCountByManager :one
SELECT COALESCE(COUNT(DISTINCT courses.id), 0)::BIGINT AS course_count
FROM courses
JOIN course_managers ON courses.id = course_managers.course_id
WHERE (sqlc.narg(userId)::BIGINT IS NULL OR course_managers.user_id = sqlc.narg(userId)::BIGINT)
  AND (sqlc.narg(name)::TEXT IS NULL OR LOWER(courses.name) LIKE LOWER('%' || sqlc.narg(name) || '%'));