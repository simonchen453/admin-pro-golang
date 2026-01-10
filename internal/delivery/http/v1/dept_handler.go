package v1

import (
	"net/http"

	"admin-pro/internal/domain/entity"
	"admin-pro/internal/usecase"
	"admin-pro/pkg/response"
	"admin-pro/pkg/validator"
	"github.com/gin-gonic/gin"
)

type DeptHandler struct {
	deptUsecase usecase.DeptUsecase
}

func NewDeptHandler(r *gin.Engine, uc usecase.DeptUsecase, authMw gin.HandlerFunc, permMw func(string) gin.HandlerFunc) {
	handler := &DeptHandler{
		deptUsecase: uc,
	}
	g := r.Group("/api/v1/system/dept")
	g.Use(authMw) // JWT 认证（会加载权限）
	{
		g.GET("/list", permMw("system:dept:list"), handler.List)
		g.GET("/:id", permMw("system:dept:query"), handler.Get)
		g.POST("", permMw("system:dept:add"), handler.Add)
		g.PUT("", permMw("system:dept:edit"), handler.Update)
		g.DELETE("/:id", permMw("system:dept:remove"), handler.Delete)
	}
}

// @Summary 获取部门列表
// @Tags system/dept
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/v1/system/dept/list [get]
func (h *DeptHandler) List(c *gin.Context) {
	list, err := h.deptUsecase.GetDeptList(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, list)
}

// @Summary 获取部门详情
// @Tags system/dept
// @Param id path string true "部门ID"
// @Success 200 {object} response.Response
// @Router /api/v1/system/dept/{id} [get]
func (h *DeptHandler) Get(c *gin.Context) {
	id := c.Param("id")
	dept, err := h.deptUsecase.GetDept(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, dept)
}

// @Summary 新增部门
// @Tags system/dept
// @Accept json
// @Produce json
// @Param dept body entity.Dept true "部门信息"
// @Success 200 {object} response.Response
// @Router /api/v1/system/dept [post]
func (h *DeptHandler) Add(c *gin.Context) {
	var dept entity.Dept
	if err := c.ShouldBindJSON(&dept); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", "invalid params")
		return
	}

	// 校验输入
	if err := validator.Validate(&dept); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", err.Error())
		return
	}

	// 从上下文获取基础信息
	dept.CreatedBy = c.GetString("userID")

	if err := h.deptUsecase.CreateDept(c.Request.Context(), &dept); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}

// @Summary 更新部门
// @Tags system/dept
// @Accept json
// @Produce json
// @Param dept body entity.Dept true "部门信息"
// @Success 200 {object} response.Response
// @Router /api/v1/system/dept [put]
func (h *DeptHandler) Update(c *gin.Context) {
	var dept entity.Dept
	if err := c.ShouldBindJSON(&dept); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", "invalid params")
		return
	}

	// Validate input
	if err := validator.Validate(&dept); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", err.Error())
		return
	}

	dept.UpdatedBy = c.GetString("userID")

	if err := h.deptUsecase.UpdateDept(c.Request.Context(), &dept); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}

// @Summary 删除部门
// @Tags system/dept
// @Param id path string true "部门ID"
// @Success 200 {object} response.Response
// @Router /api/v1/system/dept/{id} [delete]
func (h *DeptHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.deptUsecase.DeleteDept(c.Request.Context(), id); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}
