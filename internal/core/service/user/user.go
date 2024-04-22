package user

import (
	response2 "login/internal/controller/http/response"
	"login/internal/core/entity/error_code"
	"login/internal/core/port/repository"
	"login/internal/core/port/service"
	"regexp"
)

var (
	phoneRegex    = regexp.MustCompile("^\\d+$") // Basic phone number format (digits only)
	emailRegex    = regexp.MustCompile("^[a-z0-9._%+\\-]+@[a-z0-9.\\-]+\\.[a-z]{2,4}$")
	usernameRegex = regexp.MustCompile("^[a-zA-Z0-9_.-]+$")
)

type userService struct {
	userRepo  repository.UserRepository
	secretKey string
}

func NewUserService(userRepo repository.UserRepository, secretKey string) service.UserService {
	return &userService{
		userRepo:  userRepo,
		secretKey: secretKey,
	}
}

func (u userService) createFailedResponse(
	code error_code.ErrorCode, message string,
) *response2.Response {
	return &response2.Response{
		ErrorCode:    code,
		ErrorMessage: message,
		Status:       false,
	}
}

func (u userService) createSuccessResponse(data interface{}) *response2.Response {
	return &response2.Response{
		Data:         data,
		ErrorCode:    error_code.Success,
		ErrorMessage: error_code.SuccessErrMsg,
		Status:       true,
	}
}
