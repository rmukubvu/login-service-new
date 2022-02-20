package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserName string             `json:"user_name" bson:"user_name"`
	Password string             `json:"password" bson:"password"`
	IsLocked bool               `json:"is_locked" bson:"is_locked"`
}

type UserLoginResponse struct {
	Id           string `json:"id"`
	UserName     string `json:"user_name"`
	IsError      bool   `json:"is_error"`
	ErrorMessage string `json:"error"`
}
