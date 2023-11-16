package main

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadputraage/goschool-api/config"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
}

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	r.Run()
}
