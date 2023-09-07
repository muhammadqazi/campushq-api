package mappers

import (
	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
	repositories "github.com/campushq-official/campushq-api/src/internal/core/infrastructure/postgres/repositories/sqlc"
)

func SignupStudentMapper(req dtos.StudentCreateDTO, lastInsertedId int32) *repositories.InsertStudentParams {
	// semester := utils.GetCurrentSemester()
	// studentId := utils.StudentIDGenerator(12, "Fall", lastInsertedId)

	return &repositories.InsertStudentParams{
		StudentID:    lastInsertedId + 1,
		FirstName:    req.FirstName,
		Surname:      req.Surname,
		Sex:          req.Sex,
		Role:         req.Role,
		Status:       "pending",
		AccessStatus: "pending",
		DepartmentID: req.DepartmentID,
	}
}
