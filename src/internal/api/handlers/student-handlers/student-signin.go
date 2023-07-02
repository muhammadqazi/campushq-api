package handlers

import (
	"net/http"

	"github.com/campushq-official/campushq-api/src/internal/common/logs"
	services "github.com/campushq-official/campushq-api/src/internal/core/domain/services/auth0-services"
)

type studentHandler struct {
	logger       *logs.Logger
	auth0Service services.Auth0Service
}

func NewStudentHandler(logger *logs.Logger, auth0Service services.Auth0Service) *studentHandler {
	return &studentHandler{
		logger:       logger,
		auth0Service: auth0Service,
	}
}

func (l *studentHandler) StudentSignin(rw http.ResponseWriter, r *http.Request) {

	// fmt.Println("sgnininiiininiinininiinininin--------")
	// req := dtos.StudentSigninDTO{
	// 	Username: "student-21906778",
	// 	Password: "1234!@@#$qwerQWER",
	// }

	// res, err := l.auth0Service.Auth0UserSignin(req)
	// fmt.Println(res, err)
}
