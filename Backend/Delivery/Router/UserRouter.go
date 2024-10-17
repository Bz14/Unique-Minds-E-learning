package router

import (
	"fmt"
	controller "unique-minds/Delivery/Controller"
	infrastructures "unique-minds/Infrastructures"
	repositories "unique-minds/Repositories"
	useCase "unique-minds/UseCases"
	util "unique-minds/Utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRouter(server *gin.RouterGroup, database *infrastructures.Database, db *mongo.Database, config *infrastructures.Config) {
	fmt.Println(config)
	user_collection, err := database.CreateCollection(db, config.UsersCollection)
	if err != nil{
		panic(err)
	}
	validator := util.NewValidator()
	passwordService := util.NewPasswordService()
	userRepository := repositories.NewUserRepository(user_collection, *config)
	userUseCase := useCase.NewUserUseCase(userRepository, validator, passwordService)
	userController := controller.NewUserController(userUseCase)
	authGroup := server.Group("/auth")
	authGroup.POST("/signup", userController.SignUp)
	authGroup.GET("/email", userController.FindEmail)
}