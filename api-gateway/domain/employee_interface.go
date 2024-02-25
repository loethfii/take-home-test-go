package domain

import "api-gateway/dto"

type EmployeeService interface {
	FetchEmployee(accessToken string) (dto.EmployeeResponse, error)
	DetailEmployee(accessToken string, id uint64) (dto.EmployeeResponse, error)
	AddEmployee(accessToken string, employee dto.EmployeeRequest) (dto.EmployeeResponse, error)
}
