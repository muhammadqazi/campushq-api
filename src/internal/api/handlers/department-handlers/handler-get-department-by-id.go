package handlers

import (
	"net/http"

	"github.com/campushq-official/campushq-api/src/internal/common"
)

type departmentHandler struct {
	Logger *common.Logger
}

func NewDepartmentHandler(logger *common.Logger) *departmentHandler {
	return &departmentHandler{Logger: logger}
}

func (h *departmentHandler) GetDepartmentByID(rw http.ResponseWriter, r *http.Request) {}
