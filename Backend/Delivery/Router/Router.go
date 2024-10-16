package router

import (
	infrastructures "unique-minds/Infrastructures"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Router struct{}

type RouterInterface interface{}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Init(server *gin.RouterGroup, database *infrastructures.Database, db *mongo.Database, config *infrastructures.Config) {
	NewUserRouter(server, database, db, config)
} 