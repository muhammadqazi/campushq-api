package handlers

import (
	"net/http"

	"github.com/campushq-official/campushq-api/src/internal/common/logs"
)

type DepartmentHandler interface {
	GetDepartmentByID(rw http.ResponseWriter, r *http.Request)
}

type departmentHandler struct {
	Logger *logs.Logger
}

func NewDepartmentHandler(logger *logs.Logger) DepartmentHandler {
	return &departmentHandler{Logger: logger}
}
