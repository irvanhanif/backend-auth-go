package helper

import (
	"github.com/gin-gonic/gin"
)

func ERROR(ctx *gin.Context, code int, message interface{}) {
	ctx.AbortWithStatusJSON(code, gin.H{
		"success": false,
		"message": message,
	})
}

func SUCCESS(ctx *gin.Context, code int, data interface{})  {
	ctx.JSON(code, gin.H{
		"success": true,
		"data": data,
	})
}