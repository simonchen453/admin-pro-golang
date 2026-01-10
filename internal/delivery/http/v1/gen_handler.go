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

func NewGenHandler(r *gin.Engine, uc usecase.GenUsecase, authMw gin.HandlerFunc, permMw func(string) gin.HandlerFunc) {
	handler := &GenHandler{genUsecase: uc}
	g := r.Group("/api/v1/generator")
	g.Use(authMw)
	{
		g.GET("/list", permMw("tool:gen:list"), handler.List)
		g.GET("/batchGenCode", permMw("tool:gen:code"), handler.BatchGenCode)
	}
}

// @Summary 获取数据表列表
// @Tags generator
// @Param tableName query string false "表名"
// @Param pageNo query int false "页码"
// @Param pageSize query int false "每页数量"
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

// @Summary 批量生成代码
// @Tags generator
// @Param tables query string true "表名列表(逗号分隔)"
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
