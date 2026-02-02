package api

import (
	"github.com/Tedra-ez/AdvancedProgramming_Final/handlers"
	"github.com/gin-gonic/gin"
)

func SetUpRouters(r *gin.Engine, orderHandler *handlers.OrderHandler, productHandler handlers.ProductHandler, authHandler *handlers.AuthHandler) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	orders := r.Group("/orders")
	{
		orders.GET("/users/:userId/orders", orderHandler.ListOrdersByUser)
		orders.POST("", orderHandler.CreateOrder)
		orders.GET("/:id", orderHandler.GetOrderStatus)
		orders.PATCH("/:id/status", orderHandler.UpdateOrderStatus)
	}
	api := r.Group("/api")
	{
		api.GET("/product", productHandler.GetProducts)
		api.GET("/product/:id", productHandler.GetProductByID)
		api.POST("/product", productHandler.CreateProduct)
		api.PUT("/product/:id", productHandler.UpdateProduct)
		api.DELETE("product/:id", productHandler.DeleteProduct)
	}
}
