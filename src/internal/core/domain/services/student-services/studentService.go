package services

import (
	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
	repositories "github.com/campushq-official/campushq-api/src/internal/core/infrastructure/postgres/repositories/sqlc"
)

type StudentService interface {
	StudentRegister(*dtos.StudentCreateDTO) (int32, error)
	ModifyStudentById(*dtos.StudentPatchDTO, string) error
	FetchStudentById(string) ([]*dtos.StudentFetchDTO, error)
	FetchAllStudents(int64, int64) ([]*dtos.StudentFetchDTO, error)
}

type studentService struct {
	repository repositories.Store
}

func NewStudentService(repository repositories.Store) StudentService {
	return &studentService{
		repository: repository,
	}
}
