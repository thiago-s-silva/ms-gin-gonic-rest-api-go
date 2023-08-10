package handler

import (
	"ms-gin-gonic-rest-api-go/schemas"

	"github.com/gin-gonic/gin"
)

func GetAllOpenningsHandler(ctx *gin.Context) {
	var opennings []schemas.Openning

	if err := db.Find(&opennings).Error; err != nil {
		OnError(ctx, "error when trying to get all opennings from db")
		return
	}

	OnSuccess(ctx, "successfuly select query", opennings)
}
