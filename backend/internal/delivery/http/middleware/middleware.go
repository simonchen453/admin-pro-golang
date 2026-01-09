package middleware

import (
	"net/http"
	"strings"

	"admin-pro/internal/config"
	"admin-pro/pkg/response"
	"admin-pro/pkg/utils"
	"github.com/gin-gonic/gin"
)

// Cors handles Cross-Origin Resource Sharing with configurable allowed origins
// In production, set ALLOWED_ORIGINS environment variable with comma-separated origins
// Example: ALLOWED_ORIGINS=https://example.com,https://admin.example.com
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")

		if origin != "" {
			// Check if origin is allowed
			if isOriginAllowed(origin) {
				c.Header("Access-Control-Allow-Origin", origin)
				c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE, PATCH")
				c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
				c.Header("Access-Control-Allow-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
				c.Header("Access-Control-Allow-Credentials", "true")
			}
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

// isOriginAllowed checks if the given origin is allowed
// TODO: Load allowed origins from config or environment variable
// Example: ALLOWED_ORIGINS=https://example.com,https://admin.example.com
func isOriginAllowed(origin string) bool {
	// For development, allow localhost
	// For production, use specific domains
	allowedOrigins := []string{
		"http://localhost:3000",
		"http://localhost:8080",
		"http://127.0.0.1:3000",
		"http://127.0.0.1:8080",
		// Add production origins here
	}

	for _, allowed := range allowedOrigins {
		if origin == allowed {
			return true
		}
	}
	return false
}

func JWTAuth(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Check Cookie
		tokenString, err := c.Cookie("admin-pro-token")
		if err != nil {
			// 2. Check Header
			authHeader := c.GetHeader("Authorization")
			if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
				tokenString = authHeader[7:]
			}
		}

		if tokenString == "" {
			response.Fail(c, http.StatusOK, "401", "未登录或Token已过期")
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(tokenString, cfg)
		if err != nil {
			response.Fail(c, http.StatusOK, "401", "Token无效")
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("userDomain", claims.UserDomain)
		c.Set("loginName", claims.LoginName)
		c.Next()
	}
}
