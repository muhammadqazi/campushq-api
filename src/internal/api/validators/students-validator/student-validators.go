package validators

import dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"

type StudentValidators interface {
	StudentSigninValidator(req dtos.StudentSigninDTO) error
	StudentSignupValidator(req dtos.StudentCreateDTO) error
}

type studentValidators struct{}

func NewStudentValidators() StudentValidators {
	return &studentValidators{}
}
