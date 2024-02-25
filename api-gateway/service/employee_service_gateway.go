package service

import (
	"api-gateway/domain"
	"api-gateway/dto"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type employeeService struct {
}

func NewEmployeeService() domain.EmployeeService {
	return &employeeService{}
}

var employeeURI = "http://employee:3002"

//var employeeURI = "http://localhost:3002"

// FindAllEmployee 		godoc
// @Summary			Get All employee.
// @Description		Return list of employee.
// @Tags			Employee
// @Success			200 {obejct} dto.EmployeeResponse{}
// @Router			/employees [get]
// @Security ApiKeyAuth
func (s *employeeService) FetchEmployee(accessToken string) (dto.EmployeeResponse, error) {
	
	req, err := http.NewRequest("GET", employeeURI+"/api/v1/employees", nil)
	if err != nil {
		return dto.EmployeeResponse{}, err
	}
	
	req.Header.Set("Content-Type", "application/json")
	
	req.Header.Set("access_token", accessToken)
	
	client := &http.Client{}
	
	resp, err := client.Do(req)
	
	if err != nil {
		return dto.EmployeeResponse{}, err
	}
	
	defer resp.Body.Close()
	
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return dto.EmployeeResponse{}, err
	}
	
	var employees dto.EmployeeResponse
	
	if err = json.Unmarshal(responseBody, &employees); err != nil {
		return dto.EmployeeResponse{}, err
	}
	
	return employees, nil
}

// FindOneEmployee 		godoc
// @Summary			Find detail employee.
// @Description		Return one of employee.
// @Tags			Employee
// @Param			id path int true "Employee ID"
// @Success			200 {object} dto.EmployeeResponse{}
// @Router			/employees/{id} [get]
// @Security 		ApiKeyAuth
func (s *employeeService) DetailEmployee(accessToken string, id uint64) (dto.EmployeeResponse, error) {
	url := fmt.Sprintf("%s/api/v1/employees/%d", employeeURI, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return dto.EmployeeResponse{}, err
	}
	
	req.Header.Set("Content-Type", "application/json")
	
	req.Header.Set("access_token", accessToken)
	
	client := &http.Client{}
	
	resp, err := client.Do(req)
	
	if err != nil {
		return dto.EmployeeResponse{}, err
	}
	
	defer resp.Body.Close()
	
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return dto.EmployeeResponse{}, err
	}
	
	var employees dto.EmployeeResponse
	
	if err = json.Unmarshal(responseBody, &employees); err != nil {
		return dto.EmployeeResponse{}, err
	}
	
	return employees, nil
}

// CreateTags		godoc
// @Summary			Add employee
// @Description		Save employee into Db.
// @Produce			application/json
// @Param			tags body dto.EmployeeRequest true "request body to add employee"
// @Tags			Employee
// @Success			200 {object} dto.EmployeeResponse{}
// @Router			/employees [post]
// @Security 		ApiKeyAuth
func (s *employeeService) AddEmployee(accessToken string, employee dto.EmployeeRequest) (dto.EmployeeResponse, error) {
	employeeData, _ := json.Marshal(employee)
	req, err := http.NewRequest("POST", employeeURI+"/api/v1/employees", bytes.NewBuffer(employeeData))
	if err != nil {
		return dto.EmployeeResponse{}, err
	}
	
	req.Header.Set("Content-Type", "application/json")
	
	req.Header.Set("access_token", accessToken)
	
	client := &http.Client{}
	
	resp, err := client.Do(req)
	
	if err != nil {
		return dto.EmployeeResponse{}, err
	}
	
	defer resp.Body.Close()
	
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return dto.EmployeeResponse{}, err
	}
	
	var employees dto.EmployeeResponse
	
	if err = json.Unmarshal(responseBody, &employees); err != nil {
		return dto.EmployeeResponse{}, err
	}
	
	return employees, nil
	
}
