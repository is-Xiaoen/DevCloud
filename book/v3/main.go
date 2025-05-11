package main

import (
	"fmt"
	"os"

	"122.51.31.227/go-course/go18/book/v3/config"
	"122.51.31.227/go-course/go18/book/v3/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	handlers.Book.Registry(server)

	ac := config.C().Application
	// 启动服务
	if err := server.Run(fmt.Sprintf("%s:%d", ac.Host, ac.Port)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
