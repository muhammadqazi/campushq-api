package routers

import (
	handlers "github.com/campushq-official/campushq-api/src/internal/api/handlers/student-handlers"
	"github.com/campushq-official/campushq-api/src/internal/api/middlewares"
	"github.com/campushq-official/campushq-api/src/internal/common/logs"
	services "github.com/campushq-official/campushq-api/src/internal/core/domain/services/auth0-services"
	"github.com/gorilla/mux"
)

type studentRouter struct {
	router       *mux.Router
	logger       *logs.Logger
	auth0Service services.Auth0Service
}

func NewStudentRouter(r *mux.Router, logger *logs.Logger, auth0Service services.Auth0Service) *studentRouter {
	return &studentRouter{
		router:       r,
		logger:       logger,
		auth0Service: auth0Service,
	}
}

func (r *studentRouter) StudentRouter() {

	h := handlers.NewStudentHandler(r.logger, r.auth0Service)

	router := r.router
	router.HandleFunc("/student/signin", middlewares.Auth0Authorization("student.read", r.logger, h.StudentSignin)).Methods("POST")
	router.HandleFunc("/student/signup", middlewares.Auth0Authorization("student.create", r.logger, h.StudentSignup)).Methods("POST")
}
