package usecases

import (
	"errors"
	"strconv"
	"time"
	domain "unique-minds/Domain"
	infrastructure "unique-minds/Infrastructure"

	"github.com/golang-jwt/jwt/v4"
)

type UserUseCase struct {
	UserRepo domain.UserRepositoryInterface
	PassService infrastructure.PasswordService
	Config *infrastructure.Config
}


func NewUserUseCase(userRepo domain.UserRepositoryInterface, passwordService infrastructure.PasswordService, config *infrastructure.Config) *UserUseCase {
	return &UserUseCase{
		UserRepo: userRepo,
		PassService: passwordService,
		Config: config,
	}
}

// RegisterUser function registers a new user
func (uc *UserUseCase) RegisterUser(user domain.User) error {
	if user.Email == "" || user.UserName== "" || user.Password == "" {
		return errors.New("all fields are required")
	}
	if infrastructure.ValidateEmail(user.Email) != nil {
		return errors.New("invalid email format")
	}
	if infrastructure.ValidatePassword(user.Password) != nil {
		return errors.New("invalid password format")
	}

	existingUser, err := uc.UserRepo.FindUserByEmail(user.Email)

	if err == nil && existingUser.Email != "" && !existingUser.IsVerified{
		return err
	}else if err == nil{
		return err
	}

	hashedPassword, _ := uc.PassService.HashPassword(user.Password)
	user.Password = hashedPassword

	token, err  := infrastructure.GenerateVerificationToken()
	if err != nil{
		return err
	}
	err = infrastructure.SendVerificationEmail(user.Email, token)
	if err != nil {
		return err
	}

	user.IsVerified = false
	user.Created_At = time.Now()
	user.VerificationToken = token
	user.VerificationExpires = time.Now().Add(time.Hour * 24)
	user.Role = "user"
	err = uc.UserRepo.RegisterUser(user)
	if err != nil {
		return err
	}
	return nil
}

// VerifyEmail function verifies a user's email
func (uc *UserUseCase) VerifyEmail(email string, token string)error{
	user, err := uc.UserRepo.FindUserByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	if user.VerificationToken != token {
		return errors.New("invalid verification token")
	}
	if user.VerificationExpires.Before(time.Now()) {
		return errors.New("verification token expired. Please request a new one")
	}

	user.IsVerified = true
	user.VerificationToken = ""
	user.VerificationExpires = time.Time{}
	err = uc.UserRepo.UpdateUser(user.ID.Hex(), user)
	if err != nil {
		return errors.New("error verifying user")
	}
	return nil
}

// Login function logs in a user
func (uc *UserUseCase) Login(user domain.User, user_agent string)(domain.LoginResponse, error){
	var newUser domain.User
	var err error
	if user.Email == "" || user.Password == "" {
		return domain.LoginResponse{}, errors.New("all fields are required")
	}
	if user.Email != ""{
		newUser, err = uc.UserRepo.FindUserByEmail(user.Email)	
	}else if user.UserName != ""{
		newUser, err = uc.UserRepo.FindUserByUserName(user.UserName)
	}
	
	if err != nil{
		return domain.LoginResponse{}, errors.New("user not found")
	}
	if !newUser.IsVerified{
		return domain.LoginResponse{}, errors.New("user not verified. Please Verify Your Account")
	}
	if !uc.PassService.ComparePassword(user.Password, newUser.Password){
		return domain.LoginResponse{}, errors.New("invalid credentials: Password does not match")
	}
	accessToken, err := uc.CreateAccessToken(&newUser, uc.Config.AccessTokenSecret, uc.Config.AccessTokenExpiryHour)
	if err != nil {
		return domain.LoginResponse{}, errors.New("error creating access token")
	}
	refreshToken, err  := uc.CreateRefreshToken(&newUser, uc.Config.RefreshTokenSecret, uc.Config.RefreshTokenExpiryHour)
	if err != nil {
		return domain.LoginResponse{}, errors.New("error creating refresh token")
	}

	activeUser:= domain.ActiveUser{
		ID: newUser.ID,
		RefreshToken: refreshToken,
		UserAgent: user_agent,
	}
	err = uc.UserRepo.SaveAsActiveUser(activeUser, refreshToken)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	return domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// CreateAccessToken function creates an access token
func (uc *UserUseCase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	exp := time.Now().Add(5 * time.Hour * time.Duration(expiry))
	claims := &domain.JwtCustomClaims{
		ID: user.ID.Hex(),
		UserType: user.UserType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	return infrastructure.CreateToken(claims, secret)
}

// CreateRefreshToken function creates a refresh token
func (uc *UserUseCase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry))

	claims := &domain.JwtCustomClaims{
		ID: user.ID.Hex(),
		UserType: user.UserType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	return infrastructure.CreateToken(claims, secret)
}

// RefreshToken function refreshes a user's token
func (uc *UserUseCase) RefreshToken(request domain.RefreshTokenRequest, user_id string) (domain.RefreshTokenResponse, error) {
	id, err := infrastructure.ExtractID(request.RefreshToken, uc.Config.RefreshTokenSecret)
	if err != nil {
		return domain.RefreshTokenResponse{}, errors.New("user not found")
	}
	valid, err := infrastructure.IsAuthorized(request.RefreshToken, uc.Config.RefreshTokenSecret)
	if err != nil {
		return domain.RefreshTokenResponse{}, errors.New("user not found")
	}

	if id != user_id && !valid {
		return domain.RefreshTokenResponse{}, errors.New("session expired")
	}

	user, err := uc.UserRepo.FindUserByID(user_id)
	if err != nil {
		return domain.RefreshTokenResponse{}, errors.New("user not found")
	}
	accessToken, err := uc.CreateAccessToken(&user, uc.Config.AccessTokenSecret, uc.Config.AccessTokenExpiryHour)
	if err != nil {
		return domain.RefreshTokenResponse{}, errors.New(err.Error())
	}
	return domain.RefreshTokenResponse{
		AccessToken:  accessToken,
	}, nil
}

// ResetPassword function resets a user's password
func (uc *UserUseCase) ResetPassword(email string, user_id string) error{
	user, err := uc.UserRepo.FindUserByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}
	if user.ID.Hex() != user_id {
		return errors.New("unauthorized: User not found")
	}

	token, err := infrastructure.GenerateVerificationToken()
	if err != nil {
		return errors.New("error generating token")
	}
	err = infrastructure.SendResetPasswordVerificationEmail(email, token)
	if err != nil {
		return errors.New("error sending email")
	}

	user.ResetPasswordToken = token
	user.ResetPasswordExpires = time.Now().Add(time.Hour * 24)
	err = uc.UserRepo.UpdateUser(user_id, user)
	if err != nil {
		return errors.New("error updating user")
	}
	return nil
}

