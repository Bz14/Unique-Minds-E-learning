package infrastructures

import (
	"errors"
	"os"
	"strconv"
	domain "unique-minds/domain"

	"github.com/joho/godotenv"
)

func LoadConfig() (domain.Config, error) {
	err := godotenv.Load()
	if err != nil {
		return domain.Config{}, errors.New("Error loading .env file, assuming production environment")
	}

	url := os.Getenv("DATABASE_URL")
	dbName := os.Getenv("DATABASE_NAME")
	usersCollection := os.Getenv("USERS_COLLECTION")
	port := os.Getenv("PORT")

	port_str, err := strconv.Atoi(port)
	if err != nil {
		return domain.Config{}, errors.New("Invalid port value")
	}

	return domain.Config{
		DatabaseURL:     url,
		DatabaseName:    dbName,
		UsersCollection: usersCollection,
		ServerPort:      port_str,
	}, nil
}