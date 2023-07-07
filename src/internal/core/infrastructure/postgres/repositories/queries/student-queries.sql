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
    semester,
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
    $9,
    $10
) RETURNING student_id;

-- name: GetLastInsertedStudentId :one
SELECT student_id FROM students ORDER BY student_id DESC LIMIT 1;
