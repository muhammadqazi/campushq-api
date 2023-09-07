package utils

import "github.com/jackc/pgx/v5/pgtype"

func GetNullableText(value string) pgtype.Text {
	if value != "" {
		return pgtype.Text{String: value, Valid: true}
	}
	return pgtype.Text{Valid: false}
}
