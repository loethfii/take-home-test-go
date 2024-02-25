package repository

import (
	"context"
	"gorm.io/gorm"
	"service-employee/domain"
)

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) domain.EmployeeRepository {
	return &employeeRepository{db}
}

func (r *employeeRepository) FindAllEmployee(ctx context.Context) ([]domain.Employee, error) {
	var employees []domain.Employee
	err := r.db.WithContext(ctx).Order("id desc").Find(&employees).Error
	if err != nil {
		return nil, err
	}
	
	return employees, nil
}

func (r *employeeRepository) LastIDEmployee(ctx context.Context) (uint64, error) {
	var employee domain.Employee
	result := r.db.WithContext(ctx).Last(&employee)
	if result.Error != nil {
		return 0, result.Error
	}
	
	return employee.ID, nil
}

func (r *employeeRepository) FindOneEmployee(ctx context.Context, id uint64) (domain.Employee, error) {
	var employee domain.Employee
	err := r.db.WithContext(ctx).First(&employee, id).Error
	if err != nil {
		return domain.Employee{}, err
	}
	
	return employee, nil
}

func (r *employeeRepository) CreateEmployee(ctx context.Context, employee domain.Employee) error {
	err := r.db.WithContext(ctx).Create(&employee).Error
	if err != nil {
		return err
	}
	
	return nil
}
