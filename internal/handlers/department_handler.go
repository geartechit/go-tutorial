package handlers

import (
	"encoding/json"
	"go-tutorial/internal/dto"
	"go-tutorial/internal/services"
	"go-tutorial/pkg/httputil"
	"go-tutorial/pkg/logger"
	"go-tutorial/pkg/validator"
	"net/http"
)

type DepartmentHandler struct {
	svc       services.DepartmentService
	logger    logger.Logger
	validator validator.DTOValidator
}

func NewDepartmentHandler(svc services.DepartmentService, logger logger.Logger, validator validator.DTOValidator) *DepartmentHandler {
	return &DepartmentHandler{svc: svc, logger: logger, validator: validator}
}

func (h *DepartmentHandler) AddEmployee(w http.ResponseWriter, r *http.Request) {
	var addEmpReq dto.AddEmployeeRequest

	if err := json.NewDecoder(r.Body).Decode(&addEmpReq); err != nil {
		h.logger.Error("invalid request body", logger.Field{Key: "decode error", Value: err})
		httputil.WriteError(w, http.StatusBadRequest, "invalid request body", nil)
		return
	}

	if err := h.validator.Validate(&addEmpReq); err != nil {
		h.logger.Error("validation failed", logger.Field{Key: "validation error", Value: err})
		httputil.WriteError(w, http.StatusBadRequest, "invalid request body", nil)
		return
	}

	emp, err := h.svc.AddEmployee(r.Context(), &addEmpReq)
	if err != nil {
		h.logger.Error("failed to add employee", logger.Field{Key: "employee", Value: emp})
		httputil.WriteError(w, http.StatusInternalServerError, "error adding employee", nil)
		return
	}

	httputil.WriteSuccess(w, http.StatusOK, emp)
}
