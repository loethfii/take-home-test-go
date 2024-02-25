package domain

import (
	"context"
	"service-user/dto"
)

type UserRepository interface {
	Login(ctx context.Context, user *User) (User, error)
	Register(ctx context.Context, user *User) (string, error)
}

type UserService interface {
	Login(userRequest dto.UserRequest) (dto.LoginResponse, error)
	Register(userRequest dto.UserRequest) (string, error)
}
