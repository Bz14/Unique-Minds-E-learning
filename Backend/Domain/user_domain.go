package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User struct
type User struct {
	ID                   primitive.ObjectID   `bson:"_id,omitempity" json:"id" `
	UserName            string                `bson:"username"  json:"username"`
	Email 				 string 			  `bson:"email" validate:"required,email" json:"email"`
    Password             string               `bson:"password" json:"password" validate:"required"`
	IsVerified			 bool 				  `bson:"is_verified" json:"is_verified"`
	Created_At		     time.Time			  `bson:"created_at" json:"created_at"`
    ResetPasswordToken   string               `bson:"reset_password_token" json:"reset_password_token"`
    ResetPasswordExpires time.Time            `bson:"reset_password_expires" json:"reset_password_expires"`
	VerificationToken    string               `bson:"verification_token" json:"verification_token"`
	VerificationExpires  time.Time            `bson:"verification_expires" json:"verification_expires"`
	Role 			   	 string               `bson:"role" json:"role"`
	UserType        	 string 			  `bson:"user_type" json:"user_type" validate:"required"`
}

// UserData struct
type UserData struct {
    ProfileImage string
    Role         string
}


// StudentProfile struct
type StudentProfile struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FullName      string             `bson:"name" json:"name"`
	Age           string                `json:"age" bson:"age"`
	Bio           string             `json:"bio" bson:"bio"`
	GuardianEmail string             `json:"guardianEmail" bson:"guardianEmail"`
	GuardianPhone string             `json:"guardianPhone" bson:"guardianPhone"`
	Location      string             `json:"location" bson:"location"`
	ProfileImage  string             `json:"profileImage" bson:"profileImage"`
	UpdateAt      time.Time			 `json:"updateAt" bson:"updateAt"`
	Created_At	  time.Time			  `bson:"created_at" json:"created_at"`
	CourseIds     []primitive.ObjectID `bson:"course_id" json:"course_id"`
	EnrolledCourses []CourseProgress 		`bson:"courses" json:"courses"`
	Schedule 	   []Schedule         `bson:"schedules" json:"schedules"`
	Courses     []Course           `bson:"course_s" json:"course_s"`
	Condition	 string             `bson:"condition" json:"condition"`

}

// CourseProgress struct
type CourseProgress struct{
	CourseID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Progress	 int64      `json:"progress" bson:"progress"`
	CompletedParts []primitive.ObjectID  `json:"completed_parts" bson:"completed_parts"`
	IsCompleted    bool  `json:"is_completed" bson:"is_completed"`

}

// CourseDetailResponse struct
type CourseDetailResponse struct{
	Course  Course `json:"course"`
	Progress CourseProgress `json:"progress"`
}

// ProgressRequest struct
type ProgressRequest struct {
    CompletedParts []string `json:"completed_parts"`
}

// EducatorProfile struct
type EducatorProfile struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FullName      string       		 `bson:"name" json:"name"`
	Title 	   	  string             `bson:"title" json:"title"`
	ProfileImage  string             `json:"profileImage" bson:"profileImage"`
	Phone 		  string 			 `bson:"phone" json:"phone"`
	Bio 		  string             `json:"bio" bson:"bio"`
	Rating        float32		     `json:"rating" bson:"rating"`
	Reviews       []Review           `bson:"reviews" json:"reviews"`
	Availability  []string             `bson:"availability" json:"availability"`
	Email 		  string 			 `bson:"email" validate:"required,email" json:"email"`
	Created_At	  time.Time			  `bson:"created_at" json:"created_at"`
	UpdateAt      time.Time			 `json:"updateAt" bson:"updateAt"`
	Address       string            `bson:"address" json:"address"`
	Schedules     []Schedule         `bson:"schedules" json:"schedules"`
	Students 	  []Student          `bson:"students" json:"students"`
	
}

// Student struct
type Student struct{
	Student_id primitive.ObjectID `json:"student_id" bson:"student_id"`
	Course_id  primitive.ObjectID `json:"course_id" bson:"course_id"`
}

// Schedule struct
type Schedule struct {
	ID 		  primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	StudentID primitive.ObjectID `json:"student_id" bson:"student_id"`
	EducatorId primitive.ObjectID `json:"educator_id" bson:"educator_id"`
	Date 	  string          `json:"date" bson:"date"`
	GoogleMeetLink string       `json:"googleMeetLink" bson:"googleMeetLink"`
}

