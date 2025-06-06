package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func RequireAPIKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := os.Getenv("API_KEY") // read from environment variable
		if c.GetHeader("X-API-KEY") != apiKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "API key required"})
			return
		}
		c.Next()
	}
}
