package main

import (
	"github.com/Tedra-ez/AdvancedProgramming_Final/handlers"
	"github.com/Tedra-ez/AdvancedProgramming_Final/internal/api"
	"github.com/Tedra-ez/AdvancedProgramming_Final/repository"
	"github.com/Tedra-ez/AdvancedProgramming_Final/services"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// make sure to create an interface to each service handler and repository.
	//product service not the final version
	db := repository.New()
	productService := services.New(db)
	productHandler := handlers.New(productService)

	api.SetUpRouters(server, productHandler)
	server.Run()
}
