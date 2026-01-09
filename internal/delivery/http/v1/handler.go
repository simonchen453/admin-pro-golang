package v1

import (
	"net/http"

	"admin-pro/internal/config"
	"admin-pro/internal/usecase"
	"admin-pro/pkg/response"
	"github.com/gin-gonic/gin"
)

// AuthHandler 处理认证相关的请求
// 通常用于 Login, Logout
type AuthHandler struct {
	authUsecase usecase.AuthUsecase // 依赖 Usecase
	cfg         *config.Config      // 依赖配置
}

// NewAuthHandler AuthHandler 构造函数
func NewAuthHandler(r *gin.Engine, uc usecase.AuthUsecase, cfg *config.Config) {
	handler := &AuthHandler{
		authUsecase: uc,
		cfg:         cfg,
	}
	// 注册路由组 /api/v1/auth
	g := r.Group("/api/v1/auth")
	{
		g.POST("/login", handler.Login)
		g.POST("/logout", handler.Logout)
	}
}

// LoginRequest 定义登录请求 JSON 结构
type LoginRequest struct {
	LoginName string `json:"loginName" binding:"required"` // 必填校验
	Password  string `json:"password" binding:"required"`
}

// Login 处理登录请求 POST /api/v1/auth/login
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	// 1. 绑定并校验参数
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", "Invalid request parameters")
		return
	}

	// 2. 调用 Usecase 执行登录逻辑
	token, user, err := h.authUsecase.Login(c.Request.Context(), req.LoginName, req.Password)
	if err != nil {
		// 业务错误，通常返回 200 + 错误码，这里示例简化处理
		response.Fail(c, http.StatusOK, "500", "用户名或密码错误")
		return
	}

	// 3. 设置 Cookie (可选)
	c.SetCookie("admin-pro-token", token, h.cfg.JWT.Expire*3600, "/", "", false, true)

	// 4. 返回成功响应
	response.Success(c, gin.H{
		"token": token,
		"user":  user,
	})
}

// Logout 处理登出请求
func (h *AuthHandler) Logout(c *gin.Context) {
	// 清除 Cookie
	c.SetCookie("admin-pro-token", "", -1, "/", "", false, true)
	response.Success(c, nil)
}

// UserHandler 处理用户相关的请求
type UserHandler struct {
	userUsecase usecase.UserUsecase
}

// NewUserHandler 构造函数
func NewUserHandler(r *gin.Engine, uc usecase.UserUsecase, mw gin.HandlerFunc) {
	handler := &UserHandler{
		userUsecase: uc,
	}
	g := r.Group("/api/v1") // 通用路由组
	g.Use(mw)               // 使用中间件 (JWT 认证)
	{
		g.GET("/auth/userinfo", handler.GetUserInfo) // 获取当前用户信息
	}
}

// GetUserInfo 获取当前登录用户的信息（包含角色和权限）
func (h *UserHandler) GetUserInfo(c *gin.Context) {
	// 从 Gin Context 中获取 userID (由 JWT 中间件解析并存入)
	userID, exists := c.Get("userID")
	if !exists {
		response.Fail(c, http.StatusUnauthorized, "401", "Unauthorized")
		return
	}

	// 调用 Usecase
	userInfo, err := h.userUsecase.GetUserInfo(c.Request.Context(), userID.(string))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, userInfo)
}
