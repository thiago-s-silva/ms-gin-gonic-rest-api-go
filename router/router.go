package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// initialize Router
	r := gin.Default()

	// initialize Routes
	initializeRoutes(r)

	// initialize Server
	r.Run("localhost:8080")
}
