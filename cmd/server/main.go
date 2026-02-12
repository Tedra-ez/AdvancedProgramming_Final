package main

import (
	"context"
	"log"
	"time"

	"github.com/Tedra-ez/AdvancedProgramming_Final/internal/api"
	"github.com/Tedra-ez/AdvancedProgramming_Final/internal/config"
	"github.com/Tedra-ez/AdvancedProgramming_Final/internal/db"
	"github.com/Tedra-ez/AdvancedProgramming_Final/internal/handlers"
	"github.com/Tedra-ez/AdvancedProgramming_Final/internal/repository"
	"github.com/Tedra-ez/AdvancedProgramming_Final/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("env not loaded")
	}

	cfg := config.Load()

	server := gin.Default()
	server.Static("/static", "static")
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "pong"})
	})
	if cfg.MongoURI == "" {
		log.Fatalf("error when connecting to mongo, please specify MONGO_URI in .env")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	mongoClient, err := db.NewMongoDBClient(ctx, cfg.MongoURI)
	cancel()
	if err != nil {
		log.Fatalf("MongoDB: %v", err)
	}

	defer func() {
		if err := mongoClient.Close(context.Background()); err != nil {
			log.Println("MongoDB close:", err)
		}
	}()

	productCol := mongoClient.Collection("products")
	productRepo := repository.NewProductRepositoryMongo(productCol)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	userCol := mongoClient.Collection("users")
	userRepo := repository.NewUserRepository(userCol)
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	orderItemCol := mongoClient.Collection("order_items")
	orderCol := mongoClient.Collection("orders")
	orderItemRepo := repository.NewOrderItemRepositoryMongo(orderItemCol)
	orderRepo := repository.NewOrderRepositoryMongo(orderCol, orderItemRepo)
	orderService := services.NewOrderService(orderRepo, productRepo, userRepo)
	orderHandler := handlers.NewOrderHandler(orderService)

	analyticsService := services.NewAnalyticsService(orderRepo, productRepo, userRepo)
	analyticsHandler := handlers.NewAnalyticsHandler(analyticsService)

	pageHandler, err := handlers.NewPageHandler(productService, orderService, authService, analyticsService, "templates")
	if err != nil {
		log.Fatalf("templates: %v", err)
	}

	api.SetUpRouters(server, orderHandler, productHandler, authHandler, pageHandler, analyticsHandler, authService)

	addr := ":" + cfg.Port
	if err := server.Run(addr); err != nil {
		log.Fatal(err)
	}
}
