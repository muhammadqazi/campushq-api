package handlers

import (
	"net/http"

	validators "github.com/campushq-official/campushq-api/src/internal/api/validators/students-validator"
	"github.com/campushq-official/campushq-api/src/internal/common/logs"
	auth0Service "github.com/campushq-official/campushq-api/src/internal/core/domain/services/auth0-services"
	services "github.com/campushq-official/campushq-api/src/internal/core/domain/services/student-services"
)

type StudentHandler interface {
	StudentSignin(rw http.ResponseWriter, r *http.Request)
	StudentSignup(rw http.ResponseWriter, r *http.Request)
	StudentPatchByID(rw http.ResponseWriter, r *http.Request)
}

type studentHandler struct {
	logger         *logs.Logger
	auth0Service   auth0Service.Auth0Service
	studentService services.StudentService
	validator      validators.StudentValidators
}

func NewStudentHandler(logger *logs.Logger, auth0Service auth0Service.Auth0Service, studentService services.StudentService, validator validators.StudentValidators) StudentHandler {
	return &studentHandler{
		logger:         logger,
		auth0Service:   auth0Service,
		validator:      validator,
		studentService: studentService,
	}
}
