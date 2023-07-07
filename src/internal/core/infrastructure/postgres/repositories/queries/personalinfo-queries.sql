-- name: InsertPersonalInfo :exec
INSERT INTO personal_information (
    nationality,
    city,
    address,
    dob,
    place_of_birth,
    father_name,
    mother_name,
    guardian_phone_number,
    passport_number,
    passport_issue_date,
    passport_expiry_date,
    passport_issuing_authority,
    student_id,
    is_staff
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
    $10,
    $11,
    $12,
    $13,
    $14
);