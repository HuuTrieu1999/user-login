package repository

import (
	"errors"
	"login/internal/core/entity"
)

var (
	DuplicateUser = errors.New("duplicate user")
)

type UserRepository interface {
	QueryByUserName(userName string) (entity.UserEntity, error)
	QueryByEmail(email string) (entity.UserEntity, error)
	QueryByPhoneNumber(phoneNumber string) (entity.UserEntity, error)
	Insert(userEntity entity.UserEntity) error
}
