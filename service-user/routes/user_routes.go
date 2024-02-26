package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"service-user/domain"
	"service-user/dto"
	"service-user/helper"
)

type userRoutes struct {
	userService domain.UserService
}

func NewUserRoutes(app *fiber.App, userService domain.UserService) {
	handler := &userRoutes{userService}
	
	g := app.Group("/api/v1/user")
	app.Get("/", helper.TokenMiddleware(handler.HomeUser))
	g.Post("/register", handler.RegisterUser)
	g.Post("/login", handler.LoginUser)
}

func (r *userRoutes) HomeUser(ctx *fiber.Ctx) error {
	cekLocals := ctx.Locals("email")
	
	if emailMap, ok := cekLocals.(map[string]string); ok {
		// Mengakses nilai dari map dengan key "email"
		email := emailMap["email"]
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":     "Hello from user service",
			"auth status": "Authentication OK",
			"who_login":   email,
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello from user service",
		})
	}
	
}

func (r *userRoutes) RegisterUser(ctx *fiber.Ctx) error {
	var userRequest dto.UserRequest
	if err := ctx.BodyParser(&userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{
			Code:    400,
			Status:  "Bad Request",
			Message: err.Error(),
		})
	}
	
	validate := validator.New()
	
	if err := validate.Struct(userRequest); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{
			Code:    400,
			Status:  "Bad Request",
			Message: validationErrors.Error(),
		})
	}
	
	userRes, err := r.userService.Register(userRequest)
	if err != nil {
		if err.Error() == "email already exist" {
			return ctx.Status(fiber.StatusConflict).JSON(dto.ResponseError{
				Code:    409,
				Status:  "Conflict",
				Message: err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ResponseError{
			Code:    500,
			Status:  "Internal Server Error",
			Message: err.Error(),
		})
	}
	
	return ctx.Status(fiber.StatusCreated).JSON(dto.WebResponse{
		Code:   201,
		Status: "created",
		Data:   userRes,
	})
}

func (r *userRoutes) LoginUser(ctx *fiber.Ctx) error {
	var requestUser dto.UserRequest
	
	if err := ctx.BodyParser(&requestUser); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{
			Code:    400,
			Status:  "Bad Request",
			Message: err.Error(),
		})
	}
	
	validate := validator.New()
	
	if err := validate.Struct(requestUser); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{
			Code:    400,
			Status:  "Bad Request",
			Message: validationErrors.Error(),
		})
	}
	
	user, err := r.userService.Login(requestUser)
	if err != nil {
		if err.Error() == "email or password is incorrect" {
			return ctx.Status(fiber.StatusForbidden).JSON(dto.ResponseError{
				Code:    403,
				Status:  "Forbidden",
				Message: err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ResponseError{
			Code:    500,
			Status:  "Internal Server Error",
			Message: err.Error(),
		})
	}
	
	return ctx.Status(fiber.StatusOK).JSON(dto.LoginSuccessResponse{
		Code:        200,
		Status:      "ok",
		AccessToken: user.AccessToken,
		Data:        user.Data,
	})
	
}
