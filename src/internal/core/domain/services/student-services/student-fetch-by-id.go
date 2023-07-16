package services

import (
	"context"
	"strconv"

	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
	mappers "github.com/campushq-official/campushq-api/src/internal/core/infrastructure/postgres/mappers/student-mappers"
)

func (s *studentService) StudentFetchByID(id string) ([]*dtos.StudentFetchDTO, error) {

	ctx := context.TODO()

	studentId, _ := strconv.ParseInt(id, 10, 32)

	studentEntity, err := s.repository.SelectStudentById(ctx, int32(studentId))
	if err != nil {
		return nil, err
	}

	return mappers.StudentFetchMapper(studentEntity), nil

}
