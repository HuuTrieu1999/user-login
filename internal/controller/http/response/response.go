package response

import (
	"login/internal/core/dto"
	"login/internal/core/entity/error_code"
)

type Response struct {
	Data         interface{}          `json:"data"`
	Status       bool                 `json:"status"`
	ErrorCode    error_code.ErrorCode `json:"errorCode"`
	ErrorMessage string               `json:"errorMessage"`
}

type LoginDataResponse struct {
	Token string      `json:"token"`
	User  dto.UserDTO `json:"user"`
}

type RegisterDataResponse struct {
	User dto.UserDTO `json:"user"`
}
