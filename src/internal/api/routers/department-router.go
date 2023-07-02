package routers

import (
	handlers "github.com/campushq-official/campushq-api/src/internal/api/handlers/department-handlers"
	"github.com/campushq-official/campushq-api/src/internal/common/logs"
	"github.com/gorilla/mux"
)

type departmentRouter struct {
	router *mux.Router
	logger *logs.Logger
}

func NewDepartmentRouter(r *mux.Router, logger *logs.Logger) *departmentRouter {
	return &departmentRouter{
		router: r,
		logger: logger,
	}
}

func (r *departmentRouter) DepartmentRouter() {

	h := handlers.NewDepartmentHandler(r.logger)

	router := r.router
	router.HandleFunc("/department/{id}", h.GetDepartmentByID).Methods("GET")
}
