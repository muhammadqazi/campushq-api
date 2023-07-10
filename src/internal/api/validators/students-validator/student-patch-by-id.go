package validators

import (
	"context"
	"fmt"
	"strconv"
	"time"

	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jackc/pgx/v5/pgtype"
)

func (v *studentValidators) StudentPatchByIdValidator(req dtos.StudentPatchDTO, id string) error {
	ctx := context.TODO()

	studentId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return fmt.Errorf("invalid student ID")
	}

	_, err = v.repository.GetStudentById(ctx, int32(studentId))
	if err != nil {
		return fmt.Errorf("student with ID %s does not exist", id)
	}

	err = validation.ValidateStruct(&req,
		validation.Field(&req.FirstName, validation.Length(1, 50)),
		validation.Field(&req.Surname, validation.Length(1, 50)),
		validation.Field(&req.Role, validation.In("student", "teacher")), // todo: add more roles
		validation.Field(&req.Nationality, validation.Length(1, 50)),
		validation.Field(&req.City, validation.Length(1, 50)),
		validation.Field(&req.Address, validation.Length(1, 255)),
		validation.Field(&req.Dob, validation.By(func(value interface{}) error {
			dob, ok := value.(pgtype.Date)
			if !ok {
				return fmt.Errorf("invalid type for DOB")
			}

			today := time.Now().UTC().Truncate(24 * time.Hour)
			dobTime := dob.Time.UTC().Truncate(24 * time.Hour)
			if dobTime.After(today) {
				return fmt.Errorf("DOB cannot be in the future")
			}

			return nil
		})),
		validation.Field(&req.PlaceOfBirth, validation.Length(1, 50)),
		validation.Field(&req.FatherName, validation.Length(1, 50)),
		validation.Field(&req.MotherName, validation.Length(1, 50)),
		validation.Field(&req.GuardianPhoneNumber, validation.Length(1, 20)),
		validation.Field(&req.PassportNumber, validation.Length(1, 20)),
		validation.Field(&req.PassportIssueDate, validation.By(func(value interface{}) error {
			issueDate, ok := value.(pgtype.Date)
			if !ok {
				return fmt.Errorf("invalid type for issueDate")
			}

			if issueDate.Time.IsZero() {
				return nil
			}

			today := time.Now().UTC().Truncate(24 * time.Hour)
			issueTime := issueDate.Time.UTC().Truncate(24 * time.Hour)
			if issueTime.After(today) {
				return fmt.Errorf("issueDate cannot be in the future")
			}

			return nil
		})),
		validation.Field(&req.PassportExpiryDate, validation.By(func(value interface{}) error {
			expiryDate, ok := value.(pgtype.Date)
			if !ok {
				return fmt.Errorf("invalid type for expiryDate")
			}

			if expiryDate.Time.IsZero() {
				return nil
			}

			today := time.Now().UTC().Truncate(24 * time.Hour)
			expiryTime := expiryDate.Time.UTC().Truncate(24 * time.Hour)
			if expiryTime.Before(today) {
				return fmt.Errorf("expiryDate cannot be in the past")
			}

			return nil
		})),
		validation.Field(&req.PassportIssuingAuthority, validation.Length(1, 100)),
		validation.Field(&req.DepartmentID, validation.By(
			func(value interface{}) error {
				departmentID, ok := value.(pgtype.Int4)
				if !ok {
					return fmt.Errorf("invalid type for department_id")
				}

				if departmentID.Int32 == 0 {
					return nil
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
		// validation.Field(&req.SupervisorID, validation.By(
		// 	func(value interface{}) error {
		// 		supervisorID, ok := value.(pgtype.Int4)
		// 		if !ok {
		// 			return fmt.Errorf("invalid type for supervisor_id")
		// 		}
		// 		model, err := v.repository.GetSupervisorById(ctx, supervisorID.Int32)
		// 		if err != nil {
		// 			return err
		// 		}

		// 		if model.SupervisorID == 0 {
		// 			return fmt.Errorf("no supervisor found with id %d", supervisorID.Int32)
		// 		}

		// 		return nil
		// 	}),
		// ),
	)

	return err
}
