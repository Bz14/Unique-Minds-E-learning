package controller

import domain "unique-minds/Domain"

type UserController struct {
	userController domain.UserUseCaseInterface
}

type UserControllerInterface interface{}

func NewUserController(userUseCase domain.UserUseCaseInterface) *UserController {
	return &UserController{
		userController: userUseCase,
	}
}
