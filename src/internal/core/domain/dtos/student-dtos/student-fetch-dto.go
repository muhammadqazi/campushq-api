package dtos

import "time"

type StudentFetchDTO struct {
	StudentID      int64     `json:"studentId"`
	FirstName      string    `json:"firstName"`
	Surname        string    `json:"surname"`
	Sex            string    `json:"sex"`
	Role           string    `json:"role"`
	Status         string    `json:"status"`
	AccessStatus   string    `json:"accessStatus"`
	AcceptanceType string    `json:"acceptanceType"`
	GraduationDate time.Time `json:"graduationDate"`
	DepartmentID   int64     `json:"departmentId"`
	SupervisorID   int64     `json:"supervisorId"`
	IsActive       bool      `json:"isActive"`
}
