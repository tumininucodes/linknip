package main

import (
	"fmt"
	"linknip/internal/data"
	"linknip/internal/helpers"

	"github.com/gin-gonic/gin"
)

func main() {
	
	server := gin.Default()

	db := data.OpenDB()

	server.POST("/shorten", func(ctx *gin.Context) {
		var linkRequest data.LinkRequest
		err := ctx.ShouldBindJSON(&linkRequest)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
		}

		if len(linkRequest.CustomSlug) != 0 {
			linkId, err := helpers.Base62Decode(linkRequest.CustomSlug)
			if err != nil {
				fmt.Println("reaached here")
				ctx.JSON(400, gin.H{"error": err.Error()})
			}
			link := data.Link {
				Id: linkId,
				Url: linkRequest.Url,
			}

			ctx.JSON(201, data.InsertLink(db, &link))
		}

	})

	server.GET("/{slug}", func(ctx *gin.Context) {
		slug := ctx.Param("slug")
		id, err := helpers.Base62Decode(slug)
		if err != nil {
			ctx.JSON(400, gin.H {"error": err.Error()})
		}

		ctx.JSON(200, data.GetLink(db, id))
	})


	server.Run(":8080")
}