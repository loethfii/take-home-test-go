package config

import (
	"gorm.io/gorm"
	"service-employee/domain"
)

func dbMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&domain.Employee{})
}
