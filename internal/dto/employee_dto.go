package dto

import (
	"github.com/google/uuid"
	"time"
)

type EmployeeResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	DOB        string `json:"dob"`
	Department string `json:"department"`
	JobTitle   string `json:"jobTitle"`
	Address    string `json:"address"`
	JoinedAt   string `json:"joinedAt"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

type CreateEmployeeRequest struct {
	Name       string    `json:"name" validate:"required,min=2,max=100"`
	Dob        time.Time `json:"dob"  validate:"required"`
	Department string    `json:"department"  validate:"required"`
	JobTitle   string    `json:"jobTitle"  validate:"required"`
	Address    string    `json:"address"   validate:"required"`
	JoinedAt   time.Time `json:"joinedAt"   validate:"required"`
}

type UpdateEmployeeRequest struct {
	ID         uuid.UUID `json:"id" validate:"required,uuid"`
	Name       string    `json:"name,omitempty"`
	Dob        time.Time `json:"dob,omitempty"`
	Department string    `json:"department,omitempty"`
	JobTitle   string    `json:"jobTitle,omitempty"`
	Address    string    `json:"address,omitempty"`
	JoinedAt   time.Time `json:"joinedAt,omitempty"`
}
