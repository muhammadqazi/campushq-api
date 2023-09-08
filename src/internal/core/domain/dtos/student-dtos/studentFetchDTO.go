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

type PaginationMetadata struct {
	Limit      int32  `json:"limit"`
	From       int32  `json:"from"`
	To         int32  `json:"to"`
	TotalCount int32  `json:"totalCount"`
	Links      []Link `json:"links"`
}

type Link struct {
	Self     interface{} `json:"self"`
	First    interface{} `json:"first"`
	Previous interface{} `json:"previous"`
	Next     interface{} `json:"next"`
}
