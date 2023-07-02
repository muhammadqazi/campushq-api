package handlers

import (
	"net/http"

	"github.com/campushq-official/campushq-api/src/internal/common/logs"
)

type departmentHandler struct {
	Logger *logs.Logger
}

func NewDepartmentHandler(logger *logs.Logger) *departmentHandler {
	return &departmentHandler{Logger: logger}
}

func (h *departmentHandler) GetDepartmentByID(rw http.ResponseWriter, r *http.Request) {}
