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

func CreateOpenningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Create openning",
	})
}

func DeleteOpenningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete openning",
	})
}

func UpdateOpenningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update openning",
	})
}

func GetAllOpenningsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GetAll opennings",
	})
}
