package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DefaultTask as an example
func DefaultTask(ctx *gin.Context) {
	// TODO: finish the default task
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
