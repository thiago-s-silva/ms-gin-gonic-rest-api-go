package handler

import (
	"ms-gin-gonic-rest-api-go/schemas"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Create Openning
// @Description Create a new job openning
// @Tags Opennings
// @Accept json
// @Produce json
// @Param request body CreateOpenningRequest true "Request body"
// @Success 200 {object} CreateOpenningResponseType
// @Failure 400 {object} ErrorResponseType
// @Failure 500 {object} ErrorResponseType
// @Router /openning [post]
func CreateOpenningHandler(ctx *gin.Context) {
	var payload CreateOpenningRequest

	ctx.BindJSON(&payload)

	if err := payload.Validate(); err != nil {
		OnBadRequest(ctx, err.Error())
		return
	}

	openning := schemas.Openning{
		Role:     payload.Role,
		Company:  payload.Company,
		Location: payload.Location,
		Remote:   *payload.Remote,
		Link:     payload.Link,
		Salary:   payload.Salary,
	}

	dbResult := db.Create(&openning)

	if dbResult.Error != nil {
		OnError(ctx, "error when trying to create a new openning")
		return
	}

	OnCreated(ctx, "openning created", openning)
}
