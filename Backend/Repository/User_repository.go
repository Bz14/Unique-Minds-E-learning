package repository

import (
	"context"
	"errors"
	"math"
	"strings"
	"time"
	domain "unique-minds/Domain"
	infrastructure "unique-minds/Infrastructure"
	utils "unique-minds/Utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	collection *mongo.Collection
	activeUserCollection *mongo.Collection
	studentProfileCollection *mongo.Collection
	educatorProfileCollection *mongo.Collection
	courseCollection *mongo.Collection
	config   *infrastructure.Config
}

func NewUserRepository(collection *mongo.Collection, activeUserColl *mongo.Collection, studentProfile *mongo.Collection, educatorProfile *mongo.Collection, courseColl *mongo.Collection, config *infrastructure.Config) *UserRepository {
	return &UserRepository{
		collection: collection,
		activeUserCollection: activeUserColl,
		config: config,
		studentProfileCollection: studentProfile,
		educatorProfileCollection: educatorProfile,
		courseCollection: courseColl,
	}
}

// FindUserByEmail function finds a user by email
func (ur *UserRepository) FindUserByEmail(email string) (domain.User, error) {
	var user domain.User
	context, _ := context.WithTimeout(context.Background(), time.Duration(ur.config.ContextTimeout) * time.Second)
	filter := bson.M{"email": email}
	err := ur.collection.FindOne(context, filter).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
// FindUserByUserName function finds a user by username
func (ur *UserRepository) FindUserByUserName(username string) (domain.User, error) {
	var user domain.User
	context, _ := context.WithTimeout(context.Background(), time.Duration(ur.config.ContextTimeout) * time.Second)
	filter := bson.M{"user_name": username}
	err := ur.collection.FindOne(context, filter).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// RegisterUser function registers a user
func (ur *UserRepository) RegisterUser(user domain.User) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(ur.config.ContextTimeout) * time.Second)
	defer cancel()
	user.ID = primitive.NewObjectID()
	_, err := ur.collection.InsertOne(context, user)
	if err != nil {
		return err
	}
	if strings.ToLower(user.UserType) == "student"{
		var student = domain.StudentProfile{
			ID: user.ID,
			FullName: "Your Full Name Here",
			Age: "Your",
			Bio: "Your Bio Here",
			GuardianEmail: user.Email,
			GuardianPhone: "Your Guardian Phone Number Here",
			Location: "Your Location Here",
			ProfileImage: "../../Assets/educator.jpg",
			Created_At: user.Created_At,
			UpdateAt: user.Created_At,
			CourseIds: []primitive.ObjectID{},
			EnrolledCourses: []domain.CourseProgress{},
			Schedule: []domain.Schedule{},
			Courses: []domain.Course{},
			Condition: "Your Condition Here",
		}
		_, err = ur.studentProfileCollection.InsertOne(context, student)

	}else{
		var educator = domain.EducatorProfile{
			ID: user.ID,
			FullName: "Your Full Name Here",
			Title: "Your Title Here",
			ProfileImage: "../../Assets/educator.jpg",
			Phone: "Your Phone Number Here",
			Bio: "Your Bio Here",
			Rating: 0,
			Reviews: []domain.Review{},
			Availability: []string{},
			Email: user.Email,
			Created_At: user.Created_At,
			UpdateAt: user.Created_At,
			Address: "Your Address Here",
			Schedules: []domain.Schedule{},
			Students: []domain.Student{},
		}
		_, err = ur.educatorProfileCollection.InsertOne(context, educator)

	}
	return err
}

// UpdateUser function updates a user
func (ur *UserRepository) UpdateUser(id string, user domain.User) error {
	context, _ := context.WithTimeout(context.Background(), time.Duration(ur.config.ContextTimeout) * time.Second)
	user_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": user_id}
	_, err := ur.collection.UpdateOne(context, filter, bson.M{"$set": user})
	if err != nil {
		return err
	}
	return nil
}

// FindUserByID function finds a user by ID
func (ur *UserRepository) FindUserByID(id string)(domain.User, error){
	var user domain.User
	context, _ := context.WithTimeout(context.Background(), time.Duration(ur.config.ContextTimeout) * time.Second)
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}
	err := ur.collection.FindOne(context, filter).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// SaveAsActiveUser function saves a user as active
