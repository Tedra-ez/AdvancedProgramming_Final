package api

import (
	"github.com/Tedra-ez/AdvancedProgramming_Final/handlers"
	"github.com/gin-gonic/gin"
)

func SetUpRouters(r *gin.Engine, orderHandler *handlers.OrderHandler, authHandler *handlers.AuthHandler) {
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
}
