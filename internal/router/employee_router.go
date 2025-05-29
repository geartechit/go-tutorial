package router

import (
	"github.com/go-chi/chi/v5"
	"go-tutorial/internal/handlers"
)

func RegisterEmployeeRoutes(r chi.Router, h *handlers.EmployeeHandler) {
	r.Route("/employees", func(rg chi.Router) {
		rg.Get("/department/{id}", h.GetAllEmployeesByDepartmentID)
		rg.Get("/", h.GetEmployees)
		rg.Get("/{id}", h.GetEmployeeByID)
		rg.Post("/", h.CreateEmployee)
		rg.Put("/{id}", h.UpdateEmployee)
		rg.Delete("/{id}", h.DeleteEmployee)
	})
}
