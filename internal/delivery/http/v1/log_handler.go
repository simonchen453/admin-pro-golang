package v1

import (
	"net/http"

	"admin-pro/internal/usecase"
	"admin-pro/pkg/response"
	"github.com/gin-gonic/gin"
)

type LogHandler struct {
	logUsecase usecase.LogUsecase
}

func NewLogHandler(r *gin.Engine, uc usecase.LogUsecase, mw gin.HandlerFunc) {
	handler := &LogHandler{
		logUsecase: uc,
	}

	gLogin := r.Group("/api/v1/monitor/logininfor")
	gLogin.Use(mw)
	{
		gLogin.GET("/list", handler.ListLoginLog)
		gLogin.DELETE("/:id", handler.DeleteLoginLog)
		gLogin.DELETE("/clean", handler.CleanLoginLog)
	}

	gOper := r.Group("/api/v1/monitor/operlog")
	gOper.Use(mw)
	{
		gOper.GET("/list", handler.ListOperLog)
		gOper.DELETE("/:id", handler.DeleteOperLog)
		gOper.DELETE("/clean", handler.CleanOperLog)
	}
}

// --- Login Log ---

func (h *LogHandler) ListLoginLog(c *gin.Context) {
	list, err := h.logUsecase.GetLoginLogList(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, list)
}

func (h *LogHandler) DeleteLoginLog(c *gin.Context) {
	id := c.Param("id")
	if err := h.logUsecase.DeleteLoginLog(c.Request.Context(), id); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *LogHandler) CleanLoginLog(c *gin.Context) {
	if err := h.logUsecase.CleanLoginLog(c.Request.Context()); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}

// --- Oper Log ---

func (h *LogHandler) ListOperLog(c *gin.Context) {
	list, err := h.logUsecase.GetOperLogList(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, list)
}

func (h *LogHandler) DeleteOperLog(c *gin.Context) {
	id := c.Param("id")
	if err := h.logUsecase.DeleteOperLog(c.Request.Context(), id); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *LogHandler) CleanOperLog(c *gin.Context) {
	if err := h.logUsecase.CleanOperLog(c.Request.Context()); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}
