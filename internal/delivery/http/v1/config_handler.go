package v1

import (
	"net/http"

	"admin-pro/internal/domain/entity"
	"admin-pro/internal/usecase"
	"admin-pro/pkg/response"
	"github.com/gin-gonic/gin"
)

type ConfigHandler struct {
	configUsecase usecase.ConfigUsecase
}

func NewConfigHandler(r *gin.Engine, uc usecase.ConfigUsecase, mw gin.HandlerFunc) {
	handler := &ConfigHandler{
		configUsecase: uc,
	}
	g := r.Group("/api/v1/system/config")
	g.Use(mw)
	{
		g.GET("/list", handler.List)
		g.GET("/:id", handler.Get)
		g.GET("/key/:key", handler.GetByKey) // Get by key
		g.POST("", handler.Add)
		g.PUT("", handler.Update)
		g.DELETE("/:id", handler.Delete)
	}
}

// @Summary Get config list
// @Tags system/config
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/v1/system/config/list [get]
func (h *ConfigHandler) List(c *gin.Context) {
	list, err := h.configUsecase.GetConfigList(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, list)
}

// @Summary Get config info
// @Tags system/config
// @Param id path string true "Config ID"
// @Success 200 {object} response.Response
// @Router /api/v1/system/config/{id} [get]
func (h *ConfigHandler) Get(c *gin.Context) {
	id := c.Param("id")
	res, err := h.configUsecase.GetConfig(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, res)
}

// @Summary Get config by key
// @Tags system/config
// @Param key path string true "Config Key"
// @Success 200 {object} response.Response
// @Router /api/v1/system/config/key/{key} [get]
func (h *ConfigHandler) GetByKey(c *gin.Context) {
	key := c.Param("key")
	res, err := h.configUsecase.GetConfigByKey(c.Request.Context(), key)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, res)
}

// @Summary Add config
// @Tags system/config
// @Accept json
// @Produce json
// @Param config body entity.Config true "Config Info"
// @Success 200 {object} response.Response
// @Router /api/v1/system/config [post]
func (h *ConfigHandler) Add(c *gin.Context) {
	var config entity.Config
	if err := c.ShouldBindJSON(&config); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", "invalid params")
		return
	}
	config.CreatedBy = c.GetString("userID")
	if err := h.configUsecase.CreateConfig(c.Request.Context(), &config); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}

// @Summary Update config
// @Tags system/config
// @Accept json
// @Produce json
// @Param config body entity.Config true "Config Info"
// @Success 200 {object} response.Response
// @Router /api/v1/system/config [put]
func (h *ConfigHandler) Update(c *gin.Context) {
	var config entity.Config
	if err := c.ShouldBindJSON(&config); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", "invalid params")
		return
	}
	config.UpdatedBy = c.GetString("userID")
	if err := h.configUsecase.UpdateConfig(c.Request.Context(), &config); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}

// @Summary Delete config
// @Tags system/config
// @Param id path string true "Config ID"
// @Success 200 {object} response.Response
// @Router /api/v1/system/config/{id} [delete]
func (h *ConfigHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.configUsecase.DeleteConfig(c.Request.Context(), id); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}
