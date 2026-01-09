package v1

import (
	"net/http"

	"admin-pro/internal/domain/entity"
	"admin-pro/internal/usecase"
	"admin-pro/pkg/response"
	"admin-pro/pkg/validator"
	"github.com/gin-gonic/gin"
)

type MenuHandler struct {
	menuUsecase usecase.MenuUsecase
}

func NewMenuHandler(r *gin.Engine, uc usecase.MenuUsecase, mw gin.HandlerFunc) {
	handler := &MenuHandler{
		menuUsecase: uc,
	}
	g := r.Group("/api/v1/system/menu")
	g.Use(mw)
	{
		g.GET("/list", handler.List)
		g.GET("/:id", handler.Get)
		g.POST("", handler.Add)
		g.PUT("", handler.Update)
		g.DELETE("/:id", handler.Delete)
	}
}

// @Summary Get menu list
// @Tags system/menu
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/v1/system/menu/list [get]
func (h *MenuHandler) List(c *gin.Context) {
	list, err := h.menuUsecase.GetMenuList(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, list)
}

// @Summary Get menu info
// @Tags system/menu
// @Param id path string true "Menu ID"
// @Success 200 {object} response.Response
// @Router /api/v1/system/menu/{id} [get]
func (h *MenuHandler) Get(c *gin.Context) {
	id := c.Param("id")
	menu, err := h.menuUsecase.GetMenu(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, menu)
}

// @Summary Add menu
// @Tags system/menu
// @Accept json
// @Produce json
// @Param menu body entity.Menu true "Menu Info"
// @Success 200 {object} response.Response
// @Router /api/v1/system/menu [post]
func (h *MenuHandler) Add(c *gin.Context) {
	var menu entity.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", "invalid params")
		return
	}

	// Validate input
	if err := validator.Validate(&menu); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", err.Error())
		return
	}

	if err := h.menuUsecase.CreateMenu(c.Request.Context(), &menu); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}

// @Summary Update menu
// @Tags system/menu
// @Accept json
// @Produce json
// @Param menu body entity.Menu true "Menu Info"
// @Success 200 {object} response.Response
// @Router /api/v1/system/menu [put]
func (h *MenuHandler) Update(c *gin.Context) {
	var menu entity.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", "invalid params")
		return
	}

	// Validate input
	if err := validator.Validate(&menu); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", err.Error())
		return
	}

	if err := h.menuUsecase.UpdateMenu(c.Request.Context(), &menu); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}

// @Summary Delete menu
// @Tags system/menu
// @Param id path string true "Menu ID"
// @Success 200 {object} response.Response
// @Router /api/v1/system/menu/{id} [delete]
func (h *MenuHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.menuUsecase.DeleteMenu(c.Request.Context(), id); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}
