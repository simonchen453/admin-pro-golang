package middleware

import (
	"time"

	"admin-pro/pkg/logger"
	"github.com/gin-gonic/gin"
)

// RequestLogger is a middleware that logs HTTP requests
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Get status code
		status := c.Writer.Status()
		method := c.Request.Method
		clientIP := c.ClientIP()

		// Build log message
		if query != "" {
			path = path + "?" + query
		}

		// Log based on status code
		if status >= 500 {
			logger.Error("[%s] %s %s - Status: %d - Latency: %v - Client: %s",
				method, path, c.Request.Proto, status, latency, clientIP)
		} else if status >= 400 {
			logger.Warn("[%s] %s %s - Status: %d - Latency: %v - Client: %s",
				method, path, c.Request.Proto, status, latency, clientIP)
		} else {
			logger.Info("[%s] %s %s - Status: %d - Latency: %v - Client: %s",
				method, path, c.Request.Proto, status, latency, clientIP)
		}
	}
}