func (ur *UserRepository) SaveAsActiveUser(user domain.ActiveUser, refreshToken string) error {
	_, err := ur.FindActiveUser(user.ID.Hex(), user.UserAgent)
	if err == nil {
		return errors.New("user already logged in")
	}
	return ur.CreateActiveUser(user)
}

// CreateActiveUser function creates an active user
func (ur *UserRepository) CreateActiveUser(au domain.ActiveUser) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(ur.config.ContextTimeout) * time.Second)
	defer cancel()
	_, err := ur.activeUserCollection.InsertOne(context, au)

	return err
}

// DeleteActiveUser function deletes an active user
func (ur *UserRepository) DeleteActiveUser(ids string, user_agent string) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(ur.config.ContextTimeout) * time.Second)
	defer cancel()
	id, err := primitive.ObjectIDFromHex(ids)
	if err != nil {
		return err
	}
	_, err = ur.activeUserCollection.DeleteOne(context, bson.M{"_id": id, "user_agent": user_agent})
	return err
}

// FindActiveUser function finds an active user
func (ur *UserRepository) FindActiveUser(ids string, user_agent string) (domain.ActiveUser, error) {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(ur.config.ContextTimeout) * time.Second)
	defer cancel()
	id, err := primitive.ObjectIDFromHex(ids)
	if err != nil {
		return domain.ActiveUser{}, err
	}
	var au domain.ActiveUser
	err = ur.activeUserCollection.FindOne(context, bson.M{"_id": id, "user_agent": user_agent}).Decode(&au)
	return au, err
}

// GetEducators function fetches educators
func (ur *UserRepository) GetEducators(pageNo int64, pageSize int64, search string) ([]domain.EducatorProfile, domain.Pagination, error) {
	pagination := utils.PaginationByPage(pageNo, pageSize)

	totalResults, err := ur.educatorProfileCollection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return []domain.EducatorProfile{}, domain.Pagination{}, err
	}
    filter := bson.M{}
    if search != "" {
        filter["name"] = bson.M{"$regex": search, "$options": "i"}
    }
   
	totalPages := int64(math.Ceil(float64(totalResults) / float64(pageSize)))

	cursor, err := ur.educatorProfileCollection.Find(context.TODO(), filter, pagination)
    
	if err != nil {
		return []domain.EducatorProfile{}, domain.Pagination{}, err
	}
	var educators []domain.EducatorProfile
	for cursor.Next(context.TODO()) {
		var educator domain.EducatorProfile
		if err := cursor.Decode(&educator); err != nil {
			return []domain.EducatorProfile{}, domain.Pagination{}, err
		}
		educators = append(educators, educator)
	}
	paginationInfo := domain.Pagination{
		CurrentPage: pageNo,
		PageSize:    pageSize,
		TotalPages:  totalPages,
		TotatRecord: totalResults,
	}

	return educators, paginationInfo, nil
}

// GetEducatorById function fetches an educator by ID
func (ur *UserRepository) GetEducatorsById(id string) (domain.EducatorProfile, error) {
	var educator domain.EducatorProfile

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.EducatorProfile{}, err
	}

	filter := bson.M{"_id": objID}
	err = ur.educatorProfileCollection.FindOne(context.TODO(), filter).Decode(&educator)
	if err != nil { 
		return domain.EducatorProfile{}, err
	}

	return educator, nil
}

// SaveReview function saves a review
func (ur *UserRepository) SaveReview(review domain.Review, id string) error {
	uid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": uid}

	review.EducatorID = uid

	var educatorProfile domain.EducatorProfile
	err := ur.educatorProfileCollection.FindOne(context.TODO(), filter).Decode(&educatorProfile)
	if err != nil {
		return err
	}

	update := bson.M{"$push": bson.M{"reviews": review}}
	_, err = ur.educatorProfileCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	totalRating := float64(educatorProfile.Rating) + (review.Rating)

	newAverageRating := totalRating / float64(len(educatorProfile.Reviews))

	if newAverageRating > 5 {
		newAverageRating = 5
	}

	updateRating := bson.M{
		"$set": bson.M{"rating": newAverageRating},
	}
	_, err = ur.educatorProfileCollection.UpdateOne(context.TODO(), filter, updateRating)
	if err != nil {
		return err
	}

	return nil
}

