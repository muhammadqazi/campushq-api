package dtos

type StudentPatchDTO struct {
	FirstName    string `json:"firstName,omitempty"`
	Surname      string `json:"surname,omitempty"`
	Role         string `json:"role,omitempty"`
	Status       string `json:"status,omitempty"`
	AccessStatus string `json:"accessStatus,omitempty"`
	DepartmentID int    `json:"departmentId,omitempty"`
	SupervisorID int    `json:"supervisorId,omitempty"`
}
