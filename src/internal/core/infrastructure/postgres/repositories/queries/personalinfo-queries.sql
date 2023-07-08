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

-- name: UpdatePersonalInfo :exec
UPDATE personal_information
SET 
    nationality = COALESCE($2, nationality),
    city = COALESCE($3, city),
    address = COALESCE($4, address),
    dob = COALESCE($5, dob),
    place_of_birth = COALESCE($6, place_of_birth),
    local_address = COALESCE($7, local_address),
    father_name = COALESCE($7, father_name),
    mother_name = COALESCE($8, mother_name),
    guardian_phone_number = COALESCE($9, guardian_phone_number),
    local_phone_number = COALESCE($10, local_phone_number),
    passport_number = COALESCE($11, passport_number),
    passport_issue_date = COALESCE($12, passport_issue_date),
    passport_expiry_date = COALESCE($13, passport_expiry_date),
    passport_issuing_authority = COALESCE($14, passport_issuing_authority),
    id_card_number = COALESCE($15, id_card_number),
    student_id = COALESCE($16, student_id),
    staff_id = COALESCE($17, staff_id),
    updated_at = CURRENT_TIMESTAMP
WHERE pid = $1;

    