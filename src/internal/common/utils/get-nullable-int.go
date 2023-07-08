package utils

import "github.com/jackc/pgx/v5/pgtype"

func GetNullableInt(value int) pgtype.Int4 {
	if value != 0 {
		return pgtype.Int4{Int32: int32(value), Valid: true}
	}
	return pgtype.Int4{Valid: false}
}
