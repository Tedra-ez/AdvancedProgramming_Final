package main

import (
	"github.com/Tedra-ez/AdvancedProgramming_Final/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Setup(r)
	r.Run(":8080") //Test3
}
