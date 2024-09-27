package controller

import (
	"net/http"
	domain "unique-minds/Domain"

	"github.com/gin-gonic/gin"
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
	if err := ctx.ShouldBindJSON(&signUpRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		})
		return
	}
	err := uc.userUseCase.SignUp(signUpRequest)
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
