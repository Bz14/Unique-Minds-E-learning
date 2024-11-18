package repository

import (
	"context"
	"errors"
	domain "unique-minds/Domain"

	infrastructure "unique-minds/Infrastructures"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type StudentCourseRepository struct {
    collection *mongo.Collection
	config   *infrastructure.Config
	studentColl  *mongo.Collection
}
func NewStudentCourseRepository(collection *mongo.Collection, studColl *mongo.Collection, config *infrastructure.Config) *StudentCourseRepository{
    return &StudentCourseRepository{
		collection: collection,
		config: config,
		studentColl: studColl,
    }
}

func (r *StudentCourseRepository) GetMyCourse(id string) ([]domain.Course, error) {
	var courses []domain.Course
	var student domain.StudentProfile
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = r.studentColl.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&student)
	if err != nil {
		return nil, err
	}

	for id := range student.CourseIds {
		var course domain.Course
		err = r.collection.FindOne(context.TODO(), bson.M{"_id": student.CourseIds[id]}).Decode(&course)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func (r *StudentCourseRepository) GetCourseProgress(courseID, userID string) (*domain.CourseProgress, error) {
	var user domain.StudentProfile

	cid, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, errors.New("invalid course ID")
	}
	uid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	err = r.studentColl.FindOne(context.TODO(), bson.M{"_id": uid}).Decode(&user)
	if err != nil {
		return nil, err
	}
	for _, courseProgress := range user.EnrolledCourses {
		if courseProgress.CourseID == cid {
			return &courseProgress, nil
		}
	}
	return nil, errors.New("course progress not found")
}

func (r *StudentCourseRepository) UpdateCourseProgress(courseID, userID string, completedParts []string) (domain.CourseProgress, error) {
	cid, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return domain.CourseProgress{}, err
	}
	uid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return domain.CourseProgress{}, err
	}

	course, err := r.GetCourseById(cid.Hex())
	if err != nil{
		return domain.CourseProgress{}, err
	}

	total_parts := len(course.Parts)

	filter := bson.M{"_id": uid, "courses._id": cid}
	update := bson.M{
		"$set": bson.M{
			"courses.$.progress":       int64(len(completedParts)) * 100 / int64(total_parts),
			"courses.$.completed_parts": completedParts,
			"courses.$.is_completed":    len(completedParts) == total_parts,
		},
	}
	_, err = r.studentColl.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return domain.CourseProgress{}, err
	}
	var student domain.StudentProfile
	err = r.studentColl.FindOne(context.TODO(), bson.M{"_id":uid}).Decode(&student)
	for _, course := range student.EnrolledCourses{
		if course.CourseID == cid{
			return course, nil
		}
	}
	return domain.CourseProgress{}, errors.New("Progress not found")
}

func (r *StudentCourseRepository) GetCourseById(id string) (domain.Course, error) {
	var course domain.Course

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Course{}, err
	}

	filter := bson.M{"_id": objID}
	err = r.collection.FindOne(context.TODO(), filter).Decode(&course)
	if err != nil { 
		return domain.Course{}, err
	}

	return course, nil
}


func (r *StudentCourseRepository) SaveCourse(userID string, courseID string) error {
	studentObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("invalid student ID")
	}
	courseObjID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return errors.New("invalid course ID")
	}

	initialProgress := domain.CourseProgress{
		CourseID:       courseObjID,
		Progress:       0, 
		CompletedParts: []primitive.ObjectID{},
		IsCompleted:    false,
	}

	filter := bson.M{"_id": studentObjID}
	update := bson.M{
		"$push": bson.M{
			"course_id":    courseObjID,      
			"courses":      initialProgress, 
		},
	}

	_, err = r.studentColl.UpdateOne(context.TODO(), filter, update)
	if err != nil{
		return err
	}
	_, err = r.collection.UpdateOne(context.TODO(), bson.M{"_id" : courseObjID},  bson.M{
        "$inc": bson.M{"count": 1},
    })
	if err != nil{
		return err
	}
	_, err = r.collection.UpdateOne(context.TODO(), bson.M{"_id" : courseObjID},  bson.M{"$push": bson.M{"students":   studentObjID,}})
	return err
}