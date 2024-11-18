package repository

import (
	"context"
	"errors"
	"time"
	domain "unique-minds/Domain"

	infrastructure "unique-minds/Infrastructures"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EducatorCourseRepository struct {
    collection *mongo.Collection
	config   *infrastructure.Config
	eduCollection *mongo.Collection
}

func NewEducatorCourseRepository(collection *mongo.Collection, eduColl *mongo.Collection, config *infrastructure.Config) *EducatorCourseRepository{
    return &EducatorCourseRepository{
		collection: collection,
		eduCollection : eduColl,
		config: config,
    }
}

func (r *EducatorCourseRepository) Save(course *domain.Course, user_id string) error {
	uid, _ := primitive.ObjectIDFromHex(user_id)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	course.ID = primitive.NewObjectID()
	course.CreatedDate = time.Now()
	course.LastUpdated = time.Now()
	course.Creator_id = uid

	for i := range course.Parts {
		course.Parts[i].ID = primitive.NewObjectID()
		course.Parts[i].CreatedDate = time.Now()
		course.Parts[i].LastUpdated = time.Now()

		for j := range course.Parts[i].Materials {
			course.Parts[i].Materials[j].ID = primitive.NewObjectID()
			course.Parts[i].Materials[j].CreatedDate = time.Now()
			course.Parts[i].Materials[j].LastUpdated = time.Now()
		}
	}

	filter := bson.M{"_id": course.ID}
	update := bson.M{"$set": course}
	opts := options.Update().SetUpsert(true)

	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

func (r *CourseRepository) DeleteCourse(courseID string) error {
	cid, _ := primitive.ObjectIDFromHex(courseID)
    filter := bson.M{"_id": cid, "count": 0}
    result, err := r.collection.DeleteOne(context.TODO(), filter)
    if err != nil {
        return errors.New("Students already enrolled in this course, cannot delete")
    }
    
    if result.DeletedCount == 0 {
        return errors.New("Students already enrolled in this course, cannot delete")
    }
    
    return nil
}