package handler

import (
	"ms-gin-gonic-rest-api-go/schemas"
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
	logger.Debugf("created new openning: %v", data)
	handleResponseWithData(ctx, http.StatusCreated, message, data)
}

func OnError(ctx *gin.Context, message string) {
	logger.Errorf("error when creating new openning: %v", message)
	handleResponse(ctx, http.StatusInternalServerError, message)
}

func OnNotFound(ctx *gin.Context, message string) {
	logger.Debug(message)
	handleResponse(ctx, http.StatusNotFound, message)
}

func OnSuccess(ctx *gin.Context, message string, data interface{}) {
	logger.Debug(message)
	handleResponseWithData(ctx, http.StatusOK, message, data)
}

type ErrorResponseType struct {
	Message string `json:"message"`
}

type CreateOpenningResponseType struct {
	Message string                   `json:"message"`
	Data    schemas.OpenningResponse `json:"data"`
}

type UpdateOpenningResponseType struct {
	Message string                   `json:"message"`
	Data    schemas.OpenningResponse `json:"data"`
}

type DeleteOpenningResponseType struct {
	Message string                   `json:"message"`
	Data    schemas.OpenningResponse `json:"data"`
}

type GetOneOpenningResponseType struct {
	Message string                   `json:"message"`
	Data    schemas.OpenningResponse `json:"data"`
}

type GetAllOpenningsResponseType struct {
	Message string                     `json:"message"`
	Data    []schemas.OpenningResponse `json:"data"`
}
