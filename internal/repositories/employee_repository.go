package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"go-tutorial/internal/database/sqlc/queries"
	"go-tutorial/internal/domain/employee"
)

type employeeRepository struct {
	q *queries.Queries
}

func NewEmployeeRepository(q *queries.Queries) domain.EmployeeRepository {
	return &employeeRepository{
		q: q,
	}
}

func (r *employeeRepository) Create(ctx context.Context, e *domain.Employee) (*domain.Employee, error) {
	params := queries.CreateEmployeeParams{
		Name:       e.Name,
		Dob:        pgtype.Date{Time: e.DOB, Valid: true},
		Department: e.Department,
		JobTitle:   e.JobTitle,
		Address:    e.Address,
		JoinedAt:   e.JoinedAt,
	}
	row, err := r.q.CreateEmployee(ctx, params)
	if err != nil {
		return nil, err
	}

	return &domain.Employee{
		ID:         row.ID,
		Name:       row.Name,
		DOB:        row.Dob.Time,
		Department: row.Department,
		JobTitle:   row.JobTitle,
		Address:    row.Address,
		JoinedAt:   row.JoinedAt,
		CreatedAt:  row.CreatedAt,
		UpdatedAt:  row.UpdatedAt,
	}, nil
}

func (r *employeeRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Employee, error) {
	row, err := r.q.GetEmployeeById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.Employee{
		ID:         row.ID,
		Name:       row.Name,
		DOB:        row.Dob.Time,
		Department: row.Department,
		JobTitle:   row.JobTitle,
		Address:    row.Address,
		JoinedAt:   row.JoinedAt,
		CreatedAt:  row.CreatedAt,
		UpdatedAt:  row.UpdatedAt,
	}, nil
}

func (r *employeeRepository) GetAll(ctx context.Context) ([]*domain.Employee, error) {
	rows, err := r.q.GetAllEmployee(ctx)
	if err != nil {
		return nil, err
	}

	employees := make([]*domain.Employee, len(rows))
	for i, row := range rows {
		employees[i] = &domain.Employee{
			ID:         row.ID,
			Name:       row.Name,
			DOB:        row.Dob.Time,
			Department: row.Department,
			JobTitle:   row.JobTitle,
			Address:    row.Address,
			JoinedAt:   row.JoinedAt,
			CreatedAt:  row.CreatedAt,
			UpdatedAt:  row.UpdatedAt,
		}
	}

	return employees, nil
}

func (r *employeeRepository) Update(ctx context.Context, e *domain.Employee) (*domain.Employee, error) {
	existingEmp, err := r.GetByID(ctx, e.ID)
	if err != nil {
		return nil, err
	}

	if existingEmp == nil {
		return nil, fmt.Errorf("employee not found")
	}

	if e.Name != "" {
		existingEmp.Name = e.Name
	}
	if !e.DOB.IsZero() {
		existingEmp.DOB = e.DOB
	}
	if e.Department != "" {
		existingEmp.Department = e.Department
	}
	if e.JobTitle != "" {
		existingEmp.JobTitle = e.JobTitle
	}
	if e.Address != "" {
		existingEmp.Address = e.Address
	}
	if !e.JoinedAt.IsZero() {
		existingEmp.JoinedAt = e.JoinedAt
	}

	params := queries.UpdateEmployeeParams{
		Name:       existingEmp.Name,
		Dob:        pgtype.Date{Time: existingEmp.DOB, Valid: true},
		Department: existingEmp.Department,
		JobTitle:   existingEmp.JobTitle,
		Address:    existingEmp.Address,
		JoinedAt:   existingEmp.JoinedAt,
		ID:         existingEmp.ID,
	}
	row, err := r.q.UpdateEmployee(ctx, params)
	if err != nil {
		return nil, err
	}

	return &domain.Employee{
		ID:         row.ID,
		Name:       row.Name,
		DOB:        row.Dob.Time,
		Department: row.Department,
		JobTitle:   row.JobTitle,
		Address:    row.Address,
		JoinedAt:   row.JoinedAt,
		CreatedAt:  row.CreatedAt,
		UpdatedAt:  row.UpdatedAt,
	}, nil
}

func (r *employeeRepository) Delete(ctx context.Context, id uuid.UUID) (string, error) {
	deletedID, err := r.q.DeleteEmployee(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", fmt.Errorf("employee with id %s does not exist", id)
	}
	if err != nil {
		return "", err
	}
	return deletedID.String(), nil
}
