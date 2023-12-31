package handler

import (
	"fmt"
	"ms-gin-gonic-rest-api-go/schemas"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Update Openning
// @Description Update an openning
// @Tags Opennings
// @Accept json
// @Produce json
// @Param request body CreateOpenningRequest true "Request body"
// @Param id query string false "openning id"
// @Success 200 {object} UpdateOpenningResponseType
// @Failure 400 {object} ErrorResponseType
// @Failure 404 {object} ErrorResponseType
// @Failure 500 {object} ErrorResponseType
// @Router /openning [put]
func UpdateOpenningHandler(ctx *gin.Context) {
	var payload UpdateOpenningRequest

	ctx.BindJSON(&payload)

	if err := payload.Validate(); err != nil {
		OnBadRequest(ctx, err.Error())
		return
	}

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

	if payload.Role != "" {
		openning.Role = payload.Role
	}

	if payload.Company != "" {
		openning.Company = payload.Company
	}

	if payload.Location != "" {
		openning.Location = payload.Location
	}

	if payload.Remote != nil {
		openning.Remote = *payload.Remote
	}

	if payload.Link != "" {
		openning.Link = payload.Link
	}

	if payload.Salary > 0 {
		openning.Salary = payload.Salary
	}

	if err := db.Save(&openning).Error; err != nil {
		OnError(ctx, "error when tried to update openning")
		return
	}

	OnSuccess(ctx, "successfuly updated", openning)
}
