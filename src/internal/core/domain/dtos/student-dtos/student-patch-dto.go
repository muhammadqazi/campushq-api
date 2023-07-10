package dtos

import "github.com/jackc/pgx/v5/pgtype"

type StudentPatchDTO struct {
	FirstName                string      `json:"firstName,omitempty"`
	Surname                  string      `json:"surname,omitempty"`
	Role                     string      `json:"role,omitempty"`
	Status                   string      `json:"status,omitempty"`
	AccessStatus             string      `json:"accessStatus,omitempty"`
	DepartmentID             pgtype.Int4 `json:"departmentId,omitempty"`
	SupervisorID             pgtype.Int4 `json:"supervisorId,omitempty"`
	Nationality              string      `json:"nationality,omitempty"`
	City                     string      `json:"city,omitempty"`
	Address                  string      `json:"address,omitempty"`
	Dob                      pgtype.Date `json:"dob,omitempty"`
	PlaceOfBirth             string      `json:"place_of_birth,omitempty"`
	LocalAddress             string      `json:"local_address,omitempty"`
	FatherName               string      `json:"father_name,omitempty"`
	MotherName               string      `json:"mother_name,omitempty"`
	GuardianPhoneNumber      string      `json:"guardian_phone_number,omitempty"`
	LocalPhoneNumber         string      `json:"local_phone_number,omitempty"`
	PassportNumber           string      `json:"passport_number,omitempty"`
	PassportIssueDate        pgtype.Date `json:"passport_issue_date,omitempty"`
	PassportExpiryDate       pgtype.Date `json:"passport_expiry_date,omitempty"`
	PassportIssuingAuthority string      `json:"passport_issuing_authority,omitempty"`
	IDCardNumber             string      `json:"id_card_number,omitempty"`
}
