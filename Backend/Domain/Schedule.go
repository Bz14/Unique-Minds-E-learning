package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Schedule struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	StudentID      primitive.ObjectID `json:"student_id" bson:"student_id"`
	EducatorId     primitive.ObjectID `json:"educator_id" bson:"educator_id"`
	Date           string             `json:"date" bson:"date"`
	GoogleMeetLink string             `json:"googleMeetLink" bson:"googleMeetLink"`
}