// GetStudentById function fetches a student by ID
func (ur *UserRepository) GetStudentById(id string) (domain.StudentProfile, error){
	var student domain.StudentProfile
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.StudentProfile{}, err
	}
	filter := bson.M{"_id": objID}
	err = ur.studentProfileCollection.FindOne(context.TODO(), filter).Decode(&student)
	if err != nil {
		return domain.StudentProfile{}, err
	}
	return student, nil
}

// UpdateEducatorProfile function updates an educator profile
func (ur *UserRepository) UpdateEducatorProfile(user_id string, educator domain.EducatorProfile) (domain.EducatorProfile, error){
	context, cancel := context.WithTimeout(context.Background(), time.Duration(ur.config.ContextTimeout) * time.Second)
	defer cancel()
	uid, _ := primitive.ObjectIDFromHex(user_id)
	filter := bson.M{"_id": uid}
	_, err := ur.educatorProfileCollection.UpdateOne(context, filter, bson.M{"$set": educator})
	if err != nil {
		return domain.EducatorProfile{}, err
	}
	return educator, nil
}

// UpdateStudentProfile function updates a student profile
func (ur *UserRepository) UpdateStudentProfile(user_id string, student domain.StudentProfile) (domain.StudentProfile, error){
	context, cancel := context.WithTimeout(context.Background(), time.Duration(ur.config.ContextTimeout) * time.Second)
	defer cancel()
	uid, _ := primitive.ObjectIDFromHex(user_id)
	filter := bson.M{"_id": uid}
	_, err := ur.studentProfileCollection.UpdateOne(context, filter, bson.M{"$set": student})
	if err != nil {
		return domain.StudentProfile{}, err
	}
	return student, nil
}

// GetStudentProfile function fetches a student profile
func (ur *UserRepository) SetAvailability(userID, availability string) error {
	uid, _ := primitive.ObjectIDFromHex(userID)
	filter := bson.M{"_id": uid}
    update := bson.M{
        "$push": bson.M{
            "availability": availability,
        },
    }

    _, err := ur.educatorProfileCollection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        return errors.New("unable to set availability")
    }

    return nil
}

// GetEducatorSchedules function fetches educator schedules
func (ur *UserRepository) FindEducatorSchedules(educatorId string) (interface{}, error) {
    user_id, err := primitive.ObjectIDFromHex(educatorId)
    if err != nil {
        return nil, err
    }
    var educator struct {
        Schedules []domain.Schedule `bson:"schedules"`
    }
    err = ur.educatorProfileCollection.FindOne(context.TODO(), bson.M{"_id": user_id}).Decode(&educator)
    if err != nil {
        return nil, err
    }

	type ScheduleWithStudent struct {
		ID             primitive.ObjectID `json:"id" bson:"_id"`
		Date           string             `json:"date"`
		GoogleMeetLink string             `json:"googleMeetLink"`
		StudentName    string             `json:"studentName"`
	}
    var schedulesWithStudent []ScheduleWithStudent

    for _, schedule := range educator.Schedules {
        var student struct {
            Name  string `bson:"name"`
        }

        err = ur.studentProfileCollection.FindOne(context.TODO(), bson.M{"_id": schedule.StudentID}).Decode(&student)
        if err != nil {
            return nil, err
        }

        scheduleWithStudent := ScheduleWithStudent{
            ID:             schedule.ID,
            Date:           schedule.Date, //schedule.Date.Format("2006-01-02 15:04"),
            GoogleMeetLink: schedule.GoogleMeetLink,
            StudentName:    student.Name,
        }
        schedulesWithStudent = append(schedulesWithStudent, scheduleWithStudent)
    }
    return schedulesWithStudent, nil
}

