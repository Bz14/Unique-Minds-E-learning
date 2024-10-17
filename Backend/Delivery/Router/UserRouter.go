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

func NewUserRouter(server *gin.RouterGroup, database *infrastructures.Database, db *mongo.Database, config *infrastructures.Config) {
	user_collection, err := database.CreateCollection(db, config.UsersCollection)
	if err != nil{	
		panic(err)
	}
	unverified_collection, err := database.CreateCollection(db, config.UnverifiedCollection)
	if err != nil{	
		panic(err)
	}
	client, err := database.Connect(config.DatabaseURL)
	if err != nil{
		panic(err)
	}
	session, err := client.StartSession()
	if err != nil{
		panic(err)
	}
	validator := util.NewValidator()
	passwordService := util.NewPasswordService()
	userRepository := repositories.NewUserRepository(session, user_collection, unverified_collection, *config)
	userUseCase := useCase.NewUserUseCase(userRepository, validator, passwordService)
	userController := controller.NewUserController(userUseCase)
	authGroup := server.Group("/auth")
	authGroup.POST("/signup", userController.SignUp)
	authGroup.GET("/email", userController.FindEmail)
	authGroup.GET("/reset", userController.VerifyEmail)
}