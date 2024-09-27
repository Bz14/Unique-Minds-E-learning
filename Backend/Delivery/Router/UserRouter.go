package router

import (
	controller "unique-minds/Delivery/Controller"
	infrastructures "unique-minds/Infrastructures"
	repositories "unique-minds/Repositories"
	useCase "unique-minds/UseCases"
	util "unique-minds/Utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRouter(server *gin.Engine, user_collection *mongo.Collection, config *infrastructures.Config) {
	validator := util.NewValidator()
	passwordService := util.NewPasswordService()
	userRepository := repositories.NewUserRepository(user_collection, *config)
	userUseCase := useCase.NewUserUseCase(userRepository, validator, passwordService)
	userController := controller.NewUserController(userUseCase)
	server.POST("/signUp", userController.SignUp)
}