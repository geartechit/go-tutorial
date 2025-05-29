package domain

import (
	"context"
	"github.com/google/uuid"
	domain "go-tutorial/internal/domain/employee"
)

type DepartmentRepository interface {
	AddEmployee(ctx context.Context, id string, employeeID uuid.UUID) (*domain.Employee, error)
}
