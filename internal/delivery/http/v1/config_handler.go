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

func NewConfigHandler(r *gin.Engine, uc usecase.ConfigUsecase, authMw gin.HandlerFunc, permMw func(string) gin.HandlerFunc) {
	handler := &ConfigHandler{configUsecase: uc}
	g := r.Group("/api/v1/system/config")
	g.Use(authMw)
	{
		g.GET("/list", permMw("system:config:list"), handler.List)
		g.GET("/:id", permMw("system:config:query"), handler.Get)
		g.GET("/key/:key", permMw("system:config:query"), handler.GetByKey)
		g.POST("", permMw("system:config:add"), handler.Add)
		g.PUT("", permMw("system:config:edit"), handler.Update)
		g.DELETE("/:id", permMw("system:config:remove"), handler.Delete)
	}
}

// @Summary 获取参数配置列表
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

// @Summary 获取参数配置详情
// @Tags system/config
// @Param id path string true "配置ID"
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

// @Summary 根据Key获取配置
// @Tags system/config
// @Param key path string true "配置Key"
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

// @Summary 新增参数配置
// @Tags system/config
// @Accept json
// @Produce json
// @Param config body entity.Config true "配置信息"
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

// @Summary 更新参数配置
// @Tags system/config
// @Accept json
// @Produce json
// @Param config body entity.Config true "配置信息"
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

// @Summary 删除参数配置
// @Tags system/config
// @Param id path string true "配置ID"
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
