package handlers

import (
	"net/http"

	validators "github.com/campushq-official/campushq-api/src/internal/api/validators/students-validator"
	"github.com/campushq-official/campushq-api/src/internal/common/logs"
	services "github.com/campushq-official/campushq-api/src/internal/core/domain/services/auth0-services"
)

type StudentHandler interface {
	StudentSignin(rw http.ResponseWriter, r *http.Request)
	StudentSignup(rw http.ResponseWriter, r *http.Request)
}

type studentHandler struct {
	logger       *logs.Logger
	auth0Service services.Auth0Service
	validator    validators.StudentValidators
}

func NewStudentHandler(logger *logs.Logger, auth0Service services.Auth0Service) StudentHandler {
	v := validators.NewStudentValidators()
	return &studentHandler{
		logger:       logger,
		auth0Service: auth0Service,
		validator:    v,
	}
}
