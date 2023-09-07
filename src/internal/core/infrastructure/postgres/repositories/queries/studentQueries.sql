-- name: InsertStudent :one
INSERT INTO students (
    student_id,
    first_name,
    surname,
    sex,
    role,
    status,
    access_status,
    acceptance_type,
    department_id
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9
) RETURNING student_id;

-- name: SelectLastInsertedStudentId :one
SELECT student_id FROM students ORDER BY student_id DESC LIMIT 1;

-- name: SelectStudentById :one
SELECT * FROM students WHERE student_id = $1;

-- name: UpdateStudent :exec
UPDATE 
    students
SET
    first_name = COALESCE(sqlc.narg(first_name), first_name),
    surname = COALESCE(sqlc.narg(surname), surname),
    role = COALESCE(sqlc.narg(role), role),
    status = COALESCE(sqlc.narg(status), status),
    access_status = COALESCE(sqlc.narg(access_status), access_status),
    department_id = COALESCE(sqlc.narg(department_id), department_id),
    supervisor_id = COALESCE(sqlc.narg(supervisor_id), supervisor_id),
    updated_at = CURRENT_TIMESTAMP
WHERE
    student_id = $1;

-- name: SelectAllStudents :many
SELECT * FROM students ORDER BY student_id DESC LIMIT $1 OFFSET $2;