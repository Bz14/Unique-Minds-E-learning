package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	FullName  string             `json:"full_name" bson:"full_name" binding:"required"`
	ID         primitive.ObjectID `json:"id" bson:"_id" `
	Email      string             `json:"email" bson:"email" binding:"required, email"`
	Password   string             `json:"password" bson:"password" binding:"required, min=8 max=20"`
	Role       string             `json:"role" bson:"role" binding:"required"`
	IsVerified bool 			  `json:"is_verified" bson:"is_verified"`
	Created_at time.Time          `json:"created_at" bson:"created_at"`
	Updated_at time.Time          `json:"updated_at" bson:"updated_at"`
}

type UserUseCaseInterface interface {
	SignUp(user User) error
}

type UserRepoInterface interface {
	FindUserByEmail(email string) error
	CreateUser(user *User) error
}