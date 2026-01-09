package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"admin-pro/internal/usecase"
	"admin-pro/pkg/response"
)

type MonitorHandler struct {
	monitorUsecase usecase.MonitorUsecase
}

func NewMonitorHandler(r *gin.Engine, uc usecase.MonitorUsecase, mw gin.HandlerFunc) {
	handler := &MonitorHandler{
		monitorUsecase: uc,
	}
	
	// Server Info
	r.GET("/api/v1/monitor/server", mw, handler.GetServerInfo)

	// Online Sessions
	gSession := r.Group("/api/v1/monitor/online") // verify path with frontend
	gSession.Use(mw)
	{
		gSession.GET("/list", handler.ListSessions)
		gSession.DELETE("/:id", handler.ForceLogout)
	}
}

// @Summary Get server info
// @Tags monitor
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

// @Summary Get online sessions
// @Tags monitor
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

// @Summary Force logout session
// @Tags monitor
// @Param id path string true "Session ID"
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
