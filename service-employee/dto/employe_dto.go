package dto

type EmployeeResponse struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type EmployeeRequest struct {
	Name string `json:"name" validate:"required"`
}
