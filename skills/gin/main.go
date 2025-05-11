package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()

	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(200, "Gin Hello World!")
	})

	if err := server.Run(":8080"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
