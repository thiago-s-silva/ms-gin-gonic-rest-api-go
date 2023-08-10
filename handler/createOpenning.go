package handler

import (
	"ms-gin-gonic-rest-api-go/schemas"

	"github.com/gin-gonic/gin"
)

func CreateOpenningHandler(ctx *gin.Context) {
	payload := CreateOpenningRequest{}

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
