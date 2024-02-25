package dto

type UserRequest struct {
	Email    string `json:"email" validate:"required,email,omitempty"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponseLogin struct {
	Code        int    `json:"code"`
	Status      string `json:"status"`
	AccessToken string `json:"access_token"`
	Data        User   `json:"data"`
}

type UserResponseRegister struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   string `json:"data"`
}

//{
//"code": 200,
//"status": "OK",
//"data": "forman222@gmail.com"
//}
