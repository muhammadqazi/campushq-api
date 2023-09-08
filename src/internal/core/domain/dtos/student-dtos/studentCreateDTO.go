package dtos

import "github.com/jackc/pgx/v5/pgtype"

type StudentCreateDTO struct {
	Username               string      `json:"username"`
	Email                  string      `json:"email"`
	Password               string      `json:"password"`
	FirstName              string      `json:"firstName"`
	Surname                string      `json:"surname"`
	Sex                    string      `json:"sex"`
	Role                   string      `json:"role"`
	Nationality            string      `json:"nationality"`
	City                   string      `json:"city"`
	Address                string      `json:"address"`
	DOB                    pgtype.Date `json:"dob"`
	PlaceOfBirth           string      `json:"placeOfBirth"`
	FatherName             string      `json:"fatherName"`
	MotherName             string      `json:"motherName"`
	GuardianPhoneNumber    string      `json:"guardianPhoneNumber"`
	PassportNumber         string      `json:"passportNumber"`
	PassportIssueDate      pgtype.Date `json:"passportIssueDate"`
	PassportExpiryDate     pgtype.Date `json:"passportExpiryDate"`
	PassportIssueAuthority string      `json:"passportIssueAuthority"`
	DepartmentID           pgtype.Int4 `json:"departmentId"`
}

type Auth0SignupDTO struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Connection string `json:"connection"`
}

type Auth0SignupResponseDTO struct {
	CreatedAt     string `json:"created_at"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Identities    []struct {
		Connection string `json:"connection"`
		UserId     string `json:"user_id"`
		Provider   string `json:"provider"`
		IsSocial   bool   `json:"isSocial"`
	} `json:"identities"`
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	Picture   string `json:"picture"`
	UpdatedAt string `json:"updated_at"`
	UserId    string `json:"user_id"`
	Username  string `json:"username"`
}

type Auth0AssignRoleDTO struct {
	Roles []string `json:"roles"`
}
