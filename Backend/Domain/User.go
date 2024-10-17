package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	FullName  string             `json:"fullName" bson:"fullName" binding:"required"`
	ID         primitive.ObjectID `json:"id" bson:"_id" `
	Email      string             `json:"email" bson:"email" binding:"required, email"`
	Password   string             `json:"password" bson:"password" binding:"required, min=8 max=20"`
	Role       string             `json:"userType" bson:"userType" binding:"required"`
	IsVerified bool 			  `json:"is_verified" bson:"is_verified"`
	Created_at time.Time          `json:"created_at" bson:"created_at"`
	Updated_at time.Time          `json:"updated_at" bson:"updated_at"`
	VerificationToken string      `json:"verification_token" bson:"verification_token"`
	VerificationTokenExpire time.Time `json:"token_expire" bson:"token_expires"`
}

type UserUseCaseInterface interface {
	SignUp(User) (bool, error)
	FindEmail(string) error
}

type UserRepoInterface interface {
	FindUserByEmail(string) error
	CreateUser(*User) error
	SaveUnverifiedUser(*User) error
	FindUnverifiedUserByEmail(string) (User, error)
	UpdateUnverifiedUser(string, time.Time, string, time.Time)error
}