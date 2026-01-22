package router

import (
	"github.com/Tedra-ez/AdvancedProgramming_Final/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {

	api := r.Group("/api")
	{
		api.GET("/products", handlers.ProductGetAll)
	}
}
