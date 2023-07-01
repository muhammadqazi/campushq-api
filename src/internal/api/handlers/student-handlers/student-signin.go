package handlers

import (
	"net/http"

	"github.com/campushq-official/campushq-api/src/internal/common"
)

type studentHandler struct {
	logger *common.Logger
}

func NewStudentHandler(logger *common.Logger) *studentHandler {
	return &studentHandler{logger: logger}
}

func (l *studentHandler) StudentSignin(rw http.ResponseWriter, r *http.Request) {
	l.logger.Info("Student Signin", r.URL.Path)
}
