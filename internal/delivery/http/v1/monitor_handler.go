package v1

import (
	"net/http"

	"admin-pro/internal/usecase"
	"admin-pro/pkg/response"
	"github.com/gin-gonic/gin"
)

type MonitorHandler struct {
	monitorUsecase usecase.MonitorUsecase
}

func NewMonitorHandler(r *gin.Engine, uc usecase.MonitorUsecase, authMw gin.HandlerFunc, permMw func(string) gin.HandlerFunc) {
	handler := &MonitorHandler{monitorUsecase: uc}

	// 服务器信息
	r.GET("/api/v1/monitor/server", authMw, permMw("monitor:server:query"), handler.GetServerInfo)

	// 在线用户
	gSession := r.Group("/api/v1/monitor/online")
	gSession.Use(authMw)
	{
		gSession.GET("/list", permMw("monitor:online:list"), handler.ListSessions)
		gSession.DELETE("/:id", permMw("monitor:online:forceLogout"), handler.ForceLogout)
	}
}

// @Summary 获取服务器监控信息
// @Success 200 {object} response.Response
// @Router /api/v1/monitor/server [get]
func (h *MonitorHandler) GetServerInfo(c *gin.Context) {
	info, err := h.monitorUsecase.GetServerInfo(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, info)
}

// @Summary 获取在线用户列表
// @Success 200 {object} response.Response
// @Router /api/v1/monitor/online/list [get]
func (h *MonitorHandler) ListSessions(c *gin.Context) {
	list, err := h.monitorUsecase.GetSessionList(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, list)
}

// @Summary 强退用户
// @Tags monitor
// @Param id path string true "会话ID"
// @Success 200 {object} response.Response
// @Router /api/v1/monitor/online/{id} [delete]
func (h *MonitorHandler) ForceLogout(c *gin.Context) {
	id := c.Param("id")
	if err := h.monitorUsecase.DeleteSession(c.Request.Context(), id); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}
