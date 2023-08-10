package router

import (
	"ms-gin-gonic-rest-api-go/handler"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	// init Handler
	handler.InitializeHandler()

	v1 := r.Group("/api/v1")
	{
		// openning domain
		v1.GET("/openning", handler.GetOneOpenningHandler)
		v1.POST("/openning", handler.CreateOpenningHandler)
		v1.DELETE("/openning", handler.DeleteOpenningHandler)
		v1.PUT("/openning", handler.UpdateOpenningHandler)
		v1.GET("/opennings", handler.GetAllOpenningsHandler)
	}
}
