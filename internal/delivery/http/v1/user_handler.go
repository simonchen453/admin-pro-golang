package v1

import (
	"net/http"

	"admin-pro/internal/usecase"
	"admin-pro/pkg/response"
	"github.com/gin-gonic/gin"
)

// UserHandler 处理用户相关的请求
type UserHandler struct {
	userUsecase usecase.UserUsecase
}

// NewUserHandler 构造函数
func NewUserHandler(r *gin.Engine, uc usecase.UserUsecase, authMw gin.HandlerFunc) {
	handler := &UserHandler{userUsecase: uc}
	g := r.Group("/api/v1")
	g.Use(authMw)
	{
		g.GET("/auth/userinfo", handler.GetUserInfo) // 获取用户信息不需要额外权限
	}
}

// GetUserInfo 获取当前登录用户的信息（包含角色和权限）
// @Summary 获取当前用户信息
// @Tags user
// @Security BearerAuth
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/v1/auth/userinfo [get]
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
