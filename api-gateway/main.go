package main

import (
	_ "api-gateway/docs"
	"api-gateway/routes"
	"api-gateway/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"log"
)

// @title 	Api Gateway | service user & service employee
// @version	1.0.0
// @description Api Gateway with fiber, part of service user & service employee with postgresql

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name access_token

// @host 	localhost:3000
// @BasePath /api/v1
func main() {
	employeeService := service.NewEmployeeService()
	userService := service.NewUserService()
	
	app := fiber.New()
	app.Use(cors.New())
	
	routes.NewEmployeeRoutes(app, employeeService)
	routes.NewUserRoutes(app, userService)
	
	app.Get("/swagger/*", swagger.HandlerDefault) // default
	
	log.Fatalln(app.Listen(":3000"))
}
