package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// initialize Router
	r := gin.Default()

	// initialize Routes
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "its working"})
	})

	// initialize Server
	r.Run("localhost:8080")
}
