package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func GetConnection() *gorm.DB {
	//var dsn = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", "localhost", 5432, "postgres", "service-user", "root")
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error Database : ", err)
	}
	
	err = dbMigrate(db)
	if err != nil {
		log.Fatalln("Error Database : ", err)
	}
	
	return db
}
