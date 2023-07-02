package routers

import (
	handlers "github.com/campushq-official/campushq-api/src/internal/api/handlers/student-handlers"
	"github.com/campushq-official/campushq-api/src/internal/common"
	services "github.com/campushq-official/campushq-api/src/internal/core/domain/services/auth0-services"
	"github.com/gorilla/mux"
)

func StudentRouter(r *mux.Router, logger *common.Logger, auth0Service services.Auth0Service) {

	h := handlers.NewStudentHandler(logger, auth0Service)

	r.HandleFunc("/student/signin", h.StudentSignin).Methods("POST")
	r.HandleFunc("/student/signup", h.StudentSignup).Methods("POST")
}
