package main

import (
	"fmt"
	"linknip/internal/data"
	"github.com/gin-gonic/gin"
)

func main() {
	
	server := gin.Default()

	db := data.OpenDB()
	fmt.Println(db.Stats().OpenConnections)

	server.Run(":8080")
}