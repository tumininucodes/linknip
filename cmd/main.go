package main

import (
	"linknip/internal/data"
	"linknip/internal/helpers"

	"github.com/gin-gonic/gin"
)

func main() {
	
	server := gin.Default()

	server.POST("/shorten", func(ctx *gin.Context) {
		var linkRequest data.LinkRequest
		err := ctx.ShouldBindJSON(linkRequest)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
		}

		if len(linkRequest.CustomSlug) != 0 {
			linkId, err := helpers.Base62Decode(linkRequest.CustomSlug)
			if err != nil {
				ctx.JSON(400, gin.H{"error": err.Error()})
			}
			link := data.Link {
				Id: linkId,
				Url: linkRequest.Url,
			}
			
		}

		// helpers.
	})


	server.Run(":8080")
}