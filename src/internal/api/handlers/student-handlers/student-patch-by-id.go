package handlers

import (
	"fmt"
	"net/http"

	"github.com/campushq-official/campushq-api/src/internal/common/response"
	"github.com/campushq-official/campushq-api/src/internal/common/utils"
	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
	"github.com/gorilla/mux"
)

func (l *studentHandler) StudentPatchByID(rw http.ResponseWriter, r *http.Request) {
	var req dtos.StudentPatchDTO
	if err := utils.RequestBodyParser(r, &req); err != nil {
		response.JSONErrorResponse(rw, err)
		l.logger.PrintHTTPResponse(r, http.StatusBadRequest, "Invalid request body.")
		return
	}

	studentId := mux.Vars(r)["student-id"]

	err := l.studentService.StudentModifyByID(&req, studentId)
	if err == nil {
		response.JSONMessageResponse(rw, http.StatusOK, "Student updated successfully.")
		l.logger.PrintHTTPResponse(r, http.StatusOK, fmt.Sprintf("Student with ID %s updated successfully.", studentId))
		return
	}

	response.JSONMessageResponse(rw, http.StatusInternalServerError, fmt.Sprintf("Error updating student with ID %s.", studentId))
	l.logger.PrintHTTPResponse(r, http.StatusInternalServerError, fmt.Sprintf("Error updating student with ID %s. %s", studentId, err.Error()))
}
