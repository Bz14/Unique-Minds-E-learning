package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Email      string             `json:"email" bson:"email"`
	Password   string             `json:"password" bson:"password"`
	Role       string             `json:"role" bson:"role"`
	IsVerified bool 			  `json:"is_verified" bson:"is_verified"`
	Created_at time.Time          `json:"created_at" bson:"created_at"`
	Updated_at time.Time          `json:"updated_at" bson:"updated_at"`
}

type UserUseCaseInterface interface {
	SignUp(signUpRequest SignUpRequest) error
}

type UserRepoInterface interface {
	FindUserByEmail(email string) error
}