package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/campushq-official/campushq-api/src/internal/common/response"
	"github.com/campushq-official/campushq-api/src/internal/common/utils"
	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
)

func (l *studentHandler) SigninStudent(rw http.ResponseWriter, r *http.Request) {
	var req dtos.StudentSigninDTO
	if err := utils.RequestBodyParser(r, &req); err != nil {
		response.JSONErrorResponse(rw, err)
		l.logger.PrintHTTPResponse(r, http.StatusBadRequest, "Invalid request body.")
		return
	}

	if err := l.validator.StudentSigninSchema(req); err != nil {
		response.JSONErrorResponse(rw, err)
		l.logger.PrintHTTPResponse(r, http.StatusBadRequest, "validationErrors")
		return
	}

	if token, err := l.auth0Service.Auth0UserSignin(req); err == nil {

		json.NewEncoder(rw).Encode(
			map[string]interface{}{
				"status": true,
				"token":  token,
			},
		)
		l.logger.PrintHTTPResponse(r, http.StatusOK, "Student signed in successfully.")
		return
	}

	response.JSONMessageResponse(rw, http.StatusInternalServerError, "Internal server error.")
	l.logger.PrintHTTPResponse(r, http.StatusInternalServerError, "Internal server error.")
}
