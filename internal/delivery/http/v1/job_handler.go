package v1

import (
	"net/http"

	"admin-pro/internal/domain/entity"
	"admin-pro/internal/usecase"
	"admin-pro/pkg/response"
	"admin-pro/pkg/validator"
	"github.com/gin-gonic/gin"
)

type JobHandler struct {
	jobUsecase usecase.JobUsecase
}

func NewJobHandler(r *gin.Engine, uc usecase.JobUsecase, authMw gin.HandlerFunc, permMw func(string) gin.HandlerFunc) {
	handler := &JobHandler{jobUsecase: uc}

	g := r.Group("/api/v1/job")
	g.Use(authMw)
	{
		g.GET("/list", permMw("job:job:list"), handler.List)
		g.GET("/:id", permMw("job:job:query"), handler.Get)
		g.POST("", permMw("job:job:add"), handler.Add)
		g.PUT("", permMw("job:job:edit"), handler.Update)
		g.DELETE("/:id", permMw("job:job:remove"), handler.Delete)
	}

	gLog := r.Group("/api/v1/job/log")
	gLog.Use(authMw)
	{
		gLog.GET("/list", permMw("job:log:list"), handler.ListLog)
		gLog.GET("/:id", permMw("job:log:query"), handler.GetLog)
	}
}

// @Summary 获取定时任务列表
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

// @Summary 获取定时任务详情
// @Tags job
// @Param id path string true "任务ID"
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

// @Summary 新增定时任务
// @Tags job
// @Accept json
// @Produce json
// @Param job body entity.Job true "任务信息"
// @Success 200 {object} response.Response
// @Router /api/v1/job [post]
func (h *JobHandler) Add(c *gin.Context) {
	var job entity.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", "invalid params")
		return
	}

	// Validate input
	if err := validator.Validate(&job); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", err.Error())
		return
	}

	if err := h.jobUsecase.CreateJob(c.Request.Context(), &job); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}

// @Summary 更新定时任务
// @Tags job
// @Accept json
// @Produce json
// @Param job body entity.Job true "任务信息"
// @Success 200 {object} response.Response
// @Router /api/v1/job [put]
func (h *JobHandler) Update(c *gin.Context) {
	var job entity.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", "invalid params")
		return
	}

	// Validate input
	if err := validator.Validate(&job); err != nil {
		response.Fail(c, http.StatusBadRequest, "400", err.Error())
		return
	}

	if err := h.jobUsecase.UpdateJob(c.Request.Context(), &job); err != nil {
		response.Fail(c, http.StatusOK, "500", err.Error())
		return
	}
	response.Success(c, nil)
}

// @Summary 删除定时任务
// @Tags job
// @Param id path string true "任务ID"
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

// @Summary 获取任务日志列表
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

// @Summary 获取任务日志详情
// @Tags job
// @Param id path string true "日志ID"
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
