package routers

import (
	handlers "github.com/campushq-official/campushq-api/src/internal/api/handlers/department-handlers"
	"github.com/campushq-official/campushq-api/src/internal/common"
	"github.com/gorilla/mux"
)

func DepartmentRouter(r *mux.Router, logger *common.Logger) {

	h := handlers.NewDepartmentHandler(logger)

	r.HandleFunc("/department/{id}", h.GetDepartmentByID).Methods("GET")
}
