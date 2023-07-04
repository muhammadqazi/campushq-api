package validators

import (
	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func (v *studentValidators) StudentSignupValidator(req dtos.StudentCreateDTO) error {
	err := validation.ValidateStruct(&req,
		validation.Field(&req.Email, validation.Required, is.Email, validation.Length(3, 255)),
		validation.Field(&req.Username, validation.Required, validation.Length(3, 50)),
		validation.Field(&req.FirstName, validation.Required, validation.Length(1, 50)),
		validation.Field(&req.LastName, validation.Required, validation.Length(1, 50)),
		validation.Field(&req.Sex, validation.Required, validation.In("male", "female")),
		validation.Field(&req.Role, validation.Required, validation.In("student", "teacher")),
		validation.Field(&req.Nationality, validation.Required, validation.Length(1, 50)),
		validation.Field(&req.City, validation.Required, validation.Length(1, 50)),
		validation.Field(&req.Address, validation.Required, validation.Length(1, 255)),
		validation.Field(&req.DOB, validation.Required, validation.Length(10, 10)),
		validation.Field(&req.PlaceOfBirth, validation.Required, validation.Length(1, 50)),
		validation.Field(&req.FaherName, validation.Required, validation.Length(1, 50)),
		validation.Field(&req.MotherName, validation.Required, validation.Length(1, 50)),
		validation.Field(&req.PassportNumber, validation.Required, validation.Length(1, 20)),
		validation.Field(&req.PassportIssueDate, validation.Required, validation.Length(10, 10)),
		validation.Field(&req.PassportExpiryDate, validation.Required, validation.Length(10, 10)),
		validation.Field(&req.PassportIssueAuthority, validation.Required, validation.Length(1, 100)),
		validation.Field(&req.DepartmentID, validation.Required),
	)

	return err
}
