package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"login/internal/core/entity"
	"login/internal/core/port/repository"
)

type userRepo struct {
	col *mongo.Collection
}

func (u userRepo) QueryByUserName(userName string) (entity.UserEntity, error) {
	var user entity.UserEntity
	filter := bson.M{"userName": userName}
	err := u.col.FindOne(context.Background(), filter).Decode(&user)
	return user, err
}

func (u userRepo) QueryByEmail(email string) (entity.UserEntity, error) {
	var user entity.UserEntity
	filter := bson.M{"email": email}
	err := u.col.FindOne(context.Background(), filter).Decode(&user)
	return user, err
}

func (u userRepo) QueryByPhoneNumber(phoneNumber string) (entity.UserEntity, error) {
	var user entity.UserEntity
	filter := bson.M{"phoneNumber": phoneNumber}
	err := u.col.FindOne(context.Background(), filter).Decode(&user)
	return user, err
}

func (u userRepo) Insert(userEntity entity.UserEntity) error {
	_, err := u.col.InsertOne(context.Background(), userEntity)
	return err
}

func NewUserRepo(col *mongo.Collection) repository.UserRepository {
	return &userRepo{
		col: col,
	}
}
