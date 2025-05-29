package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"go-tutorial/internal/dto"
	"go-tutorial/internal/services"
	"go-tutorial/pkg/httputil"
	"go-tutorial/pkg/logger"
	"go-tutorial/pkg/mapper"
	"go-tutorial/pkg/validator"
	"net/http"
)

type EmployeeHandler struct {
	svc       services.EmployeeService
	logger    logger.Logger
	validator validator.DTOValidator
}

func NewEmployeeHandler(svc services.EmployeeService, logger logger.Logger, validator validator.DTOValidator) *EmployeeHandler {
	return &EmployeeHandler{svc: svc, logger: logger, validator: validator}
}

func (h *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateEmployeeRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("invalid request body", logger.Field{Key: "err", Value: err.Error()})
		httputil.WriteError(w, http.StatusBadRequest, "invalid request body", nil)
		return
	}

	validationErrors := h.validator.Validate(&req)
	if len(validationErrors) > 0 {
		h.logger.Error("validation failed", logger.Field{Key: "validation error", Value: validationErrors})
		httputil.WriteError(w, http.StatusBadRequest, "invalid request body", validationErrors)
		return
	}

	emp, err := h.svc.Create(r.Context(), &req)
	if err != nil {
		h.logger.Error("error creating employee", logger.Field{Key: "err", Value: err.Error()})
		httputil.WriteError(w, http.StatusInternalServerError, "error creating employee", nil)
		return
	}

	httputil.WriteSuccess(w, http.StatusCreated, emp)
}

func (h *EmployeeHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	employeeID, err := uuid.Parse(id)
	if err != nil {
		h.logger.Error("error parsing employee ID", logger.Field{Key: "err", Value: err.Error()})
		httputil.WriteError(w, http.StatusBadRequest, "error updating employee ID", nil)
		return
	}

	var req dto.UpdateEmployeeRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("invalid request body", logger.Field{Key: "decode error", Value: err})
		httputil.WriteError(w, http.StatusBadRequest, "invalid request body", nil)
		return
	}

	req.ID = employeeID
	validationErrors := h.validator.Validate(req)
	if len(validationErrors) > 0 {
		h.logger.Error("validation failed", logger.Field{Key: "validation error", Value: validationErrors})
		httputil.WriteError(w, http.StatusBadRequest, "invalid request body", validationErrors)
		return
	}

	emp, err := h.svc.Update(r.Context(), &req)
	if err != nil {
		h.logger.Error("error updating employee", logger.Field{Key: "err", Value: err.Error()})
		httputil.WriteError(w, http.StatusInternalServerError, "error updating employee", nil)
		return
	}

	httputil.WriteSuccess(w, http.StatusOK, emp)
}

func (h *EmployeeHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	employeeID, err := uuid.Parse(id)
	if err != nil {
		h.logger.Error("error parsing employee ID", logger.Field{Key: "err", Value: err.Error()})
		httputil.WriteError(w, http.StatusBadRequest, "error updating employee", nil)
		return
	}

	empID, err := h.svc.Delete(r.Context(), employeeID)
	if err != nil {
		h.logger.Error("error getting employee", logger.Field{Key: "err", Value: err.Error()})
		httputil.WriteError(w, http.StatusInternalServerError, "error getting employee", nil)
		return
	}

	httputil.WriteSuccess(w, http.StatusOK, empID)
}

func (h *EmployeeHandler) GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	employeeID, err := uuid.Parse(id)
	if err != nil {
		h.logger.Error("error parsing employee ID", logger.Field{Key: "err", Value: err.Error()})
		httputil.WriteError(w, http.StatusBadRequest, "error getting employee", nil)
		return
	}

	emp, err := h.svc.GetByID(r.Context(), employeeID)
	if err != nil {
		h.logger.Error("error getting employee", logger.Field{Key: "err", Value: err.Error()})
		httputil.WriteError(w, http.StatusInternalServerError, "error getting employee", nil)
		return
	}

	resp := mapper.ToEmployeeResponse(emp)
	httputil.WriteSuccess(w, http.StatusOK, resp)
	return
}

func (h *EmployeeHandler) GetEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := h.svc.GetAll(r.Context())
	if err != nil {
		h.logger.Error("failed to get employees", logger.Field{Key: "err", Value: err.Error()})
		httputil.WriteError(w, http.StatusInternalServerError, "failed to get employees", nil)
		return
	}

	resp := make([]*dto.EmployeeResponse, len(employees))
	for i, emp := range employees {
		resp[i] = mapper.ToEmployeeResponse(emp)
	}

	httputil.WriteSuccess(w, http.StatusOK, resp)
}

func (h *EmployeeHandler) GetAllEmployeesByDepartmentID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		h.logger.Error("missing employee ID")
		httputil.WriteError(w, http.StatusBadRequest, "missing employee ID", nil)
		return
	}

	employees, err := h.svc.GetAllByDepartmentID(r.Context(), id)
	if err != nil {
		h.logger.Error("error getting employees", logger.Field{Key: "err", Value: err.Error()})
		httputil.WriteError(w, http.StatusInternalServerError, "error getting employees", nil)
		return
	}

	resp := make([]*dto.EmployeeResponse, len(employees))
	for i, emp := range employees {
		resp[i] = mapper.ToEmployeeResponse(emp)
	}

	httputil.WriteSuccess(w, http.StatusOK, resp)
}
