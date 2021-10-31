package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maiingonly/web-api/handlers"
)

func main() {
	r := gin.Default()
	r.GET("/ping", handlers.Pong)
	r.GET("/signup/:age", handlers.SignupHandler)
	r.GET("/query", handlers.QueryHandler)
	r.POST("/top-up", handlers.TopupHandler)

	r.Run(":8888")
}
