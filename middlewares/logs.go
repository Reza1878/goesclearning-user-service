package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Logger(router *gin.Engine) gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] | %d | %v | %s | %s %s | %s | %s\n",
			param.TimeStamp.Format("2006-01-02 15:04:05"),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method, param.Path,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}
