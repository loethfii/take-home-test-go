package service

import (
	"api-gateway/domain"
	"api-gateway/dto"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type userService struct {
}

func NewUserService() domain.UserService {
	return &userService{}
}

var userURI = "http://user:3001"

//var userURI = "http://localhost:3001"

// CreateTags		godoc
// @Summary			Login user.
// @Description		Authentication user for access fitur.
// @Produce			application/json
// @Param			tags body dto.UserRequest true "request body for login user"
// @Tags			User
// @Success			200 {object} dto.UserResponseLogin{}
// @Router			/login [post]
func (u *userService) Login(userBody dto.UserRequest) (dto.UserResponseLogin, error) {
	userData, _ := json.Marshal(userBody)
	req, err := http.NewRequest("POST", userURI+"/api/v1/user/login", bytes.NewBuffer(userData))
	if err != nil {
		return dto.UserResponseLogin{}, err
	}
	
	req.Header.Set("Content-Type", "application/json")
	
	client := &http.Client{}
	
	resp, err := client.Do(req)
	if err != nil {
		return dto.UserResponseLogin{}, err
	}
	
	defer resp.Body.Close()
	
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return dto.UserResponseLogin{}, err
	}
	
	var user dto.UserResponseLogin
	if err = json.Unmarshal(responseBody, &user); err != nil {
		return dto.UserResponseLogin{}, err
	}
	
	return user, nil
}

// CreateTags		godoc
// @Summary			Registe user.
// @Description		Registe user for new user.
// @Produce			application/json
// @Param			tags body dto.UserRequest true "request body for register user"
// @Tags			User
// @Success			200 {object} dto.UserResponseLogin{}
// @Router			/register [post]
func (u *userService) Register(userBody dto.UserRequest) (map[string]any, error) {
	userData, _ := json.Marshal(userBody)
	req, err := http.NewRequest("POST", userURI+"/api/v1/user/register", bytes.NewBuffer(userData))
	if err != nil {
		return nil, err
	}
	
	req.Header.Set("Content-Type", "application/json")
	
	client := &http.Client{}
	
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	
	defer resp.Body.Close()
	
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	var user map[string]any
	if err = json.Unmarshal(responseBody, &user); err != nil {
		return nil, err
	}
	
	msg, _ := user["message"].(string)
	code, _ := user["code"]
	status, _ := user["status"].(string)
	data, _ := user["data"].(string)
	if msg == "" {
		return map[string]any{
			"code":   code,
			"status": status,
			"data":   data,
		}, nil
	}
	
	return map[string]any{
		"code":    code,
		"status":  status,
		"message": msg,
	}, nil
}
