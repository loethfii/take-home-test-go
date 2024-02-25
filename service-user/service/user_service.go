package service

import (
	"context"
	"errors"
	"service-user/domain"
	"service-user/dto"
	"service-user/helper"
)

type userService struct {
	userRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) domain.UserService {
	return &userService{userRepository}
}

func (u *userService) Login(userRequest dto.UserRequest) (dto.LoginResponse, error) {
	ctx := context.Background()
	
	var userData = domain.User{
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}
	
	user, err := u.userRepository.Login(ctx, &userData)
	if err != nil {
		return dto.LoginResponse{}, errors.New("email or password is incorrect")
	}
	
	isSuccess := helper.ComparePassword([]byte(user.Password), []byte(userRequest.Password))
	if !isSuccess {
		return dto.LoginResponse{}, errors.New("email or password is incorrect")
	}
	
	token := helper.GenerateToken(userData)
	
	return dto.LoginResponse{
		AccessToken: token,
		Data: dto.UserResponse{
			ID:       user.ID,
			Email:    user.Email,
			Password: user.Password,
		},
	}, nil
}

func (u *userService) Register(userRequest dto.UserRequest) (string, error) {
	ctx := context.Background()
	
	hashPassword := helper.HashPassword([]byte(userRequest.Password))
	
	var regUser = domain.User{
		Email:    userRequest.Email,
		Password: hashPassword,
	}
	
	email, err := u.userRepository.Register(ctx, &regUser)
	if err != nil {
		return "", err
	}
	
	return email, nil
}
