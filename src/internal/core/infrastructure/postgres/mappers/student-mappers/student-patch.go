package mappers

import (
	"strconv"

	"github.com/campushq-official/campushq-api/src/internal/common/utils"
	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
	repositories "github.com/campushq-official/campushq-api/src/internal/core/infrastructure/postgres/repositories/sqlc"
)

func StudentPatchMapper(req dtos.StudentPatchDTO, id string) *repositories.UpdateStudentParams {
	studentId, _ := strconv.ParseInt(id, 10, 32)
	studentEntity := repositories.UpdateStudentParams{
		StudentID:    int32(studentId),
		FirstName:    utils.GetNullableText(req.FirstName),
		Surname:      utils.GetNullableText(req.Surname),
		Role:         utils.GetNullableText(req.Role),
		Status:       utils.GetNullableText(req.Status),
		AccessStatus: utils.GetNullableText(req.AccessStatus),
		DepartmentID: utils.GetNullableInt(req.DepartmentID),
		SupervisorID: utils.GetNullableInt(req.SupervisorID),
	}
	return &studentEntity
}
