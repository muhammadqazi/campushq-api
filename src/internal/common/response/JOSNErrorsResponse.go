package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
)

type validationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type errorResponse struct {
	Status  bool              `json:"status"`
	Message string            `json:"message"`
	Errors  []validationError `json:"errors"`
}

func JSONErrorResponse(rw http.ResponseWriter, err error) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusBadRequest)
	var validationErrors []validationError

	switch err := err.(type) {
	case validation.Errors:
		for field, validationErr := range err {
			validationErrors = append(validationErrors, validationError{
				Field:   field,
				Message: validationErr.Error(),
			})
		}
	default:
		validationErrors = append(validationErrors, validationError{
			Field:   "general",
			Message: err.Error(),
		})
	}

	fmt.Println(validationErrors)

	response := errorResponse{
		Status:  false,
		Message: "Validation errors.",
		Errors:  validationErrors,
	}

	jsonResponse, _ := json.Marshal(response)
	rw.Write(jsonResponse)
}
