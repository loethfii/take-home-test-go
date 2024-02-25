package domain

import (
	"context"
	"service-employee/dto"
)

type EmployeeRepository interface {
	FindAllEmployee(ctx context.Context) ([]Employee, error)
	FindOneEmployee(ctx context.Context, id uint64) (Employee, error)
	LastIDEmployee(ctx context.Context) (uint64, error)
	CreateEmployee(ctx context.Context, employee Employee) error
}

type EmployeeService interface {
	FetchEmployee() ([]dto.EmployeeResponse, error)
	DetailEmployee(id uint64) (dto.EmployeeResponse, error)
	AddEmployee(employee dto.EmployeeRequest) (dto.EmployeeResponse, error)
}
