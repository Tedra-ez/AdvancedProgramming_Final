package main

import (
	"context"
	"log"
	"time"

	"github.com/Tedra-ez/AdvancedProgramming_Final/handlers"
	"github.com/Tedra-ez/AdvancedProgramming_Final/internal/api"
	"github.com/Tedra-ez/AdvancedProgramming_Final/internal/config"
	"github.com/Tedra-ez/AdvancedProgramming_Final/internal/db"
	"github.com/Tedra-ez/AdvancedProgramming_Final/repository"
	"github.com/Tedra-ez/AdvancedProgramming_Final/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("env not loaded")
	}
	cfg := config.Load()
	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "pong"})
	})

	var mongoClient *db.MongoDBClient
	if cfg.MongoURI != "" {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		client, err := db.NewMongoDBClient(ctx, cfg.MongoURI)
		cancel()
		if err != nil {
			log.Fatalf("MongoDB: %v", err)
		}
		mongoClient = client
		defer func() {
			if err := mongoClient.Close(context.Background()); err != nil {
				log.Println("MongoDB close:", err)
			}
		}()
	}

	orderStore := repository.NewOrderStore(mongoClient)
	orderService := services.NewOrderService(orderStore)
	orderHandler := handlers.NewOrderHandler(orderService)

	api.SetUpRouters(server, orderHandler)
	addr := ":" + cfg.Port
	server.Run(addr)
}
