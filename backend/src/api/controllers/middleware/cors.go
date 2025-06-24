package middleware

import (
	"backend/src/api/config"
	"fmt"
	"slices"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	allowedOrigins := config.AllowedOrigin
	fmt.Println("CORS middleware initialized with allowed origins:", allowedOrigins)
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		allowed := slices.Contains(allowedOrigins, origin)
		if allowed {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, csrf_token, hx-current-url, hx-request")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		if !allowed {
			fmt.Println("CORS middleware: Origin not allowed:", origin)
			c.Abort()
			return
		}
		c.Next()
	}
}
