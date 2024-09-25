package domain

import "time"

type User struct {
	ID         primitive.ObjectId `json:"id" bson:"_id"`
	Email      string             `json:"email" bson:"email"`
	Password   string             `json:"password" bson:"password"`
	Role       string             `json:"role" bson:"role"`
	Created_at time.Time          `json:"created_at" bson:"created_at"`
	Updated_at time.Time          `json:"updated_at" bson:"updated_at"`
}