package Controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	domain "unique-minds/Domain"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
)


type UserControllers struct{
	userUserCase domain.UserUseCaseInterface
}

func NewUserControllers(userUseCase domain.UserUseCaseInterface) *UserControllers {
	return &UserControllers{
		userUserCase: userUseCase,
	}
}
func (uc *UserControllers) RegisterUser(c *gin.Context){
	var signUp domain.SignUpRequest
	var user domain.User

	err := c.BindJSON(&signUp)
	
	if err != nil {
		c.JSON(400, domain.ErrorResponse{
			Message: err.Error(),
			Status:  400,
		})
		return
	}
	validate := validator.New()

	if err := validate.Struct(signUp); err != nil {
		c.JSON(400, domain.ErrorResponse{
			Message: err.Error(),
			Status:  400,
		})
		return
	}
	copier.Copy(&user, &signUp)

	err = uc.userUserCase.RegisterUser(user)
	if err != nil{
		c.JSON(500, domain.ErrorResponse{
			Message: err.Error(),
			Status:  500,
		})
		return
	}

	c.JSON(200, domain.SuccessResponse{
		Message: "User registered successfully. Please verify your email",
		Status:  200,
	})

}

func (uc *UserControllers) VerifyEmail(c *gin.Context){
	token := c.Query("token")
	email := c.Query("email")

	if token == "" || email == "" {
		c.JSON(400, domain.ErrorResponse{
			Message: "Both token and email required",
			Status:  400,
		})
		return
	}

	err := uc.userUserCase.VerifyEmail(email, token)
	if err != nil{
		c.JSON(400, domain.ErrorResponse{
			Message: err.Error(),
			Status:  400,
		})
		return
	}
	c.Redirect(http.StatusFound, "http://localhost:3000/verification")
}

func (uc *UserControllers) Login(c *gin.Context){
	var loginRequest domain.LoginRequest
	var user domain.User

	err := c.BindJSON(&loginRequest)
	if err != nil {
		c.JSON(400, domain.ErrorResponse{
			Message: "Invalid request",
			Status:  400,
		})
		return
	}
	validate := validator.New()

	if err := validate.Struct(loginRequest); err != nil {
		c.JSON(400, domain.ErrorResponse{
			Message: "Invalid request",
			Status:  400,
		})
		return
	}
	copier.Copy(&user, &loginRequest)
	response, err := uc.userUserCase.Login(user, c.Request.UserAgent())
	if err != nil{
		c.JSON(500, domain.ErrorResponse{
			Message: err.Error(),
			Status:  500,
		})
		return
	}
	
	c.JSON(200, domain.SuccessResponse{
		Message: "User Logged in successfully",
		Data: response,
		Status:  200,
	})
}


func (uc *UserControllers) RefreshToken(c *gin.Context){
	var request domain.RefreshTokenRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(400, domain.ErrorResponse{
			Message: "Invalid request",
			Status:  400,
		})
		return
	}
	user_id := c.GetString("user_id")
	if user_id == "" {
		c.JSON(500, domain.ErrorResponse{
			Message: "Unauthorized: Authorization header required",
			Status:  500,
		})
	}
	response , err := uc.userUserCase.RefreshToken(request, user_id)
	if err != nil{
		c.JSON(400, domain.ErrorResponse{
			Message: err.Error(),
			Status:  400,
		})
		return
	}
	c.JSON(
		200, domain.SuccessResponse{
			Message: "Token refreshed successfully",
			Data: response,
			Status:  200,
		},
	)

}

func (uc *UserControllers) ResetPassword(c *gin.Context){
	var request domain.ResetPasswordRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(400, domain.ErrorResponse{
			Message: "Invalid request",
			Status:  400,
		})
		return
	}
	user_id := c.GetString("user_id")
	if user_id == "" {
		c.JSON(500, domain.ErrorResponse{
			Message: "Unauthorized: Authorization header required",
			Status:  500,
		})
	}
	err = uc.userUserCase.ResetPassword(request.Email, user_id)
	if err != nil{
		c.JSON(400, domain.ErrorResponse{
			Message: err.Error(),
			Status:  400,
		})
		return
	}
	c.JSON(200, domain.SuccessResponse{
		Message: "Password reset link sent successfully",
		Status:  200,
	})
}

