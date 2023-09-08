package validators

import (
	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
	repositories "github.com/campushq-official/campushq-api/src/internal/core/infrastructure/postgres/repositories/sqlc"
)

type StudentValidators interface {
	StudentSignupSchema(req dtos.StudentCreateDTO) error
	StudentSigninSchema(req dtos.StudentSigninDTO) error
	PatchStudentSchema(dtos.StudentPatchDTO, string) error
	StudentIdSchema(string) error
}

type studentValidators struct {
	repository repositories.Store
}

func NewStudentValidators(repository repositories.Store) StudentValidators {
	return &studentValidators{
		repository: repository,
	}
}
