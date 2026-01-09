package v1

import (
	"net/http"
	"strconv"
	"strings"

	"admin-pro/internal/usecase"
	"admin-pro/pkg/response"
	"github.com/gin-gonic/gin"
)

type GenHandler struct {
	genUsecase usecase.GenUsecase
}

func NewGenHandler(r *gin.Engine, uc usecase.GenUsecase, mw gin.HandlerFunc) {
	handler := &GenHandler{
		genUsecase: uc,
	}
	g := r.Group("/api/v1/generator") // Use /api/v1/generator to match typical pattern, frontend might need adjust or proxy
	// Frontend uses /admin/generator -> /api/v1/generator via proxy
	g.Use(mw)
	{
		g.GET("/list", handler.List)
		g.GET("/batchGenCode", handler.BatchGenCode) // Download
		// g.GET("/genAll", handler.GenAll)
	}
}

// @Summary Get table list
// @Tags generator
// @Param tableName query string false "Table Name"
// @Param pageNo query int false "Page Number"
// @Param pageSize query int false "Page Size"
// @Success 200 {object} response.Response
// @Router /api/v1/generator/list [get]
func (h *GenHandler) List(c *gin.Context) {
	tableName := c.Query("tableName")
	pageNo, err := strconv.Atoi(c.DefaultQuery("pageNo", "1"))
	if err != nil {
		response.Fail(c, http.StatusBadRequest, "400", "invalid pageNo parameter")
		return
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if err != nil {
		response.Fail(c, http.StatusBadRequest, "400", "invalid pageSize parameter")
		return
	}

	list, count, err := h.genUsecase.GetTableList(c.Request.Context(), tableName, pageSize, pageNo)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.PageSuccess(c, list, count, pageNo, pageSize)
}

// @Summary Batch generate code
// @Tags generator
// @Param tables query string true "Table Names (comma separated)"
// @Success 200 {file} application/zip
// @Router /api/v1/generator/batchGenCode [get]
func (h *GenHandler) BatchGenCode(c *gin.Context) {
	tablesStr := c.Query("tables")
	if tablesStr == "" {
		response.Fail(c, http.StatusBadRequest, "400", "tables required")
		return
	}
	tables := strings.Split(tablesStr, ",")

	data, err := h.genUsecase.GenCode(c.Request.Context(), tables)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, "500", err.Error())
		return
	}

	c.Header("Content-Disposition", "attachment; filename=\"admin-pro-gen.zip\"")
	c.Data(http.StatusOK, "application/zip", data)
}
