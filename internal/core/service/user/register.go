package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"log"
	"login/internal/controller/http/request"
	response2 "login/internal/controller/http/response"
	"login/internal/core/dto"
	"login/internal/core/entity"
	"login/internal/core/entity/error_code"
	"time"
)

func (u userService) Register(request request.RegisterRequest) *response2.Response {
	// validate request
	if request.FullName == "" {
		return u.createFailedResponse(error_code.InvalidFullNameError, error_code.InvalidFullNameMsg)
	}
	if request.UserName == "" && request.Email == "" && request.PhoneNumber == "" {
		return u.createFailedResponse(error_code.InvalidUserNameError, error_code.InvalidUserNameMsg)
	}
	if request.Email != "" && !emailRegex.MatchString(request.Email) {
		return u.createFailedResponse(error_code.InvalidEmailError, error_code.InvalidEmailMsg)
	}
	if request.PhoneNumber != "" && !phoneRegex.MatchString(request.PhoneNumber) {
		return u.createFailedResponse(error_code.InvalidPhoneNumberError, error_code.InvalidPhoneNumberMsg)
	}
	if request.UserName != "" && !usernameRegex.MatchString(request.UserName) {
		return u.createFailedResponse(error_code.InvalidUserNameError, error_code.InvalidUserNameMsg)
	}
	if request.Password == "" {
		return u.createFailedResponse(error_code.InvalidPasswordError, error_code.InvalidPasswordMsg)
	}
	if request.Birthday == "" {
		return u.createFailedResponse(error_code.InvalidBirthdayError, error_code.InvalidBirthdayMsg)
	}

	hashPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("fail to hash password %s", err)
		return u.createFailedResponse(error_code.InternalError, error_code.InternalErrMsg)
	}
	// insert database
	user := entity.UserEntity{
		UserID:       primitive.NewObjectID(),
		UserName:     request.UserName,
		FullName:     request.FullName,
		Email:        request.Email,
		PhoneNumber:  request.PhoneNumber,
		HashPassword: string(hashPasswordBytes),
		Birthday:     request.Birthday,
		CreatedDate:  time.Now(),
		UpdatedDate:  time.Now(),
	}
	err = u.userRepo.Insert(user)
	if err != nil && mongo.IsDuplicateKeyError(err) {
		return u.createFailedResponse(error_code.AccountExistError, error_code.AccountExistMsg)
	} else if err != nil {
		log.Printf("fail to insert new user %s", err)
		return u.createFailedResponse(error_code.InternalError, error_code.InternalErrMsg)
	}

	// return success response
	return u.createSuccessResponse(response2.RegisterDataResponse{
		User: dto.UserDTO{
			UserID:      user.UserID.Hex(),
			UserName:    user.UserName,
			FullName:    user.FullName,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			Birthday:    user.Birthday,
		},
	})
}
