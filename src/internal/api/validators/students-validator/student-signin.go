package validators

import (
	"context"
	"fmt"
	"strconv"

	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
	validation "github.com/go-ozzo/ozzo-validation"
)

func (v *studentValidators) StudentSigninValidator(req dtos.StudentSigninDTO) error {
	ctx := context.TODO()
	err := validation.ValidateStruct(&req,
		validation.Field(&req.Username, validation.Required, validation.By(
			func(value interface{}) error {
				studentId, err := strconv.ParseInt(req.Username, 10, 32)
				if err != nil {
					return fmt.Errorf("invalid username")
				}

				model, err := v.repository.GetStudentById(ctx, int32(studentId))
				if err != nil {
					return fmt.Errorf("no student found with id %d", studentId)
				}

				if model.StudentID == 0 {
					return fmt.Errorf("no student found with id %d", studentId)
				}

				return nil
			}),
		),
	)

	return err
}
