package dtos

type StudentCreateDTO struct {
	Username               string `json:"username"`
	Email                  string `json:"email"`
	Connection             string `json:"connection"`
	FirstName              string `json:"first_name"`
	LastName               string `json:"last_name"`
	Sex                    string `json:"sex"`
	Role                   string `json:"role"`
	Nationality            string `json:"nationality"`
	City                   string `json:"city"`
	Address                string `json:"address"`
	DOB                    string `json:"dob"`
	PlaceOfBirth           string `json:"place_of_birth"`
	FaherName              string `json:"father_name"`
	MotherName             string `json:"mother_name"`
	PassportNumber         string `json:"passport_number"`
	PassportIssueDate      string `json:"passport_issue_date"`
	PassportExpiryDate     string `json:"passport_expiry_date"`
	PassportIssueAuthority string `json:"passport_issue_authority"`
	DepartmentID           uint   `json:"department_id"`
}
