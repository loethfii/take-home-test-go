package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"service-user/domain"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db}
}

func (u userRepository) Login(ctx context.Context, user *domain.User) (domain.User, error) {
	var userRes domain.User
	err := u.db.WithContext(ctx).First(&userRes, "email = ?", user.Email).Error
	if err != nil {
		return userRes, err
	}
	
	return userRes, nil
}

func (u userRepository) Register(ctx context.Context, user *domain.User) (string, error) {
	row := u.db.WithContext(ctx).Where("email = ?", user.Email).First(&user)
	if row.RowsAffected > 0 {
		return "", errors.New("email already exist")
	}
	err := u.db.WithContext(ctx).Create(&user).Error
	
	if err != nil {
		return "", err
	}
	
	return user.Email, nil
}
