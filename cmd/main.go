package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"linknip/internal/data"
	"linknip/internal/helpers"
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

		if !helpers.IsUrl(linkRequest.Url) {
			ctx.JSON(400, gin.H {"error": "invalid url"})
			return
		}

		var slug string

		if len(linkRequest.CustomSlug) != 0 {
			slug = linkRequest.CustomSlug
		} else {
			slug, err = helpers.GenerateRandomString(5)
			if err != nil {
				panic(err.Error())
			}
		}

		linkId, err := helpers.Base62Decode(slug)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
		}
		link := data.Link{
			Id:  fmt.Sprintf("%+v", linkId),
			Url: linkRequest.Url,
		}

		result, error := data.InsertLink(db, &link)
		if error != nil {
			ctx.JSON(400, gin.H{"error": "record already exists"})
		} else {
			ctx.JSON(201, result)
		}

	})

	server.GET("/:slug", func(ctx *gin.Context) {
		slug := ctx.Param("slug")
		id, err := helpers.Base62Decode(slug)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
		} else {
			resolvedLink := data.GetLink(db, id)
			if len(resolvedLink.Url) == 0 {
				ctx.JSON(404, gin.H{"error": "no record found in database"})
			} else {
				ctx.Redirect(302, resolvedLink.Url)
			}
		}
	})

	server.Run(":8080")
}
