package repository

import (
	"context"
	"errors"
	"math"
	domain "unique-minds/Domain"
	utils "unique-minds/Utils"

	infrastructure "unique-minds/Infrastructures"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CourseRepository struct {
    collection *mongo.Collection
	config   *infrastructure.Config
}
func NewCourseRepository(collection *mongo.Collection, config *infrastructure.Config) *CourseRepository{
    return &CourseRepository{
		collection: collection,
		config: config,
    }
}
func (r *CourseRepository) FetchRecentCourses() ([]domain.Course, error) {
    var courses []domain.Course

    cur, err := r.collection.Find(context.TODO(), bson.D{}, options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}}).SetLimit(3))
    if err != nil {
        return nil, err
    }
    defer cur.Close(context.TODO())

    for cur.Next(context.TODO()) {
        var course domain.Course
        if err := cur.Decode(&course); err != nil {
            return nil, err
        }
        courses = append(courses, course)
    }

    if err := cur.Err(); err != nil {
        return nil, err
    }
	return courses, nil
}
func (r *CourseRepository) GetCourses(pageNo int64, pageSize int64, search string, tag string) ([]domain.Course, domain.Pagination, error) {
	pagination := utils.PaginationByPage(pageNo, pageSize)

	totalResults, err := r.collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return []domain.Course{}, domain.Pagination{}, err
	}

	filter := bson.M{}
	if search != "" {
		filter["name"] = bson.M{"$regex": search, "$options": "i"}
	}
	if tag != "" {
		filter["tags"] = bson.M{"$elemMatch": bson.M{"$regex": tag, "$options": "i"}}
	}

	totalPages := int64(math.Ceil(float64(totalResults) / float64(pageSize)))

	cursor, err := r.collection.Find(context.TODO(), filter, pagination)
	if err != nil {
		return []domain.Course{}, domain.Pagination{}, err
	}

	var courses []domain.Course
	for cursor.Next(context.TODO()) {
		var course domain.Course
		if err := cursor.Decode(&course); err != nil {
			return []domain.Course{}, domain.Pagination{}, err
		}
		courses = append(courses, course)
	}

	if err := cursor.Err(); err != nil {
		return []domain.Course{}, domain.Pagination{}, err
	}

	cursor.Close(context.TODO())

	paginationInfo := domain.Pagination{
		CurrentPage: pageNo,
		PageSize:    pageSize,
		TotalPages:  totalPages,
		TotalRecord: totalResults,
	}
	return courses, paginationInfo, nil
}
func (r *CourseRepository) GetCoursesByEducator(userID string) ([]domain.Course, error) {
	uid, _ := primitive.ObjectIDFromHex(userID)
    var courses []domain.Course
    filter := bson.M{"user_id": uid}

    cursor, err := r.collection.Find(context.TODO(), filter)
    if err != nil {
        return nil, errors.New("failed to get courses")
    }

    if err = cursor.All(context.TODO(), &courses); err != nil {
        return nil, errors.New("failed to get courses")
    }

    return courses, nil
}
func (r *CourseRepository) GetCourseById(id string) (domain.Course, error) {
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
