package config

import (
	"gorm.io/gorm"
	"service-user/domain"
)

func dbMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&domain.User{})
	
}
