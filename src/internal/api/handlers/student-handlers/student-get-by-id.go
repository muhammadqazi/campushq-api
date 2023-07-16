package handlers

import (
	"fmt"
	"net/http"

	"github.com/campushq-official/campushq-api/src/internal/common/response"
	"github.com/gorilla/mux"
)

func (l *studentHandler) StudentGetByID(rw http.ResponseWriter, r *http.Request) {
	studentId := mux.Vars(r)["student-id"]

	if err := l.validator.StudentIDValidator(studentId); err != nil {
		response.JSONErrorResponse(rw, err)
		l.logger.PrintHTTPResponse(r, http.StatusBadRequest, "validationErrors")
		return
	}

	student, err := l.studentService.StudentFetchByID(studentId)
	if err != nil {
		response.JSONErrorResponse(rw, err)
		l.logger.PrintHTTPResponse(r, http.StatusInternalServerError, fmt.Sprintf("Error fetching student with ID %s. %s", studentId, err.Error()))
		return
	}

	response.JSONDataResponse(rw, http.StatusOK, student)
}
