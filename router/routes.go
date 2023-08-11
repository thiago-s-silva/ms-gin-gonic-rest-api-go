package router

import (
	"ms-gin-gonic-rest-api-go/docs"
	"ms-gin-gonic-rest-api-go/handler"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(r *gin.Engine) {
	// init Handler
	handler.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := r.Group(basePath)
	{
		// openning domain
		v1.GET("/openning", handler.GetOneOpenningHandler)
		v1.POST("/openning", handler.CreateOpenningHandler)
		v1.DELETE("/openning", handler.DeleteOpenningHandler)
		v1.PUT("/openning", handler.UpdateOpenningHandler)
		v1.GET("/opennings", handler.GetAllOpenningsHandler)
	}

	// init Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
