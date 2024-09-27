package router

import (
	controller "unique-minds/Delivery/Controller"
	repositories "unique-minds/Repositories"
	useCase "unique-minds/UseCases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRouter(server *gin.Engine, user_collection *mongo.Collection) {
	userRepository := repositories.NewUserRepository(user_collection)
	userUseCase := useCase.NewUserUseCase(userRepository)
	userController := controller.NewUserController(userUseCase)

	server.POST("/signUp", userController.SignUp)
}