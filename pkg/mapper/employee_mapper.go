package mapper

import (
	"go-tutorial/internal/domain/employee"
	"go-tutorial/internal/dto"
	"time"
)

func ToEmployeeResponse(emp *domain.Employee) *dto.EmployeeResponse {
	return &dto.EmployeeResponse{
		ID:         emp.ID.String(),
		Name:       emp.Name,
		DOB:        emp.DOB.Format(time.RFC3339),
		Department: emp.Department,
		JobTitle:   emp.JobTitle,
		Address:    emp.Address,
		JoinedAt:   emp.JoinedAt.Format(time.RFC3339),
		CreatedAt:  emp.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  emp.UpdatedAt.Format(time.RFC3339),
	}
}

func ToCreateEmployeeModel(req *dto.CreateEmployeeRequest) *domain.Employee {
	return &domain.Employee{
		Name:       req.Name,
		DOB:        req.Dob,
		Department: req.Department,
		JobTitle:   req.JobTitle,
		Address:    req.Address,
		JoinedAt:   req.JoinedAt,
	}
}

func ToUpdateEmployeeModel(req *dto.UpdateEmployeeRequest) *domain.Employee {
	return &domain.Employee{
		ID:         req.ID,
		Name:       req.Name,
		DOB:        req.Dob,
		Department: req.Department,
		JobTitle:   req.JobTitle,
		Address:    req.Address,
		JoinedAt:   req.JoinedAt,
	}
}
