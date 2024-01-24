package initializer

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HeaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requiredHeaders := []string{"Authorization", "Content-Type"}

		for _, header := range requiredHeaders {
			if c.GetHeader(header) == "" {
				// If any required header is missing, return a 400 Bad Request
				c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Missing header: %s", header)})
				c.Abort()
				return
			}
		}

		// Continue to the next handler
		c.Next()
	}
}
