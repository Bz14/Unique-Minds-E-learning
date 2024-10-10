package controller

import (
	"net/http"
	domain "unique-minds/Domain"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type UserController struct {
	userUseCase domain.UserUseCaseInterface
}

type UserControllerInterface interface{}

func NewUserController(useCase domain.UserUseCaseInterface) *UserController {
	return &UserController{
		userUseCase: useCase,
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

	err := uc.userUseCase.SignUp(user)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		})
		return
	}
	ctx.JSON(http.StatusCreated, domain.SuccessResponse{
		Message: "User created successfully",
		Status: http.StatusCreated,
	})
}
