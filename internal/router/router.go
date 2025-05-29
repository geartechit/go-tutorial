package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go-tutorial/internal/handlers"
)

func New(employeeHandler *handlers.EmployeeHandler, departmentHandler *handlers.DepartmentHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/v1", func(rg chi.Router) {
		RegisterEmployeeRoutes(rg, employeeHandler)
		RegisterDepartmentRoutes(rg, departmentHandler)
	})

	return r
}
