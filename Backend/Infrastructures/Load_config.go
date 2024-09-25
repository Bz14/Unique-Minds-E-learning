package infrastructures

import (
	"log"
	"os"
	domain "unique-minds/domain"

	"github.com/joho/godotenv"
)

func LoadConfig() domain.Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, assuming production environment")
	}

	url := os.Getenv("DATABASE_URL")
	dbName := os.Getenv("DATABASE_NAME")
	usersCollection := os.Getenv("USERS_COLLECTION")


	return domain.Config{
		DatabaseURL:     url,
		DatabaseName:    dbName,
		UsersCollection: usersCollection,
	}
}