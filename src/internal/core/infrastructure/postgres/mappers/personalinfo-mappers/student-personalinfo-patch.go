package mappers

import (
	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
	repositories "github.com/campushq-official/campushq-api/src/internal/core/infrastructure/postgres/repositories/sqlc"
)

func StudentPersonalInfoPatchMapper(req dtos.StudentPatchDTO) *repositories.UpdatePersonalInfoParams {
	return &repositories.UpdatePersonalInfoParams{}
}
