package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEntity struct {
	UserID       primitive.ObjectID `json:"userID" bson:"_id"`
	UserName     string             `bson:"userName"`
	FullName     string             `bson:"fullName"`
	Email        string             `bson:"email"`
	PhoneNumber  string             `bson:"phoneNumber"`
	HashPassword string             `bson:"hashPassword"`
	Birthday     string             `bson:"birthday"`
}
