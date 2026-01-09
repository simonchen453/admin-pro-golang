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

func NewNoticeHandler(r *gin.Engine, uc usecase.NoticeUsecase, mw gin.HandlerFunc) {
	handler := &NoticeHandler{
		noticeUsecase: uc,
	}
	g := r.Group("/api/v1/system/notice")
	g.Use(mw)
	{
		g.GET("/list", handler.List)
		g.GET("/:id", handler.Get)
		g.POST("", handler.Add)
		g.PUT("", handler.Update)
		g.DELETE("/:id", handler.Delete)
	}
}

// @Summary Get notice list
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

// @Summary Get notice info
// @Tags system/notice
// @Param id path string true "Notice ID"
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

// @Summary Add notice
// @Tags system/notice
// @Accept json
// @Produce json
// @Param notice body entity.Notice true "Notice Info"
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

// @Summary Update notice
// @Tags system/notice
// @Accept json
// @Produce json
// @Param notice body entity.Notice true "Notice Info"
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

// @Summary Delete notice
// @Tags system/notice
// @Param id path string true "Notice ID"
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
