package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserProfile struct {
	ID         		primitive.ObjectID 	  `bson:"_id,omitempity" json:"id" `
	FullName   		string             	  `bson:"name" json:"name"`
	User_Name  		string             	  `bson:"username"  json:"username"`
	Email      		string             	  `bson:"email" validate:"required,email" json:"email"`
	Contact    		string             	  `bson:"contact" json:"contact"`
	Dob         	string                `bson:"dob" json:"dob"`
	Phone   		string                `bson:"phone" json:"phone"`
	Address     	Address               `bson:"address" json:"address"`
	Bio         	string                `bson:"bio" json:"bio"`
	ProfileImage 	string                `bson:"profileImage" json:"profileImage"`
}