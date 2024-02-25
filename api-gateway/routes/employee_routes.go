package routes

import (
	"api-gateway/domain"
	"api-gateway/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type employeeRoutes struct {
	employeeService domain.EmployeeService
}

func NewEmployeeRoutes(app *fiber.App, employeeService domain.EmployeeService) {
	handler := &employeeRoutes{employeeService}
	g := app.Group("/api/v1")
	//app.Get("/", handler.HomeEmployee)
	g.Get("/employees", handler.FetchEmployee)
	g.Get("/employees/:id", handler.DetailEmployee)
	g.Post("/employees", handler.AddEmployee)
}

func (r *employeeRoutes) FetchEmployee(ctx *fiber.Ctx) error {
	accessToken := ctx.Get("access_token")
	employeesRes, err := r.employeeService.FetchEmployee(accessToken)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Code":    "500",
			"Status":  "Internal Server Error",
			"Message": err.Error(),
		})
	}
	
	return ctx.Status(fiber.StatusOK).JSON(employeesRes)
	
}

func (r *employeeRoutes) DetailEmployee(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    "400",
			"Status":  "Bad Request",
			"Message": err.Error(),
		})
	}
	
	accessToken := ctx.Get("access_token")
	
	employeeRes, err := r.employeeService.DetailEmployee(accessToken, uint64(idInt))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Code":    "500",
			"Status":  "Internal Server Error",
			"Message": err.Error(),
		})
	}
	
	return ctx.Status(fiber.StatusOK).JSON(employeeRes)
}

func (r *employeeRoutes) AddEmployee(ctx *fiber.Ctx) error {
	
	var employeeReq dto.EmployeeRequest
	
	if err := ctx.BodyParser(&employeeReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    "400",
			"Status":  "Bad Request",
			"Message": err.Error(),
		})
	}
	
	validate := validator.New()
	
	if err := validate.Struct(employeeReq); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{
			Code:    400,
			Status:  "Bad Request",
			Message: validationErrors.Error(),
		})
	}
	
	accessToken := ctx.Get("access_token")
	employeeRes, err := r.employeeService.AddEmployee(accessToken, employeeReq)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Code":    "500",
			"Status":  "Internal Server Error",
			"Message": err.Error(),
		})
	}
	
	return ctx.Status(fiber.StatusOK).JSON(employeeRes)
}
