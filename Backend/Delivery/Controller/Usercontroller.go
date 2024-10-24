package controller

import (
	"net/http"
	domain "unique-minds/Domain"
	infrastructures "unique-minds/Infrastructures"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type UserController struct {
	userUseCase domain.UserUseCaseInterface
	config infrastructures.Config
}

type UserControllerInterface interface{}

func NewUserController(useCase domain.UserUseCaseInterface, config infrastructures.Config) *UserController {
	return &UserController{
		userUseCase: useCase,
		config: config,
	}
}


func (uc *UserController) SignUp(ctx *gin.Context) {
	var signUpRequest domain.SignUpRequest
	var user domain.User
	if err := ctx.ShouldBindJSON(&signUpRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		})
		return
	}
	if err := copier.Copy(&user, &signUpRequest); err != nil{
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		})
		return
	}

	isVerified, err := uc.userUseCase.SignUp(user)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		})
		return
	}
	config, err := infrastructures.LoadConfig()
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		})
		return
	}
	if isVerified{
		ctx.Redirect(http.StatusFound, config.Redirect)
	}

	ctx.JSON(http.StatusCreated, domain.SuccessResponse{
		Message: "User created successfully",
		Status: http.StatusCreated,
	})
}


func (uc *UserController) FindEmail(ctx *gin.Context){
	email := ctx.Query("email")

	err := uc.userUseCase.FindEmail(email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		})
		return
	}
	ctx.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Email already exists",
		Data: []string{},
		Status: http.StatusOK,
	})
}


func (uc *UserController) VerifyEmail(ctx *gin.Context){
	token := ctx.Query("token")
	err := uc.userUseCase.VerifyEmail(token)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		})
		return
	}
	config, err := infrastructures.LoadConfig()
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		})
		return
	}
	ctx.Redirect(http.StatusFound, config.RedirectLogin)
}


func (uc *UserController) Login(ctx *gin.Context){
	var loginRequest domain.LoginRequest

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		})
		return
	}

	loginResponse, err := uc.userUseCase.Login(loginRequest)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		})
		return
	}
	ctx.SetCookie("refreshToken", loginResponse.RefreshToken, 3600*24*7, "/", uc.config.CookieDomain, uc.config.Environment, true)

	ctx.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Login successful",
		Data: loginResponse.AccessToken,
		Status: http.StatusOK,
	})
}

func (uc *UserController) UpdateRole(ctx *gin.Context){
	email := ctx.Query("email")
	role := ctx.Query("role")
	err := uc.userUseCase.UpdateRole(email, role)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		})
		return
	}
	ctx.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Role updated successfully",
		Status: http.StatusOK,
	})
}