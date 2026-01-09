package v1

import (
	"net/http"

	"admin-pro/internal/domain/entity"
	"admin-pro/internal/usecase"
	"admin-pro/pkg/response"
	"admin-pro/pkg/validator"
	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	roleUsecase usecase.RoleUsecase
}

func NewRoleHandler(r *gin.Engine, uc usecase.RoleUsecase, mw gin.HandlerFunc) {
	handler := &RoleHandler{
		roleUsecase: uc,
	}
	g := r.Group("/api/v1/system/role")
	g.Use(mw)
	{
		g.GET("/list", handler.List)
		g.GET("/:id", handler.Get)
		g.POST("", handler.Add)
		g.PUT("", handler.Update)
		g.DELETE("/:id", handler.Delete)
	}
}

// @Summary Get role list
// @Tags system/role
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/v1/system/role/list [get]
func (h *RoleHandler) List(c *gin.Context) {
	list, err := h.roleUsecase.GetRoleList(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, list)
}

// @Summary Get role info
// @Tags system/role
// @Param id path string true "Role ID"
// @Success 200 {object} response.Response
// @Router /api/v1/system/role/{id} [get]
func (h *RoleHandler) Get(c *gin.Context) {
	id := c.Param("id")
	role, err := h.roleUsecase.GetRole(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, role)
}

// @Summary Add role
// @Tags system/role
// @Accept json
// @Produce json
// @Param role body entity.Role true "Role Info"
// @Success 200 {object} response.Response
// @Router /api/v1/system/role [post]
func (h *RoleHandler) Add(c *gin.Context) {
	var role entity.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", "invalid params")
		return
	}

	// Validate input
	if err := validator.Validate(&role); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", err.Error())
		return
	}

	if err := h.roleUsecase.CreateRole(c.Request.Context(), &role); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}

// @Summary Update role
// @Tags system/role
// @Accept json
// @Produce json
// @Param role body entity.Role true "Role Info"
// @Success 200 {object} response.Response
// @Router /api/v1/system/role [put]
func (h *RoleHandler) Update(c *gin.Context) {
	var role entity.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", "invalid params")
		return
	}

	// Validate input
	if err := validator.Validate(&role); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", err.Error())
		return
	}

	if err := h.roleUsecase.UpdateRole(c.Request.Context(), &role); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}

// @Summary Delete role
// @Tags system/role
// @Param id path string true "Role ID"
// @Success 200 {object} response.Response
// @Router /api/v1/system/role/{id} [delete]
func (h *RoleHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.roleUsecase.DeleteRole(c.Request.Context(), id); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}
