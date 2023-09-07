package utils

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func GetNullableDate(value time.Time) pgtype.Date {
	if !value.IsZero() {
		return pgtype.Date{Time: value, Valid: true}
	}
	return pgtype.Date{Valid: false}
}
