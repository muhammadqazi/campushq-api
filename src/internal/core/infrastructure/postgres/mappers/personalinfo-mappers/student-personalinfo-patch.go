package mappers

import (
	"strconv"
	"time"

	"github.com/campushq-official/campushq-api/src/internal/common/utils"
	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
	repositories "github.com/campushq-official/campushq-api/src/internal/core/infrastructure/postgres/repositories/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func StudentPersonalInfoPatchMapper(req dtos.StudentPatchDTO, id string) *repositories.UpdateStudentPersonalInfoParams {
	studentId, _ := strconv.ParseInt(id, 10, 32)

	return &repositories.UpdateStudentPersonalInfoParams{
		StudentID:                pgtype.Int4{Int32: int32(studentId), Valid: true},
		Nationality:              utils.GetNullableText(req.Nationality),
		City:                     utils.GetNullableText(req.City),
		Address:                  utils.GetNullableText(req.Address),
		PlaceOfBirth:             utils.GetNullableText(req.PlaceOfBirth),
		LocalAddress:             utils.GetNullableText(req.LocalAddress),
		Dob:                      utils.GetNullableDate(time.Time(req.Dob.Time)),
		FatherName:               utils.GetNullableText(req.FatherName),
		MotherName:               utils.GetNullableText(req.MotherName),
		GuardianPhoneNumber:      utils.GetNullableText(req.GuardianPhoneNumber),
		LocalPhoneNumber:         utils.GetNullableText(req.LocalPhoneNumber),
		PassportNumber:           utils.GetNullableText(req.PassportNumber),
		PassportIssuingAuthority: utils.GetNullableText(req.PassportIssuingAuthority),
		IDCardNumber:             utils.GetNullableText(req.IDCardNumber),
		PassportIssueDate:        utils.GetNullableDate(time.Time(req.PassportIssueDate.Time)),
		PassportExpiryDate:       utils.GetNullableDate(time.Time(req.PassportExpiryDate.Time)),
	}
}
