package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"service-employee/config"
	"service-employee/repository"
	"service-employee/routes"
	"service-employee/service"
)

func main() {
	dbConnection := config.GetConnection()
	
	employeeRepository := repository.NewEmployeeRepository(dbConnection)
	
	employeeService := service.NewEmployeeService(employeeRepository)
	
	app := fiber.New()
	app.Use(cors.New())
	
	routes.NewEmployeeRoutes(app, employeeService)
	
	log.Fatalln(app.Listen(":3002"))
}
