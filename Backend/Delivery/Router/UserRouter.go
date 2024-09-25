package router

import (
	infrastructures "unique-minds/Infrastructures"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSignUpRouter(server *gin.Engine, db *mongo.Database, config *infrastructures.Config) {
	user_collection := 
}
