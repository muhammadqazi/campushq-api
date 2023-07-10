package services

import (
	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
	repositories "github.com/campushq-official/campushq-api/src/internal/core/infrastructure/postgres/repositories/sqlc"
)

type StudentService interface {
	StudentSignup(*dtos.StudentCreateDTO) (int32, error)
	StudentModifyByID(*dtos.StudentPatchDTO, string) error
}

type studentService struct {
	repository repositories.Store
}

func NewStudentService(repository repositories.Store) StudentService {
	return &studentService{
		repository: repository,
	}
}
