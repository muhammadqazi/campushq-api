package routers

import (
	handlers "github.com/campushq-official/campushq-api/src/internal/api/handlers/student-handlers"
	"github.com/campushq-official/campushq-api/src/internal/api/middlewares"
	validators "github.com/campushq-official/campushq-api/src/internal/api/validators/students-validator"
	"github.com/campushq-official/campushq-api/src/internal/common/logs"
	auth0 "github.com/campushq-official/campushq-api/src/internal/core/domain/services/auth0-services"
	services "github.com/campushq-official/campushq-api/src/internal/core/domain/services/student-services"
	"github.com/gorilla/mux"
)

type studentRouter struct {
	router           *mux.Router
	logger           *logs.Logger
	auth0Service     auth0.Auth0Service
	studentService   services.StudentService
	studentValidator validators.StudentValidators
}

func NewStudentRouter(r *mux.Router, logger *logs.Logger, auth0Service auth0.Auth0Service, studentService services.StudentService, studentValidator validators.StudentValidators) *studentRouter {
	return &studentRouter{
		router:           r,
		logger:           logger,
		auth0Service:     auth0Service,
		studentService:   studentService,
		studentValidator: studentValidator,
	}
}

func (r *studentRouter) StudentRouter() {

	h := handlers.NewStudentHandler(r.logger, r.auth0Service, r.studentService, r.studentValidator)

	router := r.router
	router.HandleFunc("/student/signin", h.StudentSignin).Methods("POST")
	router.HandleFunc("/student/signup", middlewares.Auth0Authorization("student.create", r.logger, h.StudentSignup)).Methods("POST")
	router.HandleFunc("/student/{student-id}", middlewares.Auth0Authorization("student.update", r.logger, h.StudentPatchByID)).Methods("PATCH")
}
