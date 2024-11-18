package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StudentProfile struct {
	ID              primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	FullName        string               `bson:"name" json:"name"`
	Age             string               `json:"age" bson:"age"`
	Bio             string               `json:"bio" bson:"bio"`
	GuardianEmail   string               `json:"guardianEmail" bson:"guardianEmail"`
	GuardianPhone   string               `json:"guardianPhone" bson:"guardianPhone"`
	Location        string               `json:"location" bson:"location"`
	ProfileImage    string               `json:"profileImage" bson:"profileImage"`
	UpdateAt        time.Time            `json:"updateAt" bson:"updateAt"`
	Created_At      time.Time            `bson:"created_at" json:"created_at"`
	CourseIds       []primitive.ObjectID `bson:"course_id" json:"course_id"`
	EnrolledCourses []CourseProgress     `bson:"courses" json:"courses"`
	Schedule        []Schedule           `bson:"schedules" json:"schedules"`
	Courses         []Course             `bson:"course_s" json:"course_s"`
	Condition       string               `bson:"condition" json:"condition"`
}