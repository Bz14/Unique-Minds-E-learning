package main

import (
	"fmt"
	router "unique-minds/Delivery/Routers"
	infrastructure "unique-minds/Infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {
    config, err := infrastructure.LoadEnv()
    if err != nil {
        fmt.Println(err.Error())
    }
    database := infrastructure.NewDatabase()

    corsMiddleware := infrastructure.NewCorsMiddleware()

    server := gin.Default()
    server.Use(corsMiddleware.CORSMiddleware())
    server.Static("/uploads", "./uploads")
    serverGroup := server.Group("/api")
    router.Routers(serverGroup, database, config)
    server.Run(fmt.Sprintf(":%d", config.Port))
}
