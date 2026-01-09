package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	RestCode string      `json:"restCode"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		RestCode: "200",
		Message:  "success",
		Data:     data,
	})
}

func Fail(c *gin.Context, httpCode int, restCode string, message string) {
	c.JSON(httpCode, Response{
		RestCode: restCode,
		Message:  message,
		Data:     nil,
	})
}

func Error(c *gin.Context, httpCode int, err error) {
	Fail(c, httpCode, "500", err.Error())
}

type PageResult struct {
	List  interface{} `json:"rows"` // Consistent with many table plugins, or "list"
	Total int64       `json:"total"`
	Page  int         `json:"page"`
	Size  int         `json:"size"`
}

func PageSuccess(c *gin.Context, list interface{}, total int64, page, size int) {
	Success(c, PageResult{
		List:  list,
		Total: total,
		Page:  page,
		Size:  size,
	})
}
