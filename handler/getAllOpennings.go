package handler

import (
	"ms-gin-gonic-rest-api-go/schemas"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Get All Opennings
// @Description Get all opennings
// @Tags Opennings
// @Produce json
// @Success 200 {object} GetAllOpenningsResponseType
// @Failure 400 {object} ErrorResponseType
// @Failure 500 {object} ErrorResponseType
// @Router /opennings [get]
func GetAllOpenningsHandler(ctx *gin.Context) {
	var opennings []schemas.Openning

	if err := db.Find(&opennings).Error; err != nil {
		OnError(ctx, "error when trying to get all opennings from db")
		return
	}

	OnSuccess(ctx, "successfuly select query", opennings)
}