// ResetPasswordVerify function verifies a user's password reset
func (uc *UserUseCase) ResetPasswordVerify(email string, token string, user_id string, password string) error{
	user, err := uc.UserRepo.FindUserByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}
	if user.ID.Hex() != user_id {
		return errors.New("unauthorized: User not found")
	}

	if err := infrastructure.ValidatePassword(password); err != nil {
		return errors.New("invalid password format")
	}

	if user.ResetPasswordToken != token {
		return errors.New("invalid token")
	}
	if user.ResetPasswordExpires.Before(time.Now()) {
		return errors.New("token expired")
	}
	newPassword, err := uc.PassService.HashPassword(password)
	if err != nil{
		return errors.New("error hashing password")
	}
	user.Password = newPassword
	user.ResetPasswordToken = ""
	user.ResetPasswordExpires = time.Time{}

	err = uc.UserRepo.UpdateUser(user_id, user)
	if err != nil{
		return errors.New("can not reset password")
	}
	return nil
}

// Logout function logs out a user
func (uc *UserUseCase) Logout(user_id string, user_agent string) error {
	_, err := uc.UserRepo.FindActiveUser(user_id, user_agent)
	if err != nil {
		return errors.New("user not found, login before logout")
	}

	return uc.UserRepo.DeleteActiveUser(user_id, user_agent)
}

// GetEducators function gets all educators
func (uc *UserUseCase) GetEducators(pageNo string, pageSize string, search string) ([]domain.EducatorProfile, domain.Pagination, error) {
	PageNo, err := strconv.ParseInt(pageNo, 10, 64)
	if err != nil {
		return []domain.EducatorProfile{}, domain.Pagination{}, err
	}
	PageSize, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		return []domain.EducatorProfile{}, domain.Pagination{}, err
	}
	if PageNo <= 0 || PageSize <= 0 {
		return []domain.EducatorProfile{}, domain.Pagination{}, errors.New("invalid page number or page size")
	}

	educators, pagination, err := uc.UserRepo.GetEducators(PageNo, PageSize, search)
	if err != nil {
		return nil, domain.Pagination{}, err
	} else {
		return educators, pagination, nil
	}
}

// GetEducatorById function gets an educator by ID
func (uc *UserUseCase) GetEducatorById(id string) (domain.EducatorProfile, error) {
	result, err := uc.UserRepo.GetEducatorsById(id)
	if err != nil {
		return domain.EducatorProfile{}, err
	}
	return result, nil
}

// SaveReview function saves a review
func (uc *UserUseCase) SaveReview(review domain.Review, id string) error{
	return uc.UserRepo.SaveReview(review, id)
}

// GetEducatorReviews function gets all reviews for an educator
func (uc *UserUseCase) GetEducatorProfile(user_id string) (domain.EducatorProfile, error){
	educator, err := uc.UserRepo.GetEducatorsById(user_id)
	if err != nil {
		return domain.EducatorProfile{}, errors.New("educator not found")
	}
	return educator, nil
}

// GetStudentProfile function gets a student's profile
func (uc *UserUseCase) GetStudentProfile(user_id string) (domain.StudentProfile, error){
	student, err := uc.UserRepo.GetStudentById(user_id)
	if err != nil {
		return domain.StudentProfile{}, errors.New("student not found")
	}
	return student, nil
}

