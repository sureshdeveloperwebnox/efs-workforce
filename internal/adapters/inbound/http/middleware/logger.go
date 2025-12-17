package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger provides structured logging for HTTP requests
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Get status code
		statusCode := c.Writer.Status()

		// Build query string
		if raw != "" {
			path = path + "?" + raw
		}

		// Log request details
		log.Printf("[HTTP] %s | %3d | %13v | %15s | %-7s %s",
			time.Now().Format("2006/01/02 - 15:04:05"),
			statusCode,
			latency,
			c.ClientIP(),
			c.Request.Method,
			path,
		)
	}
}
