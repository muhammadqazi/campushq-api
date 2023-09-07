package mappers

import (
	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
	repositories "github.com/campushq-official/campushq-api/src/internal/core/infrastructure/postgres/repositories/sqlc"
)

func FetchAllStudentMapper(req []repositories.Student) []*dtos.StudentFetchDTO {
	var result []*dtos.StudentFetchDTO

	for _, student := range req {
		departmentID := int64(student.DepartmentID.Int32)
		supervisorID := int64(student.SupervisorID.Int32)
		graduationDate := student.GraduationDate.Time

		studentDTO := &dtos.StudentFetchDTO{
			StudentID:      int64(student.StudentID),
			FirstName:      student.FirstName,
			Surname:        student.Surname,
			Sex:            student.Sex,
			Role:           student.Role,
			Status:         student.Status,
			AccessStatus:   student.AccessStatus,
			AcceptanceType: student.AcceptanceType,
			GraduationDate: graduationDate,
			DepartmentID:   departmentID,
			SupervisorID:   supervisorID,
			IsActive:       student.IsActive,
		}

		result = append(result, studentDTO)
	}

	return result
}
