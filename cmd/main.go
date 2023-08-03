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

	server.GET("/:slug", func(ctx *gin.Context) {
		slug := ctx.Param("slug")
		id, err := helpers.Base62Decode(slug)
		if err != nil {
			ctx.JSON(400, gin.H {"error": err.Error()})
		}

		resolvedLink := data.GetLink(db, id)
		if len(resolvedLink.Url) == 0 {
			ctx.JSON(404, gin.H {"error": "no record found in database"})
		} else {
			ctx.Redirect(302, resolvedLink.Url)
			// ctx.JSON(200, gin.H {"shortenedUrl": resolvedLink.Url})
		}
		
	})


	server.Run(":8080")
}