func (uc *UserControllers) ResetPasswordVerify(c *gin.Context){
	var newPassword domain.ResetPassword
	token := c.Query("token")
	email := c.Query("email")

	if token == "" || email == "" {
		c.JSON(400, domain.ErrorResponse{
			Message: "Both token and email required",
			Status:  400,
		})
		return
	}
	err := c.BindJSON(&newPassword)
	if err != nil {
		c.JSON(400, domain.ErrorResponse{
			Message: "Invalid request",
			Status:  400,
		})
		return 
	}

	user_id := c.GetString("user_id")
	if user_id == "" {
		c.JSON(500, domain.ErrorResponse{
			Message: "Unauthorized: Authorization header required",
			Status:  500,
		})
	}
	err = uc.userUserCase.ResetPasswordVerify(email, token, user_id, newPassword.NewPassword)
	if err != nil{
		c.JSON(400, domain.ErrorResponse{
			Message: err.Error(),
			Status:  400,
		})
		return
	}
	c.JSON(200, domain.SuccessResponse{
		Message: "Password updated successfully",
		Status:  200,})
}


func (uc *UserControllers) Logout(c *gin.Context) {
	user_id := c.GetString("user_id")
	user_agent := c.Request.UserAgent()
	err := uc.userUserCase.Logout(user_id, user_agent)
	if err != nil {
		c.JSON(500, domain.ErrorResponse{
			Message: "Internal server error",
			Status:  500,
		})
		return
	}
	c.JSON(200, domain.SuccessResponse{
		Message: "Logout success",
		Status:  200,
	})
}


func (uc *UserControllers) UploadProfileImage (ctx *gin.Context) {
    file, err := ctx.FormFile("file")

    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
        return
    }

    filename := filepath.Base(file.Filename)
    savePath := fmt.Sprintf("./uploads/%s", filename)

    if err := ctx.SaveUploadedFile(file, savePath); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file"})
        return
    }

    fileURL := fmt.Sprintf("http://localhost:8080/uploads/%s", filename)

	err = uc.userUserCase.SaveProfileImage(ctx.GetString("user_id"), ctx.GetString("user_type"), fileURL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file"})
		return
	}

    ctx.JSON(http.StatusOK, gin.H{"fileUrl": fileURL})
}

func (uc *UserControllers) GetEducators(ctx *gin.Context) {
	pageNo := ctx.Query("page")
	pageSize := ctx.Query("limit")
    search := ctx.Query("search")

	if pageNo == "" {
		pageNo = "1"
	}
	if pageSize == "" {
		pageSize = "10"
	}

	educators, pagination, err := uc.userUserCase.GetEducators(pageNo, pageSize, search)
	if err != nil {
		ctx.JSON(500, domain.ErrorResponse{
			Message: err.Error(),
			Status: 500,

		})
		ctx.Abort()
	} else {
		ctx.JSON(http.StatusOK, domain.SuccessResponse{
			Status: http.StatusOK,
			Data: map[string]interface{}{
				"educators" : educators, 
				"pagination" : pagination,
			},
			Message: "educators fetched successfully.",
		})
	}
}

func (uc *UserControllers) GetEducatorById(ctx *gin.Context) {
    id := ctx.Param("id")
    course, err := uc.userUserCase.GetEducatorById(id)

    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, course)
}

func (uc *UserControllers) SaveReview(ctx *gin.Context) {
	var review domain.Review

	err := ctx.ShouldBind(&review)
	if err != nil{
		ctx.JSON(400, domain.ErrorResponse{
            Message: "Bad request",
            Status: 500,
        })
	}

	edu_id := ctx.Param("id")
	
	err = uc.userUserCase.SaveReview(review, edu_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Review Saved successfully"})
}

func  (uc *UserControllers) GetProfile(c *gin.Context) {
	user_id := c.GetString("user_id")
	if user_id == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	user_type := c.GetString("user_type")
	if user_type == "educator" {
		profile, err := uc.userUserCase.GetEducatorProfile(user_id)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch profile"})
			return
		}
		c.JSON(http.StatusOK, profile)
	}else{
		profile, err := uc.userUserCase.GetStudentProfile(user_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch profile"})
			return
		}
		c.JSON(http.StatusOK, profile)
		return
	}
}

