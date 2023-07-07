package mappers

import (
	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
	repositories "github.com/campushq-official/campushq-api/src/internal/core/infrastructure/postgres/repositories/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func PersonalInfoCreateMapper(req dtos.StudentCreateDTO, insertedStudentId pgtype.Int4) *repositories.InsertPersonalInfoParams {
	return &repositories.InsertPersonalInfoParams{
		Nationality:              req.Nationality,
		City:                     req.City,
		Address:                  req.Address,
		Dob:                      req.DOB,
		PlaceOfBirth:             req.PlaceOfBirth,
		FatherName:               req.FatherName,
		MotherName:               req.MotherName,
		GuardianPhoneNumber:      req.GuardianPhoneNumber,
		PassportNumber:           req.PassportNumber,
		PassportIssueDate:        req.PassportIssueDate,
		PassportExpiryDate:       req.PassportExpiryDate,
		PassportIssuingAuthority: req.PassportIssueAuthority,
		StudentID:                insertedStudentId,
		IsStaff:                  false,
	}
}
