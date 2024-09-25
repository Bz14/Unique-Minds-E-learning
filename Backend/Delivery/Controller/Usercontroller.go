package controller

type UserController struct{}

type UserControllerInterface interface{}

func NewUserController() *UserController {
	return &UserController{}
}
