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

func NewDictHandler(r *gin.Engine, uc usecase.DictUsecase, mw gin.HandlerFunc) {
	handler := &DictHandler{
		dictUsecase: uc,
	}
	// Dict Type Routes
	gType := r.Group("/api/v1/system/dict/type")
	gType.Use(mw)
	{
		gType.GET("/list", handler.ListType)
		gType.GET("/:id", handler.GetType)
		gType.POST("", handler.AddType)
		gType.PUT("", handler.UpdateType)
		gType.DELETE("/:id", handler.DeleteType)
	}

	// Dict Data Routes
	gData := r.Group("/api/v1/system/dict/data")
	gData.Use(mw)
	{
		gData.GET("/list", handler.ListData)
		gData.GET("/type/:dictType", handler.GetDataByType) // Common use: get options by type
		gData.GET("/:id", handler.GetData)
		gData.POST("", handler.AddData)
		gData.PUT("", handler.UpdateData)
		gData.DELETE("/:id", handler.DeleteData)
	}
}

// --- Type Handlers ---

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

// --- Data Handlers ---

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
