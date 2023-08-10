package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleResponse(ctx *gin.Context, code int, message string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message": message,
	})
}

func handleResponseWithData(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message": message,
		"data":    data,
	})
}

func OnBadRequest(ctx *gin.Context, message string) {
	logger.Errorf("validation error: %v", message)
	handleResponse(ctx, http.StatusBadRequest, message)
}

func OnCreated(ctx *gin.Context, message string, data interface{}) {
	handleResponseWithData(ctx, http.StatusCreated, message, data)
}

func OnError(ctx *gin.Context, message string) {
	logger.Errorf("error when creating new openning: %v", message)
	handleResponse(ctx, http.StatusInternalServerError, message)
}
