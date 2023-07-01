package routers

import (
	handlers "github.com/campushq-official/campushq-api/src/internal/api/handlers/student-handlers"
	"github.com/campushq-official/campushq-api/src/internal/common"
	"github.com/gorilla/mux"
)

func StudentRouter(r *mux.Router, logger *common.Logger) {

	h := handlers.NewStudentHandler(logger)

	r.HandleFunc("/student/signup", h.StudentSignin).Methods("POST")
	r.HandleFunc("/student/signin", h.StudentSignup).Methods("POST")
}
