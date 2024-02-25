package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"service-user/config"
	"service-user/repository"
	"service-user/routes"
	"service-user/service"
)

func main() {
	dbConnection := config.GetConnection()
	
	userRepository := repository.NewUserRepository(dbConnection)
	
	userService := service.NewUserService(userRepository)
	
	app := fiber.New()
	app.Use(cors.New())
	
	routes.NewUserRoutes(app, userService)
	
	log.Fatalln(app.Listen(":3001"))
	
}