// UpdateEducatorProfile function updates an educator's profile
func (uc *UserUseCase) UpdateEducatorProfile(user_id string, updatedProfile domain.EducatorProfile) (domain.EducatorProfile, error){
	educator, err := uc.UserRepo.GetEducatorsById(user_id)
	if err != nil {
		return domain.EducatorProfile{}, errors.New("educator not found")
	}
	educator.ID = updatedProfile.ID
	educator.FullName = updatedProfile.FullName
	educator.Title = updatedProfile.Title
	educator.Bio = updatedProfile.Bio
	educator.Email = updatedProfile.Email
	educator.Rating = updatedProfile.Rating
	educator.Phone = updatedProfile.Phone
	educator.Address = updatedProfile.Address
	educator.ProfileImage = updatedProfile.ProfileImage
	educator.Address = updatedProfile.Address
	educator.Availability = updatedProfile.Availability
	educator.UpdateAt = time.Now()
	
	result, err := uc.UserRepo.UpdateEducatorProfile(user_id, educator)
	return result, err
}

// UpdateStudentProfile function updates a student's profile
func (uc *UserUseCase) UpdateStudentProfile(user_id string, updatedProfile domain.StudentProfile)(domain.StudentProfile, error){
	student, err := uc.UserRepo.GetStudentById(user_id)
	if err != nil {
		return domain.StudentProfile{}, errors.New("student not found")
	}
	student.ID = updatedProfile.ID
	student.FullName = updatedProfile.FullName
	student.Age = updatedProfile.Age
	student.Bio = updatedProfile.Bio
	student.GuardianEmail = updatedProfile.GuardianEmail
	student.GuardianPhone = updatedProfile.GuardianPhone
	student.Location = updatedProfile.Location
	student.ProfileImage = updatedProfile.ProfileImage
	student.UpdateAt = time.Now()
	result, err := uc.UserRepo.UpdateStudentProfile(user_id, student)
	return result, err
}

// SetAvailability function sets an educator's availability
func (uc *UserUseCase) SetAvailability(userID string, availability string) error {
    err := uc.UserRepo.SetAvailability(userID, availability)
    if err != nil {
		return errors.New("error setting availability")
    }

    return nil
}

// GetEducatorSchedules function gets an educator's schedules
func (uc *UserUseCase) GetEducatorSchedules(educatorId string) (interface{}, error) {
	return uc.UserRepo.FindEducatorSchedules(educatorId)
}

// GetStudentSchedules function gets a student's schedules
func (uc *UserUseCase) GetStudentSchedules(student_id string) (interface{}, error) {
	return uc.UserRepo.FindStudentSchedules(student_id)
}

// CancelEducatorSchedule function cancels an educator's schedule
func (uc *UserUseCase) CancelEducatorSchedule(scheduleId string, user_id string) error {
    return uc.UserRepo.DeleteSchedule(scheduleId, user_id)
}

// CancelSchedule function cancels a schedule
func (uc *UserUseCase) CancelSchedule(scheduleId string, user_id string) error {
    return uc.UserRepo.DeleteEducatorSchedule(scheduleId, user_id)
}

// FetchStudentsByCourses function fetches students by courses
func (uc *UserUseCase) FetchStudentsByCourses(educatorID string) (map[string][]domain.StudentProfile, error) {
    return uc.UserRepo.GetStudentsFromEducatorProfile(educatorID)
}

// GetUserProfile function gets a user's profile
func (uc *UserUseCase) GetUserProfile(userID string) (*domain.UserData, error) {
    return uc.UserRepo.FindById(userID)
}

// GetTopEducatorsUseCase function gets top educators
func (uc *UserUseCase) GetTopEducatorsUseCase() ([]domain.EducatorProfile, error) {
	return uc.UserRepo.GetTopEducators()
}

// GetEnrolledCoursesProgress function gets a user's enrolled courses progress
func (uc *UserUseCase) GetEnrolledCoursesProgress(userID string) ([]map[string]interface{}, error) {
	enrolledCourses, err := uc.UserRepo.FetchUserEnrolledCourses(userID)
	if err != nil {
		return nil, err
	}
	var courseProgressList []map[string]interface{}

	for _, enrolledCourse := range enrolledCourses {
		courseName, err := uc.UserRepo.FetchCourseNameByID(enrolledCourse.CourseID)
		if err != nil {
			return nil, err
		}

		courseProgressList = append(courseProgressList, map[string]interface{}{
			"id":       enrolledCourse.CourseID,
			"name":     courseName,
			"progress": enrolledCourse.Progress,
		})
	}

	return courseProgressList, nil
}

// ScheduleSession function schedules a session
func (uc *UserUseCase) ScheduleSession(user_id string, educatorID string , availability string) error {
	return uc.UserRepo.UpdateSchedules(user_id, educatorID, availability)
}

// SaveProfileImage function saves a user's profile image
func (uc *UserUseCase) SaveProfileImage(user_id string, user_type string, profileImage string) error {
	return uc.UserRepo.UpdateProfileImage(user_id, user_type, profileImage)
}