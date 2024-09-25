package router

import (
	domain "unique-minds/Domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Router struct{}

type RouterInterface interface{}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Init(server *gin.Engine, db *mongo.Database, config *domain.Config) {

}