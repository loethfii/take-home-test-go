package service

import (
	"context"
	"service-employee/domain"
	"service-employee/dto"
)

type employeeService struct {
	employeeRepository domain.EmployeeRepository
}

func NewEmployeeService(employeeRepository domain.EmployeeRepository) domain.EmployeeService {
	return &employeeService{employeeRepository}
}

func (s *employeeService) FetchEmployee() ([]dto.EmployeeResponse, error) {
	ctx := context.Background()
	employees, err := s.employeeRepository.FindAllEmployee(ctx)
	if err != nil {
		return nil, err
	}
	
	var employeeResponses []dto.EmployeeResponse
	
	for _, employee := range employees {
		employeeResponses = append(employeeResponses, dto.EmployeeResponse{
			ID:   employee.ID,
			Name: employee.Name,
		})
	}
	
	return employeeResponses, nil
}

func (s *employeeService) DetailEmployee(id uint64) (dto.EmployeeResponse, error) {
	ctx := context.Background()
	
	employee, err := s.employeeRepository.FindOneEmployee(ctx, id)
	if err != nil {
		return dto.EmployeeResponse{}, err
	}
	
	return dto.EmployeeResponse{
		ID:   employee.ID,
		Name: employee.Name,
	}, nil
}

func (s *employeeService) AddEmployee(employee dto.EmployeeRequest) (dto.EmployeeResponse, error) {
	ctx := context.Background()
	
	if err := s.employeeRepository.CreateEmployee(ctx, domain.Employee{
		Name: employee.Name,
	}); err != nil {
		return dto.EmployeeResponse{}, err
	}
	
	lastID, err := s.employeeRepository.LastIDEmployee(ctx)
	if err != nil {
		return dto.EmployeeResponse{}, err
	}
	
	return dto.EmployeeResponse{
		ID:   lastID,
		Name: employee.Name,
	}, nil
}
