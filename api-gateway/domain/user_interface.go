package domain

import "api-gateway/dto"

type UserService interface {
	Login(userBody dto.UserRequest) (map[string]any, error)
	Register(user dto.UserRequest) (map[string]any, error)
}