// GetStudentSchedules function fetches student schedules
func (ur *UserRepository) FindStudentSchedules(studentId string) (interface{}, error) {
    user_id, err := primitive.ObjectIDFromHex(studentId)
    if err != nil {
        return nil, err
    }
    var students struct {
        Schedules []domain.Schedule `bson:"schedules"`
    }
    err = ur.studentProfileCollection.FindOne(context.TODO(), bson.M{"_id": user_id}).Decode(&students)
    if err != nil {
        return nil, err
    }

	type ScheduleWithStudent struct {
		ID              primitive.ObjectID `json:"id" bson:"_id"`
		Date            string             `json:"date"`
		GoogleMeetLink  string             `json:"googleMeetLink"`
		EducatorName    string            `json:"educatorName"`
		EducatorID		string            `json:"educatorID"`
		StudentID       string            `json:"studentID"`
	}
    var schedulesWithStudent []ScheduleWithStudent

    for _, schedule := range students.Schedules {
        var student struct {
            Name  string `bson:"name"`
        }

        err = ur.educatorProfileCollection.FindOne(context.TODO(), bson.M{"_id": schedule.EducatorId}).Decode(&student)
        if err != nil {
            return nil, err
        }

        scheduleWithStudent := ScheduleWithStudent{
            ID:             schedule.ID,
            Date:           schedule.Date,  
            GoogleMeetLink: schedule.GoogleMeetLink,
            EducatorName:   student.Name,
			EducatorID: schedule.EducatorId.Hex(),
			StudentID: studentId,
        }
        schedulesWithStudent = append(schedulesWithStudent, scheduleWithStudent)
    }
    return schedulesWithStudent, nil
}

// DeleteSchedule function deletes a schedule
func (ur *UserRepository) DeleteSchedule(scheduleId string, userId string) error {
    scheduleObjID, _ := primitive.ObjectIDFromHex(scheduleId)
    userObjID, _ := primitive.ObjectIDFromHex(userId)
	var eduId primitive.ObjectID

	var student domain.StudentProfile

	err := ur.studentProfileCollection.FindOne(context.TODO(), bson.M{"_id": userObjID, "schedules._id": scheduleObjID}).Decode(&student)
	if err != nil {
		return err
	}

	for _, schedule := range student.Schedule {
		if schedule.ID == scheduleObjID {
			eduId = schedule.EducatorId
			break
		}
	}
    educatorFilter := bson.M{
        "_id":           eduId,
        "schedules._id": scheduleObjID,
    }
    update := bson.M{
        "$pull": bson.M{
            "schedules": bson.M{"_id": scheduleObjID},
        },
    }

    _, err = ur.educatorProfileCollection.UpdateOne(context.TODO(), educatorFilter, update)
    if err != nil {
        return err
    }

    studentFilter := bson.M{
        "_id":           userObjID,
        "schedules._id": scheduleObjID,
    }

    _, err = ur.studentProfileCollection.UpdateOne(context.TODO(), studentFilter, update)
    if err != nil {
        return err
    }

    return nil
}

// DeleteEducatorSchedule function deletes an educator schedule
func (ur *UserRepository) DeleteEducatorSchedule(scheduleId string, userId string) error {
    scheduleObjID, _ := primitive.ObjectIDFromHex(scheduleId)
    userObjID, _ := primitive.ObjectIDFromHex(userId)
	var stuId primitive.ObjectID

	var educator domain.EducatorProfile

	err := ur.educatorProfileCollection.FindOne(context.TODO(), bson.M{"_id": userObjID, "schedules._id": scheduleObjID}).Decode(&educator)
	if err != nil {
		return err
	}

	for _, schedule := range educator.Schedules {
		if schedule.ID == scheduleObjID {
			stuId = schedule.StudentID
			break
		}
	}
    educatorFilter := bson.M{
        "_id":           userObjID,
        "schedules._id": scheduleObjID,
    }
    update := bson.M{
        "$pull": bson.M{
            "schedules": bson.M{"_id": scheduleObjID},
        },
    }

    _, err = ur.educatorProfileCollection.UpdateOne(context.TODO(), educatorFilter, update)
    if err != nil {
        return err
    }

    studentFilter := bson.M{
        "_id":           stuId,
        "schedules._id": scheduleObjID,
    }

    _, err = ur.studentProfileCollection.UpdateOne(context.TODO(), studentFilter, update)
    if err != nil {
        return err
    }

    return nil
}

// GetStudentsFromEducatorProfile function fetches students from an educator profile
func (ur *UserRepository) GetStudentsFromEducatorProfile(educatorID string) (map[string][]domain.StudentProfile, error) {
    educatorObjID, err := primitive.ObjectIDFromHex(educatorID)
    if err != nil {
        return nil, err
    }

	var courses []domain.Course

	cursor, err := ur.courseCollection.Find(context.TODO(), bson.M{"user_id": educatorObjID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var c domain.Course
		if err := cursor.Decode(&c); err != nil {
			return nil, err
		}
		courses = append(courses, c)
	}

	result := make(map[string][]domain.StudentProfile)

	for _, course := range courses {
		for _, id := range course.Students {
			var student domain.StudentProfile
			err := ur.studentProfileCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&student)
			if err != nil {
				continue
			}
			result[course.Name] = append(result[course.Name], student)
		}
	}

		return result, nil
}

