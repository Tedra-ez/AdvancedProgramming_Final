package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	server := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Fatal("env not loaded")
	}

	server.GET("/ping", func(c *gin.Context) {

		c.JSON(200, gin.H{"msg": "pong"})
	})

	server.Run(":8080")
}
