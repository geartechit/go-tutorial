package domain

import (
	"context"
	"github.com/google/uuid"
)

type EmployeeRepository interface {
	Create(ctx context.Context, e *Employee) (*Employee, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Employee, error)
	GetAll(ctx context.Context) ([]*Employee, error)
	Update(ctx context.Context, e *Employee) (*Employee, error)
	Delete(ctx context.Context, id uuid.UUID) (string, error)
}
