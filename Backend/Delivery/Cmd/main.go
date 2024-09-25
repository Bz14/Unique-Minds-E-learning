package cmd

import (
	"fmt"
	router "unique-minds/Delivery/Router"
	infrastructures "unique-minds/infrastructures"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := infrastructures.LoadConfig()

	if err != nil {
		panic(err)
	}

	database := infrastructures.NewDatabase()
	db, err := database.CreateDB(config.DatabaseURL, config.DatabaseName)

	if err != nil {
		panic(err)
	}

	server := gin.Default()

	route := router.NewRouter()

	route.Init(server, db, &config)

	server.Run(fmt.Sprintf(":%d", config.ServerPort))
}