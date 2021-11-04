package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/maiingonly/web-api/entities"
	"github.com/maiingonly/web-api/handlers"
	"github.com/maiingonly/web-api/util"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	//viper package read .env
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta", config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)

	db, err := gorm.Open(postgres.Open(viper.GetString(dsn)), &gorm.Config{})
	if err != nil {
		log.Println("Db connection error")
	}

	db.AutoMigrate(&entities.User_account{})

	r := gin.Default()
	v1 := r.Group("/v1")

	v1.GET("/ping", handlers.Pong)
	r.GET("/signup/:age", handlers.SignupHandler)
	r.GET("/query", handlers.QueryHandler)
	v1.POST("/top-up", handlers.TopupHandler)

	r.Run(config.ServerAddress)
}
