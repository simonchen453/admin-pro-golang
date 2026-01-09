package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"admin-pro/internal/domain/entity"
	"admin-pro/internal/usecase"
	"admin-pro/pkg/response"
)

type JobHandler struct {
	jobUsecase usecase.JobUsecase
}

func NewJobHandler(r *gin.Engine, uc usecase.JobUsecase, mw gin.HandlerFunc) {
	handler := &JobHandler{
		jobUsecase: uc,
	}
	
	g := r.Group("/api/v1/job")
	g.Use(mw)
	{
		g.GET("/list", handler.List)
		g.GET("/:id", handler.Get)
		g.POST("", handler.Add)
		g.PUT("", handler.Update)
		g.DELETE("/:id", handler.Delete)
	}

	gLog := r.Group("/api/v1/job/log")
	gLog.Use(mw)
	{
		gLog.GET("/list", handler.ListLog)
		gLog.GET("/:id", handler.GetLog)
	}
}

// @Summary Get job list
// @Tags job
// @Success 200 {object} response.Response
// @Router /api/v1/job/list [get]
func (h *JobHandler) List(c *gin.Context) {
	list, err := h.jobUsecase.GetJobList(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, list)
}

// @Summary Get job info
// @Tags job
// @Param id path string true "Job ID"
// @Success 200 {object} response.Response
// @Router /api/v1/job/{id} [get]
func (h *JobHandler) Get(c *gin.Context) {
	id := c.Param("id")
	res, err := h.jobUsecase.GetJob(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, res)
}

// @Summary Add job
// @Tags job
// @Accept json
// @Produce json
// @Param job body entity.Job true "Job Info"
// @Success 200 {object} response.Response
// @Router /api/v1/job [post]
func (h *JobHandler) Add(c *gin.Context) {
	var job entity.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", "invalid params")
		return
	}
	if err := h.jobUsecase.CreateJob(c.Request.Context(), &job); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}

// @Summary Update job
// @Tags job
// @Accept json
// @Produce json
// @Param job body entity.Job true "Job Info"
// @Success 200 {object} response.Response
// @Router /api/v1/job [put]
func (h *JobHandler) Update(c *gin.Context) {
	var job entity.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", "invalid params")
		return
	}
	if err := h.jobUsecase.UpdateJob(c.Request.Context(), &job); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}

// @Summary Delete job
// @Tags job
// @Param id path string true "Job ID"
// @Success 200 {object} response.Response
// @Router /api/v1/job/{id} [delete]
func (h *JobHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.jobUsecase.DeleteJob(c.Request.Context(), id); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}

// @Summary Get job log list
// @Tags job
// @Success 200 {object} response.Response
// @Router /api/v1/job/log/list [get]
func (h *JobHandler) ListLog(c *gin.Context) {
	list, err := h.jobUsecase.GetJobLogList(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, list)
}

// @Summary Get job log info
// @Tags job
// @Param id path string true "Job Log ID"
// @Success 200 {object} response.Response
// @Router /api/v1/job/log/{id} [get]
func (h *JobHandler) GetLog(c *gin.Context) {
	id := c.Param("id")
	res, err := h.jobUsecase.GetJobLog(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, res)
}
