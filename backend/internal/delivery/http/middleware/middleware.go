package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"admin-pro/internal/config"
	"admin-pro/pkg/response"
	"admin-pro/pkg/utils"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin) // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE, PATCH")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
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
