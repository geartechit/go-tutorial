package domain

import (
	"github.com/google/uuid"
	"time"
)

type Employee struct {
	ID           uuid.UUID
	Name         string
	DOB          time.Time
	Department   string
	JobTitle     string
	Address      string
	JoinedAt     time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DepartmentID string
}
