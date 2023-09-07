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
    student_id
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
    $13
);

-- name: UpdateStudentPersonalInfo :exec
UPDATE personal_information
SET 
    nationality = COALESCE(sqlc.narg(nationality), nationality),
    city = COALESCE(sqlc.narg(city), city),
    address = COALESCE(sqlc.narg(address), address),
    dob = COALESCE(sqlc.narg(dob), dob),
    place_of_birth = COALESCE(sqlc.narg(place_of_birth), place_of_birth),
    local_address = COALESCE(sqlc.narg(local_address), local_address),
    father_name = COALESCE(sqlc.narg(father_name), father_name),
    mother_name = COALESCE(sqlc.narg(mother_name), mother_name),
    guardian_phone_number = COALESCE(sqlc.narg(guardian_phone_number), guardian_phone_number),
    local_phone_number = COALESCE(sqlc.narg(local_phone_number), local_phone_number),
    passport_number = COALESCE(sqlc.narg(passport_number), passport_number),
    passport_issue_date = COALESCE(sqlc.narg(passport_issue_date), passport_issue_date),
    passport_expiry_date = COALESCE(sqlc.narg(passport_expiry_date), passport_expiry_date),
    passport_issuing_authority = COALESCE(sqlc.narg(passport_issuing_authority), passport_issuing_authority),
    id_card_number = COALESCE(sqlc.narg(id_card_number), id_card_number),
    updated_at = CURRENT_TIMESTAMP
WHERE
    student_id = $1;

    