package services

import (
	"context"
	"fmt"
	"go-tutorial/internal/domain/department"
	emp "go-tutorial/internal/domain/employee"
	"go-tutorial/internal/dto"
	"go-tutorial/pkg/logger"
)

type DepartmentService interface {
	AddEmployee(ctx context.Context, req *dto.AddEmployeeRequest) (*emp.Employee, error)
}

type departmentService struct {
	repo   domain.DepartmentRepository
	logger logger.Logger
}

func NewDepartmentService(repo domain.DepartmentRepository, logger logger.Logger) DepartmentService {
	return &departmentService{repo: repo, logger: logger}
}

func (s *departmentService) AddEmployee(ctx context.Context, req *dto.AddEmployeeRequest) (*emp.Employee, error) {
	employee, err := s.repo.AddEmployee(ctx, req.ID, req.EmployeeID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error adding employee %s to department: %s", req.EmployeeID, req.ID))
		return nil, err
	}

	return employee, nil
}
