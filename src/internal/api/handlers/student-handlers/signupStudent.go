package handlers

import (
	"fmt"
	"net/http"

	"github.com/campushq-official/campushq-api/src/internal/common/response"
	"github.com/campushq-official/campushq-api/src/internal/common/utils"
	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
)

func (l *studentHandler) SignupStudent(rw http.ResponseWriter, r *http.Request) {
	var req dtos.StudentCreateDTO

	if err := utils.RequestBodyParser(r, &req); err != nil {
		response.JSONErrorResponse(rw, err)
		l.logger.PrintHTTPResponse(r, http.StatusBadRequest, "Invalid request body.")
		return
	}

	if err := l.validator.StudentSignupSchema(req); err != nil {
		response.JSONErrorResponse(rw, err)
		l.logger.PrintHTTPResponse(r, http.StatusBadRequest, "validationErrors")
		return
	}

	oneTimePassword := utils.PasswordGenerator(16, 2)
	req.Password = oneTimePassword

	if lastStudentId, err := l.studentService.StudentRegister(&req); err == nil {
		studentId := lastStudentId + 1
		username := fmt.Sprintf("student-%d", studentId)
		auth0User := dtos.Auth0SignupDTO{
			Username: username,
			Email:    req.Email,
			Password: req.Password,
		}

		if err := l.auth0Service.Auth0Signup(&auth0User); err != nil {
			response.JSONErrorResponse(rw, err)
			l.logger.PrintHTTPResponse(r, http.StatusInternalServerError, "Internal server error.")
			return
		}

		data := map[string]interface{}{
			"studentID":         studentId,
			"temporaryPassword": oneTimePassword,
		}

		response.JSONDataResponse(rw, http.StatusOK, data)
		l.logger.PrintHTTPResponse(r, http.StatusOK, "Student created successfully.")
		return

	}

	response.JSONMessageResponse(rw, http.StatusInternalServerError, "Internal server error.")
	l.logger.PrintHTTPResponse(r, http.StatusInternalServerError, "Internal server error.")

}
