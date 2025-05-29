package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"go-tutorial/internal/database/sqlc/queries"
	"go-tutorial/internal/domain/department"
	emp "go-tutorial/internal/domain/employee"
)

type departmentRepository struct {
	q  *queries.Queries
	db *pgxpool.Pool
}

func NewDepartmentRepository(q *queries.Queries, db *pgxpool.Pool) domain.DepartmentRepository {
	return &departmentRepository{
		q:  q,
		db: db,
	}
}

func (r *departmentRepository) AddEmployee(ctx context.Context, id string, employeeID uuid.UUID) (*emp.Employee, error) {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}

	qtx := queries.New(tx)
	existingEmp, err := qtx.GetEmployeeById(ctx, employeeID)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("employee not found")
	}

	if existingEmp.DepartmentID.Valid {
		if err = tx.Rollback(ctx); err != nil {
			return nil, err
		}
		return nil, errors.New(fmt.Sprintf("employee with id %s already in department %s", employeeID.String(), id))
	}

	params := queries.UpdateEmployeeParams{
		Name:         existingEmp.Name,
		Dob:          existingEmp.Dob,
		Department:   existingEmp.Department,
		JobTitle:     existingEmp.JobTitle,
		Address:      existingEmp.Address,
		JoinedAt:     existingEmp.JoinedAt,
		ID:           existingEmp.ID,
		DepartmentID: pgtype.Text{String: id, Valid: true},
	}
	row, err := qtx.UpdateEmployee(ctx, params)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	updatedEmp := emp.Employee{
		ID:           row.ID,
		Name:         row.Name,
		DOB:          row.Dob.Time,
		Department:   row.Department,
		JobTitle:     row.JobTitle,
		Address:      row.Address,
		JoinedAt:     row.JoinedAt,
		DepartmentID: row.DepartmentID.String,
		CreatedAt:    row.CreatedAt,
		UpdatedAt:    row.UpdatedAt,
	}

	return &updatedEmp, nil
}
