package response

import (
	"github.com/Reza1878/goesclearning/user-service/model"
	"github.com/gin-gonic/gin"
)

func JSON(ctx *gin.Context, statusCode int, message string, data interface{}) {
	ctx.JSON(statusCode, model.ResponseSuccess{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	})
}
