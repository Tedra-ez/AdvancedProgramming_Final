package api

import (
	"github.com/Tedra-ez/AdvancedProgramming_Final/handlers"
	"github.com/gin-gonic/gin"
)

func SetUpRouters(r *gin.Engine, h handlers.ProductHandler) {
	api := r.Group("/api")
	{
		api.GET("/product", h.GetProducts)
		api.GET("/product/:id", h.GetProductByID)
		api.POST("/product", h.CreateProduct)
		api.PUT("/product/:id", h.UpdateProduct)
		api.DELETE("product/:id", h.DeleteProduct)

	}
}
