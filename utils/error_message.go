package utils

import "github.com/gin-gonic/gin"

func HandleError(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{
		"message": err.Error(),
	})
}
