package v1

import (
	"net/http"

	"admin-pro/internal/domain/entity"
	"admin-pro/internal/usecase"
	"admin-pro/pkg/response"
	"github.com/gin-gonic/gin"
)

type NoticeHandler struct {
	noticeUsecase usecase.NoticeUsecase
}

func NewNoticeHandler(r *gin.Engine, uc usecase.NoticeUsecase, authMw gin.HandlerFunc, permMw func(string) gin.HandlerFunc) {
	handler := &NoticeHandler{noticeUsecase: uc}
	g := r.Group("/api/v1/system/notice")
	g.Use(authMw)
	{
		g.GET("/list", permMw("system:notice:list"), handler.List)
		g.GET("/:id", permMw("system:notice:query"), handler.Get)
		g.POST("", permMw("system:notice:add"), handler.Add)
		g.PUT("", permMw("system:notice:edit"), handler.Update)
		g.DELETE("/:id", permMw("system:notice:remove"), handler.Delete)
	}
}

// @Summary 获取通知公告列表
// @Tags system/notice
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/v1/system/notice/list [get]
func (h *NoticeHandler) List(c *gin.Context) {
	list, err := h.noticeUsecase.GetNoticeList(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, list)
}

// @Summary 获取通知公告详情
// @Tags system/notice
// @Param id path string true "公告ID"
// @Success 200 {object} response.Response
// @Router /api/v1/system/notice/{id} [get]
func (h *NoticeHandler) Get(c *gin.Context) {
	id := c.Param("id")
	res, err := h.noticeUsecase.GetNotice(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, res)
}

// @Summary 新增通知公告
// @Tags system/notice
// @Accept json
// @Produce json
// @Param notice body entity.Notice true "公告信息"
// @Success 200 {object} response.Response
// @Router /api/v1/system/notice [post]
func (h *NoticeHandler) Add(c *gin.Context) {
	var notice entity.Notice
	if err := c.ShouldBindJSON(&notice); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", "invalid params")
		return
	}
	notice.CreatedBy = c.GetString("userID")
	if err := h.noticeUsecase.CreateNotice(c.Request.Context(), &notice); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}

// @Summary 更新通知公告
// @Tags system/notice
// @Accept json
// @Produce json
// @Param notice body entity.Notice true "公告信息"
// @Success 200 {object} response.Response
// @Router /api/v1/system/notice [put]
func (h *NoticeHandler) Update(c *gin.Context) {
	var notice entity.Notice
	if err := c.ShouldBindJSON(&notice); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", "invalid params")
		return
	}
	notice.UpdatedBy = c.GetString("userID")
	if err := h.noticeUsecase.UpdateNotice(c.Request.Context(), &notice); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}

// @Summary 删除通知公告
// @Tags system/notice
// @Param id path string true "公告ID"
// @Success 200 {object} response.Response
// @Router /api/v1/system/notice/{id} [delete]
func (h *NoticeHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.noticeUsecase.DeleteNotice(c.Request.Context(), id); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}