func  (uc *UserControllers) UpdateProfile(c *gin.Context) {
	user_id := c.GetString("user_id")
	user_type := c.GetString("user_type")
	if user_id == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var educator domain.EducatorProfile
	var student domain.StudentProfile

	if user_type == "educator" {
		if err := c.BindJSON(&educator); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		updatedProfile, err := uc.userUserCase.UpdateEducatorProfile(user_id, educator)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update profile"})
			return
		}
		c.JSON(http.StatusOK, updatedProfile)
	} else {
		if err := c.BindJSON(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		updatedProfile, err := uc.userUserCase.UpdateStudentProfile(user_id, student)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update profile"})
			return
		}
		c.JSON(http.StatusOK, updatedProfile)
	}
}

func (uc *UserControllers) SetAvailability(c *gin.Context) {
    user_id := c.GetString("user_id")
    if user_id == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    var requestBody struct {
        Availability   string `json:"availability"`
    }

    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }
    err := uc.userUserCase.SetAvailability(user_id, requestBody.Availability)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to set availability", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Availability set successfully"})
}

func (uc *UserControllers) GetSchedules(ctx *gin.Context) {
    user_id := ctx.GetString("user_id")
    if user_id == ""{
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
    }
    schedules, err := uc.userUserCase.GetEducatorSchedules(user_id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch schedules"})
        return
    }

    ctx.JSON(http.StatusOK, schedules)
}

func (uc *UserControllers) GetStudentSchedules(ctx *gin.Context) {
    user_id := ctx.GetString("user_id")
    if user_id == ""{
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
    }
    schedules, err := uc.userUserCase.GetStudentSchedules(user_id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch schedules"})
        return
    }

    ctx.JSON(http.StatusOK, schedules)
}

func (uc *UserControllers) CancelSchedule(ctx *gin.Context) {
	user_id := ctx.GetString("user_id")
    if user_id == ""{
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
    }
    scheduleId := ctx.Param("id") 

    err := uc.userUserCase.CancelEducatorSchedule(scheduleId, user_id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to cancel schedule"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Schedule canceled successfully"})
}

func (uc *UserControllers) CancelEducatorSchedule(ctx *gin.Context) {
	user_id := ctx.GetString("user_id")
    if user_id == ""{
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
    }
    scheduleId := ctx.Param("id") 

    err := uc.userUserCase.CancelSchedule(scheduleId, user_id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to cancel schedule"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Schedule canceled successfully"})
}


func (uc *UserControllers) GetStudentsByCourses(c *gin.Context) {
	user_id := c.GetString("user_id")
	if user_id == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

    studentsByCourses, err := uc.userUserCase.FetchStudentsByCourses(user_id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch students"})
        return
    }

    c.JSON(http.StatusOK, studentsByCourses)
}

func (uc *UserControllers) GetUserProfile(c *gin.Context) {
    userID  := c.GetString("user_id")
    if userID == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
        return
    }
	
    user, err := uc.userUserCase.GetUserProfile(userID)
    if err != nil || user == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "profileImage": user.ProfileImage, 
        "role":         user.Role,        
    })
}

func (uc *UserControllers) GetTopEducators(c *gin.Context) {
    topEducators, err := uc.userUserCase.GetTopEducatorsUseCase()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching top educators"})
        return
    }
    c.JSON(http.StatusOK, topEducators)
}

func (uc *UserControllers) GetCourseProgress(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	
	courseProgress, err := uc.userUserCase.GetEnrolledCoursesProgress(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching enrolled courses progress"})
		return
	}

	c.JSON(http.StatusOK, courseProgress)
}

func  (uc *UserControllers) ScheduleSession(c *gin.Context) {
	user_id := c.GetString("user_id")
	if user_id == ""{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var req struct {
		EducatorID        string `json:"educatorId"`
		Availability      string `json:"availability"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := uc.userUserCase.ScheduleSession(user_id, req.EducatorID, req.Availability)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Scheduled successfully"})
}