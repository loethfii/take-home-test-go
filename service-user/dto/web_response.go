package dto

type WebResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type ResponseError struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type LoginSuccessResponse struct {
	Code        int          `json:"code"`
	Status      string       `json:"status"`
	AccessToken string       `json:"access_token"`
	Data        UserResponse `json:"data"`
}
