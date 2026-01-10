package middleware

import (
	"net/http"
	"strings"

	"admin-pro/internal/config"
	"admin-pro/internal/usecase"
	"admin-pro/pkg/response"
	"admin-pro/pkg/utils"
	"github.com/gin-gonic/gin"
)

// Cors 处理跨域资源共享（CORS）配置
// 在生产环境中，可通过 ALLOWED_ORIGINS 环境变量设置允许的来源（逗号分隔）
// 示例: ALLOWED_ORIGINS=https://example.com,https://admin.example.com
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

// isOriginAllowed 检查给定的来源是否被允许
// TODO: 从配置或环境变量加载允许的来源列表
// 示例: ALLOWED_ORIGINS=https://example.com,https://admin.example.com
func isOriginAllowed(origin string) bool {
	// 开发环境：允许 localhost
	// 生产环境：使用具体域名
	allowedOrigins := []string{
		"http://localhost:3000",
		"http://localhost:8080",
		"http://127.0.0.1:3000",
		"http://127.0.0.1:8080",
		// 在此添加生产环境域名
	}

	for _, allowed := range allowedOrigins {
		if origin == allowed {
			return true
		}
	}
	return false
}

// JWTAuth JWT 认证中间件（增强版）
// 验证 Token 并加载用户权限到 Context
func JWTAuth(cfg *config.Config, userUsecase usecase.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 检查 Cookie
		tokenString, err := c.Cookie("admin-pro-token")
		if err != nil {
			// 2. 检查 Header
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

		// 3. 加载用户信息和权限（新增）
		userInfo, err := userUsecase.GetUserInfo(c.Request.Context(), claims.UserID)
		if err != nil || userInfo == nil {
			response.Fail(c, http.StatusOK, "401", "用户信息获取失败")
			c.Abort()
			return
		}

		// 4. 存入 Context
		c.Set("userID", claims.UserID)
		c.Set("userDomain", claims.UserDomain)
		c.Set("loginName", claims.LoginName)
		c.Set("permissions", userInfo.Permissions) // 存入权限列表
		c.Set("roles", userInfo.Roles)             // 存入角色列表
		c.Next()
	}
}

// RequirePermission 权限检查中间件
// 检查用户是否拥有指定权限，支持超级管理员通配符
func RequirePermission(requiredPermission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 获取用户权限列表
		permissionsVal, exists := c.Get("permissions")
		if !exists {
			response.Fail(c, http.StatusForbidden, "403", "无权限信息，请重新登录")
			c.Abort()
			return
		}

		permissions, ok := permissionsVal.([]string)
		if !ok {
			response.Fail(c, http.StatusForbidden, "403", "权限信息格式错误")
			c.Abort()
			return
		}

		// 2. 检查是否拥有所需权限
		if hasPermission(permissions, requiredPermission) {
			c.Next()
			return
		}

		// 3. 权限不足，拒绝访问
		response.Fail(c, http.StatusForbidden, "403", "权限不足")
		c.Abort()
	}
}

// hasPermission 检查权限列表中是否包含所需权限
// 支持通配符：*:*:* (超级管理员), system:*:* (系统模块所有权限), system:dept:* (部门所有权限)
func hasPermission(permissions []string, required string) bool {
	for _, perm := range permissions {
		// 完全匹配
		if perm == required {
			return true
		}

		// 超级管理员通配符
		if perm == "*:*:*" {
			return true
		}

		// 模块级通配符 (如: system:*:*)
		if strings.HasSuffix(perm, ":*:*") {
			module := strings.Split(perm, ":")[0]
			if strings.HasPrefix(required, module+":") {
				return true
			}
		}

		// 资源级通配符 (如: system:dept:*)
		if strings.HasSuffix(perm, ":*") {
			prefix := strings.TrimSuffix(perm, "*")
			if strings.HasPrefix(required, prefix) {
				return true
			}
		}
	}
	return false
}
