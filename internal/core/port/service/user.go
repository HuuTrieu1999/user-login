package service

import (
	"login/internal/controller/http/request"
	"login/internal/controller/http/response"
)

type UserService interface {
	Login(request request.LoginRequest) *response.Response
	Register(request request.RegisterRequest) *response.Response
}
