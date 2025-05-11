package main

import (
	"fmt"
	"os"

	"122.51.31.227/go-course/go18/book/v3/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	handlers.Book.Registry(server)

	if err := server.Run(":8080"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
