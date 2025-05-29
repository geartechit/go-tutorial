package dto

import "github.com/google/uuid"

type AddEmployeeRequest struct {
	ID         string    `json:"id" validate:"required"`
	EmployeeID uuid.UUID `json:"employeeId" validate:"required"`
}
