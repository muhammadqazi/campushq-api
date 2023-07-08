package services

import (
	"context"

	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
	mappers "github.com/campushq-official/campushq-api/src/internal/core/infrastructure/postgres/mappers/student-mappers"
)

func (s *studentService) StudentModifyByID(req *dtos.StudentPatchDTO, id string) error {
	ctx := context.TODO()

	studentEntity := mappers.StudentPatchMapper(*req, id)
	err := s.repository.UpdateStudent(ctx, *studentEntity)
	if err != nil {
		return err
	}

	return nil
}
