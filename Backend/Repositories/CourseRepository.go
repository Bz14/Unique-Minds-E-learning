package repository

import (
	"context"
	"errors"
	"math"
	"time"
	domain "unique-minds/Domain"
	infrastructure "unique-minds/Infrastructure"
	utils "unique-minds/Utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CourseRepository struct {
    collection *mongo.Collection
	config   *infrastructure.Config
	userCollection *mongo.Collection
	studentColl  *mongo.Collection
}

func NewCourseRepository(collection *mongo.Collection, eduColl *mongo.Collection, studColl *mongo.Collection, config *infrastructure.Config) *CourseRepository{
    return &CourseRepository{
		collection: collection,
		userCollection : eduColl,
		config: config,
		studentColl: studColl,
    }
}


func (r *CourseRepository) Save(course *domain.Course, user_id string) error {
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
	// Initialize pagination
	pagination := utils.PaginationByPage(pageNo, pageSize)

	// Count total documents
	totalResults, err := r.collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return []domain.Course{}, domain.Pagination{}, err
	}

	// Construct the filter
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

	// Decode results
	var courses []domain.Course
	for cursor.Next(context.TODO()) {
		var course domain.Course
		if err := cursor.Decode(&course); err != nil {
			return []domain.Course{}, domain.Pagination{}, err
		}
		courses = append(courses, course)
	}

	// Handle cursor errors
	if err := cursor.Err(); err != nil {
		return []domain.Course{}, domain.Pagination{}, err
	}

	// Close the cursor
	cursor.Close(context.TODO())

	// Create pagination info
	paginationInfo := domain.Pagination{
		CurrentPage: pageNo,
		PageSize:    pageSize,
		TotalPages:  totalPages,
		TotatRecord: totalResults,
	}
	return courses, paginationInfo, nil
}


func (r *CourseRepository) GetMyCourse(id string) ([]domain.Course, error) {
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

func (r *CourseRepository) GerCourseById(id string) (domain.Course, error) {
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

func (r *CourseRepository) GetCourseProgress(courseID, userID string) (*domain.CourseProgress, error) {
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

func (r *CourseRepository) UpdateCourseProgress(courseID, userID string, completedParts []string) (domain.CourseProgress, error) {
	cid, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return domain.CourseProgress{}, err
	}
	uid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return domain.CourseProgress{}, err
	}

	course, err := r.GerCourseById(cid.Hex())
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

func (r *CourseRepository) SaveCourse(userID string, courseID string) error {
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