package services

import (
	"context"
	"fmt"

	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
	mappers "github.com/campushq-official/campushq-api/src/internal/core/infrastructure/postgres/mappers/personalinfo-mappers"
	studentMapper "github.com/campushq-official/campushq-api/src/internal/core/infrastructure/postgres/mappers/student-mappers"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s *studentService) StudentSignup(req *dtos.StudentCreateDTO) (int32, error) {
	ctx := context.TODO()
	lastInsertedId, err := s.repository.GetLastInsertedStudentId(ctx)
	fmt.Println("we here ", err)
	if err != nil {
		return 0, err
	}

	studentEntity := studentMapper.StudentSignupMapper(*req, lastInsertedId)
	insertedStudentId, err := s.repository.InsertStudent(ctx, *studentEntity)
	fmt.Println("we here 2", err)
	if err != nil {
		return 0, err
	}

	personalInfoEntity := mappers.PersonalInfoCreateMapper(*req, pgtype.Int4{Int32: insertedStudentId, Valid: true})

	if err = s.repository.InsertPersonalInfo(ctx, *personalInfoEntity); err != nil {
		fmt.Println("we here 3", err)
		return 0, err
	}

	return insertedStudentId, nil
}
