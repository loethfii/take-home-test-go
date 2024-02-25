package dto

type EmployeeRequest struct {
	Name string `json:"name" validate:"required"`
}

type Employee struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type EmployeeResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data"`
}
