package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/luankkobs/go-crud/src/configuration/logger"
	"github.com/luankkobs/go-crud/src/controller/routes"
	"log"
)

func main() {
	logger.Info("Starting application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
