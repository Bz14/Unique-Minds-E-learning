package main

import (
	"fmt"

	routers "unique-minds/Delivery/Routers"
	infrastructure "unique-minds/Infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	config, err := infrastructure.LoadEnv()
	if err != nil {
		fmt.Print("Error in env.load")
	}
	db := infrastructure.NewDatabase()
	coreMiddleWare := infrastructure.NewCorsMiddleware().CORSMiddleware()
	
	server.Use(coreMiddleWare)
	
	serverGroup := server.Group("api")
	routers.Routers(serverGroup, db, config)

	server.Run(fmt.Sprintf(":%d", config.Port))

}