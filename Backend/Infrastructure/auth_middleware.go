package infrastructure

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Auth struct{
	env Config
}
type AuthInterface interface{
	AuthenticationMiddleware() gin.HandlerFunc
}

// NewAuthMiddleware creates a new instance of Auth middleware with the provided configuration.
// 
// Parameters:
// - env: A Config object that contains the environment configuration.
//
// Returns:
// - A pointer to an Auth struct initialized with the provided configuration.
func NewAuthMiddleware (env Config)*Auth{
	return &Auth{
		env : env,
	}
}




// AuthenticationMiddleware is a middleware function for the Gin framework that
// handles authentication using bearer tokens. It checks for the presence of an
// "Authorization" header in the incoming request, validates the token, and
// extracts user information from it. If the token is missing, invalid, or
// unauthorized, it responds with an appropriate HTTP status and message.
//
// The middleware performs the following steps:
// 1. Checks if the "Authorization" header is present.
// 2. Validates the format of the "Authorization" header (must be "Bearer <token>").
// 3. Extracts and validates the token using the provided secret.
// 4. Sets the user ID and user type in the request context if the token is valid.
// 5. Aborts the request with an unauthorized status if any validation fails.
//
// Returns a Gin handler function that can be used as middleware in routes.
func (authenticate *Auth) AuthenticationMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")
		if authHeader == ""{
			c.JSON(http.StatusUnauthorized, gin.H{
				"message" : "Authorization header required",
			})
			c.Abort()
			return
		}
		auth := strings.Split(authHeader, " ")
		if len(auth) != 2 || strings.ToLower(auth[0]) != "bearer"{
			c.JSON(http.StatusUnauthorized, gin.H{
				"message" : "Invalid authorization header.",
			})
			c.Abort()
			return
		}
		claims, err := ExtractIDFromToken(auth[1], authenticate.env.AccessTokenSecret)
		if err != nil{
			c.JSON(http.StatusUnauthorized, gin.H{
				"message" : "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Set("user_id" , claims["id"])
		c.Set("user_type", claims["user_type"])
		c.Next()
	}
}
