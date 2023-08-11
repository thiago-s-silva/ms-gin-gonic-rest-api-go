package handler

import (
	"fmt"
	"ms-gin-gonic-rest-api-go/schemas"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Get Openning
// @Description Get one openning by id
// @Tags Opennings
// @Produce json
// @Param id query string false "openning id"
// @Success 200 {object} GetOneOpenningResponseType
// @Failure 400 {object} ErrorResponseType
// @Failure 404 {object} ErrorResponseType
// @Failure 500 {object} ErrorResponseType
// @Router /openning [get]
func GetOneOpenningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		OnBadRequest(ctx, "invalid query parameter ID")
		return
	}

	var openning schemas.Openning

	if err := db.First(&openning, id).Error; err != nil {
		message := fmt.Sprintf("openning not found for id: %s", id)

		OnNotFound(ctx, message)
		return
	}

	OnSuccess(ctx, "message successfuly select query", openning)
}
