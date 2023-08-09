package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllOpenningsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GetAll opennings",
	})
}
