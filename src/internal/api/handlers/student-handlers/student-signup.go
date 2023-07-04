package handlers

import (
	"fmt"
	"net/http"

	"github.com/campushq-official/campushq-api/src/internal/common/response"
	"github.com/campushq-official/campushq-api/src/internal/common/utils"
	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
)

func (l *studentHandler) StudentSignup(rw http.ResponseWriter, r *http.Request) {
	var req dtos.StudentCreateDTO

	if err := utils.RequestBodyParser(r, &req); err != nil {
		fmt.Println(err)
		l.logger.PrintHTTPResponse(r, http.StatusBadRequest, "Invalid request body.")
		return
	}

	if err := l.validator.StudentSignupValidator(req); err != nil {
		response.JSONErrorResponse(rw, err)
		l.logger.PrintHTTPResponse(r, http.StatusBadRequest, "validationErrors")
		return
	}

	fmt.Println(req)
}
