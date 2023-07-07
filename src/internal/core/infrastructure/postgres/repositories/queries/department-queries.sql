-- name: GetDepartmentById :one
SELECT * FROM departments WHERE department_id = $1;