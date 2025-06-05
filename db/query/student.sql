-- name: CreateStudents :one
INSERT INTO students(
    full_name, age, group_name
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: ListStudents :many
SELECT *
FROM students
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: GetStudentByID :one
SELECT *
FROM students
WHERE id = $1;

-- name: UpdateStudent :one
UPDATE students
SET full_name = $2,
    age = $3,
    group_name = $4
WHERE id = $1
RETURNING *;

-- name: DeleteStudent :exec
DELETE FROM students
WHERE id = $1;
