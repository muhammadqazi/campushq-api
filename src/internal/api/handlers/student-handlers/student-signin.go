package handlers

import (
	"fmt"
	"net/http"

	"github.com/campushq-official/campushq-api/src/internal/common"
	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
	services "github.com/campushq-official/campushq-api/src/internal/core/domain/services/auth0-services"
)

type studentHandler struct {
	logger       *common.Logger
	auth0Service services.Auth0Service
}

func NewStudentHandler(logger *common.Logger, auth0Service services.Auth0Service) *studentHandler {
	return &studentHandler{
		logger:       logger,
		auth0Service: auth0Service,
	}
}

func (l *studentHandler) StudentSignin(rw http.ResponseWriter, r *http.Request) {

	req := dtos.StudentSigninDTO{
		Username: "student-21906778",
		Password: "1234!@@#$qwerQWER",
	}

	res, err := l.auth0Service.Auth0UserSignin(req)
	fmt.Println(res, err)
}