// FindById function finds a user by ID
func (ur *UserRepository) FindById(userID string) (*domain.UserData, error) {
    var student domain.StudentProfile
    var educator domain.EducatorProfile
	uid, _ := primitive.ObjectIDFromHex(userID)
	
    err := ur.studentProfileCollection.FindOne(context.TODO(), bson.M{"_id": uid}).Decode(&student)
    if err == nil {
        return &domain.UserData{
            ProfileImage: student.ProfileImage,
            Role:         "student",
        }, nil
    }

    err = ur.educatorProfileCollection.FindOne(context.TODO(), bson.M{"_id": uid}).Decode(&educator)
    if err == nil {
        return &domain.UserData{
            ProfileImage: educator.ProfileImage,
            Role:         "educator",
        }, nil
    }
    return nil, err
}

// GetTopEducators function fetches top educators
func (ur *UserRepository) GetTopEducators() ([]domain.EducatorProfile, error) {
    var topEducators []domain.EducatorProfile

    findOptions := options.Find()
    findOptions.SetSort(bson.D{{Key: "rating", Value: -1}})
    findOptions.SetLimit(3)

    cursor, err := ur.educatorProfileCollection.Find(context.TODO(), bson.D{}, findOptions)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())

    if err = cursor.All(context.TODO(), &topEducators); err != nil {
        return nil, err
    }

    return topEducators, nil
}

// FetchUserEnrolledCourses function fetches user enrolled courses
func (ur *UserRepository) FetchUserEnrolledCourses(userID string) ([]domain.CourseProgress, error) {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": userObjID}
	var user struct {
		EnrolledCourses []domain.CourseProgress `bson:"courses" json:"courses"`
	}
	err = ur.studentProfileCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user.EnrolledCourses, nil
}

// FetchCourseNameByID function fetches course name by ID
func (ur *UserRepository) FetchCourseNameByID(courseID primitive.ObjectID) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": courseID}
	var course struct {
		Name string `bson:"name" json:"name"`
	}
	err := ur.courseCollection.FindOne(ctx, filter).Decode(&course)
	if err != nil {
		return "", err
	}
	return course.Name, nil
}

// UpdateSchedule function updates a schedule
func (ur *UserRepository) UpdateSchedules(user_id string, educatorID string, availability string) error {
	uid, _ := primitive.ObjectIDFromHex(user_id)
	educator_id, _ := primitive.ObjectIDFromHex(educatorID)

	schedule := domain.Schedule{
		ID: primitive.NewObjectID(),
		StudentID: uid,
		EducatorId: educator_id,
		Date: availability,
		GoogleMeetLink: " https://meet.google.com/jjg-ifsj-bkm",

	}
	educatorUpdate := bson.M{
		"$push": bson.M{"schedules": schedule},
		"$pull": bson.M{"availability": availability},
	}

	_, err := ur.educatorProfileCollection.UpdateOne(context.TODO(), bson.M{"_id": educator_id}, educatorUpdate)
	if err != nil {
		return err
	}
	educatorUpdate = bson.M{
		"$push": bson.M{"students": schedule},
	}

	studentUpdate := bson.M{
		"$push": bson.M{"schedules": schedule},
	}
	_, err = ur.studentProfileCollection.UpdateOne(context.TODO(), bson.M{"_id": uid}, studentUpdate)
	if err != nil{
		return err
	}
	return nil
}

// UpdateProfileImage function updates a profile image
func (ur *UserRepository) UpdateProfileImage(user_id string, user_type, profileImage string) error {
	uid, _ := primitive.ObjectIDFromHex(user_id)
	filter := bson.M{"_id": uid}
	update := bson.M{"$set": bson.M{"profileImage": profileImage}}
	if user_type == "student"{
		_, err := ur.studentProfileCollection.UpdateOne(context.TODO(), filter, update)
		return err
	}else{
		_, err := ur.educatorProfileCollection.UpdateOne(context.TODO(), filter, update)
		return err
	}
}