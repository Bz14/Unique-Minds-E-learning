// Setting up cors middleware

package infrastructures

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type CorsMiddleware struct{

}

type CorsMiddlewareInterface interface{
	CorsMiddleware() gin.HandlerFunc
}


func NewCorsMiddleware() CorsMiddlewareInterface{
	return &CorsMiddleware{}
}


func (c CorsMiddleware) CorsMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		fmt.Println("CORS middleware executed for:", c.Request.Method, c.Request.URL.Path)
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
            fmt.Println("Preflight request handled")
            c.AbortWithStatus(204)
            return
        }
		c.Next()

	}
}