package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maiingonly/web-api/handlers"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")

	v1.GET("/ping", handlers.Pong)
	r.GET("/signup/:age", handlers.SignupHandler)
	r.GET("/query", handlers.QueryHandler)
	v1.POST("/top-up", handlers.TopupHandler)

	r.Run(":8888")
}
