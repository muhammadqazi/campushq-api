package handlers

import (
	"net/http"
	"strconv"

	"github.com/campushq-official/campushq-api/src/internal/common/response"
	"github.com/campushq-official/campushq-api/src/internal/common/utils"
	"github.com/gorilla/mux"
)

func (l *studentHandler) GetAllStudents(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	page, _ := strconv.ParseInt(vars["page"], 10, 32)
	limit, _ := strconv.ParseInt(vars["limit"], 10, 32)

	offset := (page - 1) * limit

	students, err := l.studentService.FetchAllStudents(limit, offset)
	if err != nil {
		response.JSONErrorResponse(rw, err)
		l.logger.PrintHTTPResponse(r, http.StatusInternalServerError, err.Error())
		return
	}

	totalCount := len(students)

	paginatedResponse := utils.PaginationMetadata(limit, offset, totalCount, page)

	response.JSONDataResponse(rw, http.StatusOK, students, "_metadata", paginatedResponse)
}
