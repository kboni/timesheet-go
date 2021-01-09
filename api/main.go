package main

import (
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func main() {
	Router = gin.Default()
	api := Router.Group("/api")
	{
		api.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "test successful",
			})
		})
		api.GET("/test2", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "another test",
			})
		})
	}
	Router.Run(":5000")
}
