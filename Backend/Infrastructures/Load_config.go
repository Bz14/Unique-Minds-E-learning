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
	UnverifiedCollection string
	ActiveUserCollection     string
	CourseCollection         string
	StudentProfileCollection string
	EducatorProfileCollection string
	ServerPort      int
	TimeOut 	   int
	SMTPHost        string
	SMTPPort        int
	EmailFrom       string
	EmailPassword   string
	ServerHost      string
	TokenTTlL       int
	Redirect        string
	RedirectLogin   string
	ClientId 	  string
	ClientSecret 	  string
	RedirectUrl 	  string
	State 		 string
	RoleRedirect		 string
	Secret          string
	AccessTokenExpire int
	RefreshTokenExpire int
	CookieDomain string
	Environment    bool
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("error loading .env file, assuming production environment")
	}

	url := os.Getenv("DATABASE_URL")
	dbName := os.Getenv("DATABASE_NAME")
	usersCollection := os.Getenv("USER_COLLECTION")
	activeUserCollection := os.Getenv("ACTIVE_USER_COLLECTION")
	courseCollection := os.Getenv("COURSE_COLLECTION")
	studentProfileCollection := os.Getenv("STUDENT_PROFILE_COLLECTION")
	educatorProfileCollection := os.Getenv("EDUCATOR_PROFILE_COLLECTION")
	port := os.Getenv("PORT")
	timeOut := os.Getenv("TIMEOUT")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	emailFrom := os.Getenv("EMAIL_FROM")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	serverHost := os.Getenv("SERVER_HOST")
	tokenTTL := os.Getenv("TOKEN_TTL")
	unverified := os.Getenv("UNVERIFIED_COLLECTION")
	redirect := os.Getenv("REDIRECT")
	redirectLogin := os.Getenv("REDIRECT_LOGIN")
	clientId := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	redirectUrl := os.Getenv("REDIRECT_URL")
	state := os.Getenv("STATE")
	role := os.Getenv("REDIRECT_ROLE")
	secret := os.Getenv("SECRET")
	accessTokenExpiry := os.Getenv("ACCESS_TOKEN_EXPIRES_IN")
	refreshTokenExpiry := os.Getenv("REFRESH_TOKEN_EXPIRES_IN")
	cookieDomain := os.Getenv("COOKIE_DOMAIN")
	env := os.Getenv("ENVIRONMENT")

	prodEnv := false
	if env == "production" {
		prodEnv = true
	}


	accessExp, err := strconv.Atoi(accessTokenExpiry)
	if err != nil {
		return nil, errors.New("invalid expiry value")

	}
	refreshExp, err := strconv.Atoi(refreshTokenExpiry)
	if err != nil {
		return nil, errors.New("invalid expiry value")
	}

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
		CourseCollection: courseCollection,
		StudentProfileCollection: studentProfileCollection,
		EducatorProfileCollection: educatorProfileCollection,
		ServerPort:      port_str,
		TimeOut:         timeOut_str,
		SMTPHost:        smtpHost,
		SMTPPort:        host,
		EmailFrom:       emailFrom,
		EmailPassword:   emailPassword,
		ServerHost:      serverHost,
		TokenTTlL:       ttl,
		UnverifiedCollection: unverified,
		Redirect: redirect,
		RedirectLogin: redirectLogin,
		ClientId: clientId,
		ClientSecret: clientSecret,
		RedirectUrl: redirectUrl,
		State: state,
		RoleRedirect: role,
		Secret: secret,
		AccessTokenExpire: accessExp,
		RefreshTokenExpire: refreshExp,
		CookieDomain: cookieDomain,
		Environment: prodEnv,
		ActiveUserCollection: activeUserCollection,
	}, nil
}