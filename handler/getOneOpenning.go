package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOneOpenningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get one openning",
	})
}