// Address struct
type Address struct {
	Street     string `bson:"street" json:"street"`
	City       string `bson:"city" json:"city"`
	PostalCode string `bson:"postalCode" json:"postalCode"`
}

// Review struct
type Review struct {
    Name        string  `bson:"name" json:"name"`
    Text        string  `bson:"text" json:"text"`
    Rating      float64 `bson:"rating" json:"rating"`
    EducatorID  primitive.ObjectID    `bson:"educator_id" json:"educator_id"`
}

// Profile struct
type Profile struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
}

// CourseWithStudents struct
type CourseWithStudents struct {
	CourseName string           `json:"courseName"`
	Students   []StudentProfile `json:"students"`
}

// UserUseCaseInterface Interface
type UserUseCaseInterface interface {
	RegisterUser(user User) error
	VerifyEmail(email string, token string) error
	Login(user User, user_agent string)(LoginResponse, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
	RefreshToken(request RefreshTokenRequest, user_id string) (RefreshTokenResponse, error)
	ResetPassword(email string, user_id string)error
	ResetPasswordVerify(email string, token string, user_id string, password string) error
	Logout(user_id string, user_agent string) error	
	GetEducators(pageNo string, pageSize string, search string) ([]EducatorProfile, Pagination, error)
	GetEducatorById(id string) (EducatorProfile, error)
	SaveReview(review Review, id string) error
	UpdateStudentProfile(user_id string, updatedProfile StudentProfile) (StudentProfile, error)
	GetStudentProfile(user_id string) (StudentProfile, error)
	UpdateEducatorProfile(user_id string, updatedProfile EducatorProfile) (EducatorProfile, error)
	GetEducatorProfile(user_id string) (EducatorProfile, error)
	SetAvailability(userID string, availability string) error
	GetEducatorSchedules(educatorId string) (interface{}, error)
	CancelEducatorSchedule(scheduleId string, user_id string) error
	FetchStudentsByCourses(educatorID string) (map[string][]StudentProfile, error)
	GetUserProfile(userID string) (*UserData, error)
	GetTopEducatorsUseCase() ([]EducatorProfile, error)
	GetEnrolledCoursesProgress(userID string) ([]map[string]interface{}, error)
	ScheduleSession(user_id string, educatorID string , availability string) error
	GetStudentSchedules(student_id string) (interface{}, error)
	SaveProfileImage(user_id string, user_type string, profileImage string) error
	CancelSchedule(scheduleId string, user_id string) error
}

// UserRepositoryInterface interface
type UserRepositoryInterface interface {
	RegisterUser(user User) error
	FindUserByEmail(email string) (User, error)
	FindUserByUserName(username string) (User, error)
	FindUserByID(id string)(User, error)
	SaveAsActiveUser(user ActiveUser, refreshToken string) error
	CreateActiveUser(au ActiveUser) error
	DeleteActiveUser(ids string, user_agent string) error
	FindActiveUser(ids string, user_agent string) (ActiveUser, error)
	UpdateUser(id string, user User) error
	GetEducators(pageNo int64, pageSize int64, search string) ([]EducatorProfile, Pagination, error)
	GetEducatorsById(id string) (EducatorProfile, error)
	SaveReview(review Review, id string) error
	GetStudentById(id string) (StudentProfile, error)
	UpdateEducatorProfile(user_id string, educator EducatorProfile)(EducatorProfile, error)
	UpdateStudentProfile(user_id string, student StudentProfile) (StudentProfile, error)
	SetAvailability(userID, availability string) error 
	FindEducatorSchedules(educatorId string) (interface{}, error)
	DeleteSchedule(scheduleId string, userId string) error
	GetStudentsFromEducatorProfile(educatorID string) (map[string][]StudentProfile, error)
	FindById(userID string) (*UserData, error)
	GetTopEducators() ([]EducatorProfile, error)
	FetchCourseNameByID(courseID primitive.ObjectID) (string, error)
	FetchUserEnrolledCourses(userID string) ([]CourseProgress, error)
	UpdateSchedules(user_id string, educatorID string, availability string) error
	FindStudentSchedules(studentId string) (interface{}, error)
	UpdateProfileImage(user_id string, user_type, profileImage string) error 
	DeleteEducatorSchedule(scheduleId string, userId string) error
}