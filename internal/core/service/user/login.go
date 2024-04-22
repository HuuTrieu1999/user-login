package user

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
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

func (u userService) Login(request request.LoginRequest) *response2.Response {
	// validate request
	if request.Account == "" {
		return u.createFailedResponse(error_code.InvalidAccountError, error_code.InvalidAccountMsg)
	}
	if request.Password == "" {
		return u.createFailedResponse(error_code.InvalidPasswordError, error_code.InvalidPasswordMsg)
	}

	// get from database
	var user entity.UserEntity
	var err error
	if phoneRegex.MatchString(request.Account) {
		user, err = u.userRepo.QueryByPhoneNumber(request.Account)
	} else if emailRegex.MatchString(request.Account) {
		user, err = u.userRepo.QueryByEmail(request.Account)
	} else {
		user, err = u.userRepo.QueryByUserName(request.Account)
	}
	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return u.createFailedResponse(error_code.AccountNotExistError, error_code.AccountNotExistMsg)
	} else if err != nil {
		log.Printf("fail to query user %s", err)
		return u.createFailedResponse(error_code.InternalError, error_code.InternalErrMsg)
	}

	// compare password
	if bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(request.Password)) != nil {
		return u.createFailedResponse(error_code.WrongPasswordError, error_code.WrongPasswordMsg)
	}

	token, err := u.createToken(user)
	if err != nil {
		log.Printf("fail to create token %s", err)
		return u.createFailedResponse(error_code.InternalError, error_code.InternalErrMsg)
	}

	// return success response
	return u.createSuccessResponse(response2.LoginDataResponse{
		Token: token,
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

func (u userService) createToken(user entity.UserEntity) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userName":    user.UserName,
			"userID":      user.UserID,
			"email":       user.Email,
			"phoneNumber": user.PhoneNumber,
			"exp":         time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte(u.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
