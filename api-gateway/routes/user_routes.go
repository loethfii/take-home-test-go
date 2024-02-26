package routes

import (
	"api-gateway/domain"
	"api-gateway/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type userRoutes struct {
	userService domain.UserService
}

func NewUserRoutes(app *fiber.App, userService domain.UserService) {
	handler := &userRoutes{userService}
	g := app.Group("/api/v1")
	g.Post("/login", handler.Login)
	g.Post("/register", handler.Register)
}

func (r *userRoutes) Login(ctx *fiber.Ctx) error {
	var reqLogin dto.UserRequest
	err := ctx.BodyParser(&reqLogin)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    "400",
			"Status":  "Bad Request",
			"Message": err.Error(),
		})
	}
	
	validate := validator.New()
	
	if err := validate.Struct(reqLogin); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{
			Code:    400,
			Status:  "Bad Request",
			Message: validationErrors.Error(),
		})
	}
	
	loginResponse, err := r.userService.Login(reqLogin)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Code":    "500",
			"Status":  "Internal Server Error",
			"Message": err.Error(),
		})
	}
	
	msg, _ := loginResponse["message"]
	if msg == "email or password is incorrect" {
		return ctx.Status(fiber.StatusForbidden).JSON(loginResponse)
	}
	
	return ctx.Status(fiber.StatusOK).JSON(loginResponse)
	
}

func (r *userRoutes) Register(ctx *fiber.Ctx) error {
	var reqRegister dto.UserRequest
	err := ctx.BodyParser(&reqRegister)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    "400",
			"Status":  "Bad Request",
			"Message": err.Error(),
		})
	}
	
	validate := validator.New()
	
	if err := validate.Struct(reqRegister); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{
			Code:    400,
			Status:  "Bad Request",
			Message: validationErrors.Error(),
		})
	}
	
	registerResponse, err := r.userService.Register(reqRegister)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Code":    "500",
			"Status":  "Internal Server Error kontol",
			"Message": err.Error(),
		})
	}
	
	msg, _ := registerResponse["message"]
	
	if msg == "email already exist" {
		return ctx.Status(fiber.StatusConflict).JSON(registerResponse)
	}
	
	return ctx.Status(fiber.StatusCreated).JSON(registerResponse)
}
