package handler

import (
	"fmt"
	"ms-gin-gonic-rest-api-go/schemas"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Delete Openning
// @Description Delete an openning
// @Tags Opennings
// @Produce json
// @Param id query string false "openning id"
// @Success 200 {object} DeleteOpenningResponseType
// @Failure 400 {object} ErrorResponseType
// @Failure 404 {object} ErrorResponseType
// @Failure 500 {object} ErrorResponseType
// @Router /openning [delete]
func DeleteOpenningHandler(ctx *gin.Context) {
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

	if err := db.Delete(&openning).Error; err != nil {
		message := fmt.Sprintf("failed to delete openning with id: %s", id)
		OnError(ctx, message)
		return
	}

	OnSuccess(ctx, "successfuly deleted openning", openning)
}
