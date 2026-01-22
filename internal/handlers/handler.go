package handlers

import "github.com/gin-gonic/gin"

func ProductGetAll(c *gin.Context) {
	c.IndentedJSON(201, gin.H{
		"starus": "okay",
	})
}
