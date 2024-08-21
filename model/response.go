package model

type Response struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}
