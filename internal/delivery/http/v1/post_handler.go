package v1

import (
	"net/http"

	"admin-pro/internal/domain/entity"
	"admin-pro/internal/usecase"
	"admin-pro/pkg/response"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postUsecase usecase.PostUsecase
}

func NewPostHandler(r *gin.Engine, uc usecase.PostUsecase, authMw gin.HandlerFunc, permMw func(string) gin.HandlerFunc) {
	handler := &PostHandler{postUsecase: uc}
	g := r.Group("/api/v1/system/post")
	g.Use(authMw)
	{
		g.GET("/list", permMw("system:post:list"), handler.List)
		g.GET("/:id", permMw("system:post:query"), handler.Get)
		g.POST("", permMw("system:post:add"), handler.Add)
		g.PUT("", permMw("system:post:edit"), handler.Update)
		g.DELETE("/:id", permMw("system:post:remove"), handler.Delete)
	}
}

// @Summary 获取岗位列表
// @Tags system/post
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/v1/system/post/list [get]
func (h *PostHandler) List(c *gin.Context) {
	list, err := h.postUsecase.GetPostList(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, list)
}

// @Summary 获取岗位详情
// @Tags system/post
// @Param id path string true "岗位ID"
// @Success 200 {object} response.Response
// @Router /api/v1/system/post/{id} [get]
func (h *PostHandler) Get(c *gin.Context) {
	id := c.Param("id")
	post, err := h.postUsecase.GetPost(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, post)
}

// @Summary 新增岗位
// @Tags system/post
// @Accept json
// @Produce json
// @Param post body entity.Post true "岗位信息"
// @Success 200 {object} response.Response
// @Router /api/v1/system/post [post]
func (h *PostHandler) Add(c *gin.Context) {
	var post entity.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", "invalid params")
		return
	}

	post.CreatedBy = c.GetString("userID")

	if err := h.postUsecase.CreatePost(c.Request.Context(), &post); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}

// @Summary 更新岗位
// @Tags system/post
// @Accept json
// @Produce json
// @Param post body entity.Post true "岗位信息"
// @Success 200 {object} response.Response
// @Router /api/v1/system/post [put]
func (h *PostHandler) Update(c *gin.Context) {
	var post entity.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", "invalid params")
		return
	}

	post.UpdatedBy = c.GetString("userID")

	if err := h.postUsecase.UpdatePost(c.Request.Context(), &post); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}

// @Summary 删除岗位
// @Tags system/post
// @Param id path string true "岗位ID"
// @Success 200 {object} response.Response
// @Router /api/v1/system/post/{id} [delete]
func (h *PostHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.postUsecase.DeletePost(c.Request.Context(), id); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}
