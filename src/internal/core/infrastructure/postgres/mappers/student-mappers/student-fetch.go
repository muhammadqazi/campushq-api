package mappers

import (
	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
	repositories "github.com/campushq-official/campushq-api/src/internal/core/infrastructure/postgres/repositories/sqlc"
)

func StudentFetchMapper(req repositories.Student) []*dtos.StudentFetchDTO {

	arr := []repositories.Student{req}

	var res []*dtos.StudentFetchDTO

	departmentId := int64(req.DepartmentID.Int32)
	supervisorId := int64(req.SupervisorID.Int32)

	graduationDate := req.GraduationDate.Time

	for _, v := range arr {
		res = append(res, &dtos.StudentFetchDTO{
			StudentID:      int64(v.StudentID),
			FirstName:      v.FirstName,
			Surname:        v.Surname,
			Sex:            v.Sex,
			Role:           v.Role,
			Status:         v.Status,
			AccessStatus:   v.AccessStatus,
			AcceptanceType: v.AcceptanceType,
			GraduationDate: graduationDate,
			DepartmentID:   departmentId,
			SupervisorID:   supervisorId,
			IsActive:       v.IsActive,
		})
	}

	return res
}
