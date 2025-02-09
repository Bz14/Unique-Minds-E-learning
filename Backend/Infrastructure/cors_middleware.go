// Core represents the core structure for the CORS middleware.
//
// CoreInterface defines the interface for the CORS middleware.
//
// NewCorsMiddleware creates a new instance of Core.
//
// CORSMiddleware returns a gin.HandlerFunc that handles CORS (Cross-Origin Resource Sharing) requests.
// It sets the appropriate headers to allow cross-origin requests from the specified frontend URL.
// It also handles preflight requests by responding with a 204 status code.
package infrastructure

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Core struct {
}

type CoreInterface interface {
    CORSMiddleware() gin.HandlerFunc
}

func NewCorsMiddleware() *Core {
    return &Core{}
}


func (c *Core) CORSMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        fmt.Println("Lopping through the CORS middleware")
        fmt.Println("CORS middleware executed for:", ctx.Request.Method, ctx.Request.URL.Path)

        // Update the Access-Control-Allow-Origin header to match your frontend's URL
        ctx.Writer.Header().Set("Access-Control-Allow-Origin", "https://unique-minds.vercel.app")
        ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        ctx.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")

        // Handle preflight requests
        if ctx.Request.Method == "OPTIONS" {
            fmt.Println("Preflight request handled")
            ctx.AbortWithStatus(204)
            return
        }

        // Continue to the next middleware or handler
        ctx.Next()
    }
}
