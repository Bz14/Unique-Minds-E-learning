package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Course struct

type Course struct {
    ID          primitive.ObjectID    `json:"_id" bson:"_id"`
    Name        string                `json:"name" bson:"name"`
    Description string                `json:"description" bson:"description"`
    Image       string                `json:"image" bson:"image"`
    Parts       []Part               `json:"parts" bson:"parts"`
    CreatedDate time.Time             `json:"created_date" bson:"created_date"`
    LastUpdated time.Time             `json:"last_updated" bson:"last_updated"`
	IsFeatured  bool                  `json:"is_featured" bson:"is_featured"`
    Creator_id  primitive.ObjectID    `json:"user_id" bson:"user_id"`
    Tags        []string              `json:"tags" bson:"tags"`
    Count       int                   `json:"count" bson:"count"`
    Students    []primitive.ObjectID  `json:"students" bson:"students"`
}

// Part struct
type Part struct {
    ID          primitive.ObjectID    `json:"_id" bson:"_id"`
    Name        string    `json:"name" bson:"name"`
    Description string    `json:"description" bson:"description"`
    Materials   []Material `json:"materials" bson:"materials"`
    Sequence    int    `json:"sequence" bson:"sequence"`
    CreatedDate time.Time `json:"created_date" bson:"created_date"`
    LastUpdated time.Time `json:"last_updated" bson:"last_updated"`
}

// Material struct
type Material struct {
    ID          primitive.ObjectID    `json:"_id" bson:"_id"`
    Name        string    `json:"name" bson:"name"`
    Type        string    `json:"type" bson:"type"`
    Content     string    `json:"content" bson:"content"`
    Description string    `json:"description" bson:"description"`
    CreatedDate time.Time `json:"created_date" bson:"created_date"`
    LastUpdated time.Time `json:"last_updated" bson:"last_updated"`
}

// CourseRepository interface
type CourseRepository interface {
    FetchRecentCourses() ([]Course, error)
    GetCourses(pageNo int64, pageSize int64, search string, tag string) ([]Course, Pagination, error)
    GerCourseById(id string) (Course, error)
    SaveCourse(userID string, courseID string) error
    GetMyCourse(id string) ([]Course, error)
    Save(course *Course, user_id string) error
    DeleteCourse(courseID string) error
    GetCoursesByEducator(userID string) ([]Course, error)
    UpdateCourseProgress(courseID, userID string, completedParts []string) (CourseProgress, error)
    GetCourseProgress(courseID, userID string) (*CourseProgress, error) 
}
// CourseUseCaseInterface interface
type CourseUseCaseInterface interface {
    GetRecentCourses() ([]Course, error)
    GetCourses(pageNo string, pageSize string, search string, filter string) ([]Course, Pagination, error)
    SaveCourse(studentID string, courseID string) error
    GetMyCourses(id string) ([]Course, error)
    UploadCourse(course *Course, user_id string) error
    GetEducatorCourses(id string) ([]Course, error)
    DeleteCourse(id string) error
    UpdateProgress(courseID, userID string, completedParts []string) (CourseProgress, error)
    GetCourseByID(courseID, userID string) (*CourseDetailResponse, error)
}