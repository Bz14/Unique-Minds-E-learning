package infrastructures

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)
type Config struct {
	DatabaseURL     string
	DatabaseName    string
	UsersCollection string
	ServerPort      int
	TimeOut 	   int
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("error loading .env file, assuming production environment")
	}

	url := os.Getenv("DATABASE_URL")
	dbName := os.Getenv("DATABASE_NAME")
	usersCollection := os.Getenv("USER_COLLECTION")
	port := os.Getenv("PORT")
	timeOut := os.Getenv("TIMEOUT")

	port_str, err := strconv.Atoi(port)
	if err != nil {
		return nil, errors.New("invalid port value")
	}

	timeOut_str, err := strconv.Atoi(timeOut)
	if err != nil {
		return nil, errors.New("invalid timeout value")
	}

	return &Config{
		DatabaseURL:     url,
		DatabaseName:    dbName,
		UsersCollection: usersCollection,
		ServerPort:      port_str,
		TimeOut:         timeOut_str,
	}, nil
}