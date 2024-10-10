package model

type LoginResponse struct {
	Code    int32     `json:"code"`
	Message string    `json:"message"`
	Data    DataLogin `json:"data"`
}

type DataLogin struct {
	SessionId string `json:"session_id"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
