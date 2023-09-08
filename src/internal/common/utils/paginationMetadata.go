package utils

import (
	"fmt"

	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
)

func PaginationMetadata(limit int64, offset int64, totalCount int, page int64) *dtos.PaginationMetadata {

	metadata := dtos.PaginationMetadata{
		Limit:      int32(limit),
		From:       int32(offset) + 1,
		To:         int32(offset) + int32(totalCount),
		TotalCount: int32(totalCount),
		Links: []dtos.Link{
			{
				Self:  fmt.Sprintf("api/v1/students?page=%d&limit=%d", page, limit),
				First: fmt.Sprintf("api/v1/students?page=%d&limit=%d", 1, limit),
			},
		},
	}

	previousPage := page - 1
	if page > 1 {
		metadata.Links[0].Previous = fmt.Sprintf("api/v1/students?page=%d&limit=%d", previousPage, limit)
	} else {
		metadata.Links[0].Previous = nil
	}

	nextPage := page + 1
	if totalCount > 0 {
		metadata.Links[0].Next = fmt.Sprintf("api/v1/students?page=%d&limit=%d", nextPage, limit)
	} else {
		metadata.Links[0].Next = nil
	}
	return &metadata
}
