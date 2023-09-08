package services

import (
	"context"

	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
	mappers "github.com/campushq-official/campushq-api/src/internal/core/infrastructure/postgres/mappers/personalinfo-mappers"
	studentMapper "github.com/campushq-official/campushq-api/src/internal/core/infrastructure/postgres/mappers/student-mappers"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s *studentService) StudentRegister(req *dtos.StudentCreateDTO) (int32, error) {
	ctx := context.TODO()
	lastInsertedId, err := s.repository.SelectLastInsertedStudentId(ctx)
	if err != nil {
		return 0, err
	}

	studentEntity := studentMapper.SignupStudentMapper(*req, lastInsertedId)
	insertedStudentId, err := s.repository.InsertStudent(ctx, *studentEntity)
	if err != nil {
		return 0, err
	}

	personalInfoEntity := mappers.StudentPersonalInfoCreateMapper(*req, pgtype.Int4{Int32: insertedStudentId, Valid: true})

	if err = s.repository.InsertPersonalInfo(ctx, *personalInfoEntity); err != nil {
		return 0, err
	}

	return insertedStudentId, nil
}
