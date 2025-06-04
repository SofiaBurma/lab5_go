-- name: CreateStudents :one
INSERT INTO students(
    full_name, age, group_name
) VALUES (
    $1, $2, $3
) RETURNING *;