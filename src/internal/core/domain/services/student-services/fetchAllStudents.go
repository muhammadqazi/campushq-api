package services

import (
	"context"

	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
	mappers "github.com/campushq-official/campushq-api/src/internal/core/infrastructure/postgres/mappers/student-mappers"
	repositories "github.com/campushq-official/campushq-api/src/internal/core/infrastructure/postgres/repositories/sqlc"
)

func (s *studentService) FetchAllStudents(limit int64, offset int64) ([]*dtos.StudentFetchDTO, error) {

	ctx := context.TODO()

	students, err := s.repository.SelectAllStudents(ctx,
		repositories.SelectAllStudentsParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		})
	if err != nil {
		return nil, err
	}

	return mappers.FetchAllStudentMapper(students), nil

}
