package router

import (
	"github.com/go-chi/chi/v5"
	"go-tutorial/internal/handlers"
)

func RegisterDepartmentRoutes(r chi.Router, h *handlers.DepartmentHandler) {
	r.Route("/departments", func(rg chi.Router) {
		rg.Put("/employee", h.AddEmployee)
	})
}
