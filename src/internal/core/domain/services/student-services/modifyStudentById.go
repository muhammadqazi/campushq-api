package services

import (
	"context"

	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
	personalInfo "github.com/campushq-official/campushq-api/src/internal/core/infrastructure/postgres/mappers/personalinfo-mappers"
	mappers "github.com/campushq-official/campushq-api/src/internal/core/infrastructure/postgres/mappers/student-mappers"
)

func (s *studentService) ModifyStudentById(req *dtos.StudentPatchDTO, id string) error {
	ctx := context.TODO()

	studentEntity := mappers.PatchStudentMapper(*req, id)
	err := s.repository.UpdateStudent(ctx, *studentEntity)
	if err != nil {
		return err
	}

	personalInfoEntity := personalInfo.StudentPersonalInfoPatchMapper(*req, id)
	err = s.repository.UpdateStudentPersonalInfo(ctx, *personalInfoEntity)
	if err != nil {
		return err
	}

	return nil
}
