package domain

import "api-gateway/dto"

type UserService interface {
	Login(userBody dto.UserRequest) (dto.UserResponseLogin, error)
	Register(user dto.UserRequest) (map[string]any, error)
}
