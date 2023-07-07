package validators

import (
	"context"
	"fmt"
	"time"

	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/jackc/pgx/v5/pgtype"
)

func (v *studentValidators) StudentSignupValidator(req dtos.StudentCreateDTO) error {
	ctx := context.TODO()
	err := validation.ValidateStruct(&req,
		validation.Field(&req.Email, validation.Required, is.Email, validation.Length(3, 255)),
		validation.Field(&req.FirstName, validation.Required, validation.Length(1, 50)),
		validation.Field(&req.Surname, validation.Required, validation.Length(1, 50)),
		validation.Field(&req.Sex, validation.Required, validation.In("male", "female")),
		validation.Field(&req.Role, validation.Required, validation.In("student", "teacher")),
		validation.Field(&req.Nationality, validation.Required, validation.Length(1, 50)),
		validation.Field(&req.City, validation.Required, validation.Length(1, 50)),
		validation.Field(&req.Address, validation.Required, validation.Length(1, 255)),
		validation.Field(&req.DOB, validation.Required, validation.By(func(value interface{}) error {
			dob, ok := value.(pgtype.Date)
			if !ok {
				return fmt.Errorf("invalid type for DOB")
			}

			if dob.Time.IsZero() {
				return fmt.Errorf("DOB is required")
			}

			today := time.Now().UTC().Truncate(24 * time.Hour)
			dobTime := dob.Time.UTC().Truncate(24 * time.Hour)
			if dobTime.After(today) {
				return fmt.Errorf("DOB cannot be in the future")
			}

			return nil
		})),
		validation.Field(&req.PlaceOfBirth, validation.Required, validation.Length(1, 50)),
		validation.Field(&req.FatherName, validation.Required, validation.Length(1, 50)),
		validation.Field(&req.MotherName, validation.Required, validation.Length(1, 50)),
		validation.Field(&req.GuardianPhoneNumber, validation.Required, validation.Length(1, 20)),
		validation.Field(&req.PassportNumber, validation.Required, validation.Length(1, 20)),
		validation.Field(&req.PassportIssueDate, validation.Required, validation.By(func(value interface{}) error {
			issueDate, ok := value.(pgtype.Date)
			if !ok {
				return fmt.Errorf("invalid type for issueDate")
			}

			if issueDate.Time.IsZero() {
				return fmt.Errorf("issueDate is required")
			}

			today := time.Now().UTC().Truncate(24 * time.Hour)
			issueTime := issueDate.Time.UTC().Truncate(24 * time.Hour)
			if issueTime.After(today) {
				return fmt.Errorf("issueDate cannot be in the future")
			}

			return nil
		})),
		validation.Field(&req.PassportExpiryDate, validation.Required, validation.By(func(value interface{}) error {
			expiryDate, ok := value.(pgtype.Date)
			if !ok {
				return fmt.Errorf("invalid type for expiryDate")
			}

			if expiryDate.Time.IsZero() {
				return fmt.Errorf("expiryDate is required")
			}

			today := time.Now().UTC().Truncate(24 * time.Hour)
			expiryTime := expiryDate.Time.UTC().Truncate(24 * time.Hour)
			if expiryTime.Before(today) {
				return fmt.Errorf("expiryDate cannot be in the past")
			}

			return nil
		})),
		validation.Field(&req.PassportIssueAuthority, validation.Required, validation.Length(1, 100)),
		validation.Field(&req.DepartmentID, validation.Required, validation.By(
			func(value interface{}) error {
				departmentID, ok := value.(pgtype.Int4)
				if !ok {
					return fmt.Errorf("invalid type for department_id")
				}
				model, err := v.repository.GetDepartmentById(ctx, departmentID.Int32)
				if err != nil {
					return err
				}

				if model.DepartmentID == 0 {
					return fmt.Errorf("no department found with id %d", departmentID.Int32)
				}

				return nil
			}),
		),
	)

	return err
}
