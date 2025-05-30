package services

import (
	"context"
	"github.com/google/uuid"
	"go-tutorial/internal/domain/employee"
	"go-tutorial/internal/dto"
	"go-tutorial/pkg/logger"
)

type EmployeeService interface {
	Create(ctx context.Context, req *dto.CreateEmployeeRequest) (*domain.Employee, error)
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Employee, error)
	GetAll(ctx context.Context) ([]*domain.Employee, error)
	Update(ctx context.Context, req *dto.UpdateEmployeeRequest) (*domain.Employee, error)
	Delete(ctx context.Context, id uuid.UUID) (string, error)
	GetAllByDepartmentID(ctx context.Context, departmentID string) ([]*domain.Employee, error)
}

type employeeService struct {
	repo   domain.EmployeeRepository
	logger logger.Logger
}

func NewEmployeeService(repo domain.EmployeeRepository, logger logger.Logger) EmployeeService {
	return &employeeService{
		repo:   repo,
		logger: logger,
	}
}

func (s *employeeService) Create(ctx context.Context, req *dto.CreateEmployeeRequest) (*domain.Employee, error) {
	e := &domain.Employee{
		Name:       req.Name,
		DOB:        req.Dob,
		Department: req.Department,
		JobTitle:   req.JobTitle,
		Address:    req.Address,
		JoinedAt:   req.JoinedAt,
	}
	emp, err := s.repo.Create(ctx, e)
	if err != nil {
		s.logger.Error("failed to create emp", logger.Field{Key: "err", Value: err.Error()})
	}

	return emp, err
}

func (s *employeeService) GetByID(ctx context.Context, id uuid.UUID) (*domain.Employee, error) {
	emp, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("failed to get employee", logger.Field{Key: "err", Value: err.Error()})
		return nil, err
	}

	return emp, nil
}

func (s *employeeService) GetAll(ctx context.Context) ([]*domain.Employee, error) {
	employees, err := s.repo.GetAll(ctx)
	if err != nil {
		s.logger.Error("failed to get employees", logger.Field{Key: "err", Value: err.Error()})
		return nil, err
	}

	return employees, nil
}

func (s *employeeService) Update(ctx context.Context, req *dto.UpdateEmployeeRequest) (*domain.Employee, error) {
	e := &domain.Employee{
		ID:         req.ID,
		Name:       req.Name,
		DOB:        req.Dob,
		Department: req.Department,
		JobTitle:   req.JobTitle,
		Address:    req.Address,
		JoinedAt:   req.JoinedAt,
	}

	emp, err := s.repo.Update(ctx, e)
	if err != nil {
		s.logger.Error("failed to update employee", logger.Field{Key: "err", Value: err.Error()})
		return nil, err
	}

	return emp, nil
}

func (s *employeeService) Delete(ctx context.Context, id uuid.UUID) (string, error) {
	employeeID, err := s.repo.Delete(ctx, id)
	if err != nil {
		// can be refactored
		s.logger.Error("failed to delete employee", logger.Field{Key: "err", Value: err.Error()})
		return "", err
	}

	return employeeID, nil
}

func (s *employeeService) GetAllByDepartmentID(ctx context.Context, departmentID string) ([]*domain.Employee, error) {
	employees, err := s.repo.GetAllByDepartmentID(ctx, departmentID)
	if err != nil {
		s.logger.Error("failed to get employees", logger.Field{Key: "err", Value: err.Error()})
		return nil, err
	}

	return employees, nil
}
