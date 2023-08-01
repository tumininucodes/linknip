package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	
	server := gin.Default()

	server.POST("/shorten", func(ctx *gin.Context) {
		
		ctx.ShouldBindJSON()
	})


	server.Run(":8080")
}