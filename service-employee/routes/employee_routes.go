package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"service-employee/domain"
	"service-employee/dto"
	"service-employee/helper"
	"strconv"
)

type employeeRoutes struct {
	employeeService domain.EmployeeService
}

func NewEmployeeRoutes(app *fiber.App, employeeService domain.EmployeeService) {
	handler := &employeeRoutes{employeeService}
	g := app.Group("/api/v1")
	app.Get("/", handler.HomeEmployee)
	g.Get("/employees", helper.TokenMiddleware(handler.FetchEmployee))
	g.Get("/employees/:id", helper.TokenMiddleware(handler.DetailEmployee))
	g.Post("/employees", helper.TokenMiddleware(handler.AddEmployee))
}

func (r *employeeRoutes) HomeEmployee(ctx *fiber.Ctx) error {
	
	var listAPI = map[string]string{
		"GET http://localhost:3002/api/v1/employees":     "Fetch Employee",
		"GET http://localhost:3002/api/v1/employees/:id": "Detail Employee",
		"POST http://localhost:3002/api/v1/employees":    "Add Employee",
	}
	
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"List API": listAPI,
	})
}

func (r *employeeRoutes) FetchEmployee(ctx *fiber.Ctx) error {
	employees, err := r.employeeService.FetchEmployee()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ResponseError{
			Code:    500,
			Status:  "Internal Server Error",
			Message: err.Error(),
		})
	}
	
	return ctx.Status(fiber.StatusOK).JSON(dto.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   employees,
	})
}

func (r *employeeRoutes) DetailEmployee(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, _ := strconv.Atoi(id)
	employee, err := r.employeeService.DetailEmployee(uint64(idInt))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ResponseError{
			Code:    400,
			Status:  "Bad Request",
			Message: err.Error(),
		})
	}
	
	return ctx.Status(fiber.StatusOK).JSON(dto.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   employee,
	})
}

func (r *employeeRoutes) AddEmployee(ctx *fiber.Ctx) error {
	var employeeRequest dto.EmployeeRequest
	if err := ctx.BodyParser(&employeeRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{
			Code:    400,
			Status:  "Bad Request",
			Message: err.Error(),
		})
	}
	
	validate := validator.New()
	
	if err := validate.Struct(employeeRequest); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{
			Code:    400,
			Status:  "Bad Request",
			Message: validationErrors.Error(),
		})
	}
	
	employeeRes, err := r.employeeService.AddEmployee(employeeRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ResponseError{
			Code:    500,
			Status:  "Internal Server Error",
			Message: err.Error(),
		})
	}
	
	return ctx.Status(fiber.StatusOK).JSON(dto.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   employeeRes,
	})
}
