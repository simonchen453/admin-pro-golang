package v1

import (
	"net/http"

	"admin-pro/internal/domain/entity"
	"admin-pro/internal/usecase"
	"admin-pro/pkg/response"
	"github.com/gin-gonic/gin"
)

type DictHandler struct {
	dictUsecase usecase.DictUsecase
}

func NewDictHandler(r *gin.Engine, uc usecase.DictUsecase, authMw gin.HandlerFunc, permMw func(string) gin.HandlerFunc) {
	handler := &DictHandler{dictUsecase: uc}
	// 字典类型路由
	gType := r.Group("/api/v1/system/dict/type")
	gType.Use(authMw)
	{
		gType.GET("/list", permMw("system:dict:list"), handler.ListType)
		gType.GET("/:id", permMw("system:dict:query"), handler.GetType)
		gType.POST("", permMw("system:dict:add"), handler.AddType)
		gType.PUT("", permMw("system:dict:edit"), handler.UpdateType)
		gType.DELETE("/:id", permMw("system:dict:remove"), handler.DeleteType)
	}

	// 字典数据路由
	gData := r.Group("/api/v1/system/dict/data")
	gData.Use(authMw)
	{
		gData.GET("/list", permMw("system:dict:list"), handler.ListData)
		gData.GET("/type/:dictType", permMw("system:dict:query"), handler.GetDataByType)
		gData.GET("/:id", permMw("system:dict:query"), handler.GetData)
		gData.POST("", permMw("system:dict:add"), handler.AddData)
		gData.PUT("", permMw("system:dict:edit"), handler.UpdateData)
		gData.DELETE("/:id", permMw("system:dict:remove"), handler.DeleteData)
	}
}

// --- 字典类型处理器 ---

func (h *DictHandler) ListType(c *gin.Context) {
	list, err := h.dictUsecase.GetDictTypeList(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, list)
}

func (h *DictHandler) GetType(c *gin.Context) {
	id := c.Param("id")
	res, err := h.dictUsecase.GetDictType(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, res)
}

func (h *DictHandler) AddType(c *gin.Context) {
	var dt entity.DictType
	if err := c.ShouldBindJSON(&dt); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", "invalid params")
		return
	}
	dt.CreatedBy = c.GetString("userID")
	if err := h.dictUsecase.CreateDictType(c.Request.Context(), &dt); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *DictHandler) UpdateType(c *gin.Context) {
	var dt entity.DictType
	if err := c.ShouldBindJSON(&dt); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", "invalid params")
		return
	}
	dt.UpdatedBy = c.GetString("userID")
	if err := h.dictUsecase.UpdateDictType(c.Request.Context(), &dt); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *DictHandler) DeleteType(c *gin.Context) {
	id := c.Param("id")
	if err := h.dictUsecase.DeleteDictType(c.Request.Context(), id); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}

// --- 字典数据处理器 ---

func (h *DictHandler) ListData(c *gin.Context) {
	dictType := c.Query("dictType")
	list, err := h.dictUsecase.GetDictDataByType(c.Request.Context(), dictType)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, list)
}

func (h *DictHandler) GetDataByType(c *gin.Context) {
	dictType := c.Param("dictType")
	list, err := h.dictUsecase.GetDictDataByType(c.Request.Context(), dictType)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, list)
}

func (h *DictHandler) GetData(c *gin.Context) {
	id := c.Param("id")
	res, err := h.dictUsecase.GetDictData(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, res)
}

func (h *DictHandler) AddData(c *gin.Context) {
	var dd entity.DictData
	if err := c.ShouldBindJSON(&dd); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", "invalid params")
		return
	}
	dd.CreatedBy = c.GetString("userID")
	if err := h.dictUsecase.CreateDictData(c.Request.Context(), &dd); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *DictHandler) UpdateData(c *gin.Context) {
	var dd entity.DictData
	if err := c.ShouldBindJSON(&dd); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", "invalid params")
		return
	}
	dd.UpdatedBy = c.GetString("userID")
	if err := h.dictUsecase.UpdateDictData(c.Request.Context(), &dd); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *DictHandler) DeleteData(c *gin.Context) {
	id := c.Param("id")
	if err := h.dictUsecase.DeleteDictData(c.Request.Context(), id); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}
