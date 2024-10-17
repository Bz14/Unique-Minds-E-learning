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
	SMTPHost        string
	SMTPPort        int
	EmailFrom       string
	EmailPassword   string
	ServerHost      string
	TokenTTlL       int
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
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	emailFrom := os.Getenv("EMAIL_FROM")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	serverHost := os.Getenv("SERVER_HOST")
	tokenTTL := os.Getenv("TOKEN_TTL")

	host, err := strconv.Atoi(smtpPort)
	if err != nil {
		return nil, errors.New("invalid smtp port value")
	}

	port_str, err := strconv.Atoi(port)
	if err != nil {
		return nil, errors.New("invalid port value")
	}

	timeOut_str, err := strconv.Atoi(timeOut)
	if err != nil {
		return nil, errors.New("invalid timeout value")
	}

	ttl, err := strconv.Atoi(tokenTTL)
	if err != nil {
		return nil, errors.New("invalid token ttl value")
	} 

	return &Config{
		DatabaseURL:     url,
		DatabaseName:    dbName,
		UsersCollection: usersCollection,
		ServerPort:      port_str,
		TimeOut:         timeOut_str,
		SMTPHost:        smtpHost,
		SMTPPort:        host,
		EmailFrom:       emailFrom,
		EmailPassword:   emailPassword,
		ServerHost:      serverHost,
		TokenTTlL:       ttl,
	}, nil
}