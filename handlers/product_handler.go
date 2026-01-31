package handlers

import (
	"errors"

	"github.com/Tedra-ez/AdvancedProgramming_Final/models"
	"github.com/Tedra-ez/AdvancedProgramming_Final/services"
	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	GetProducts(ctx *gin.Context)
	GetProductByID(ctx *gin.Context)
	CreateProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
}

type productHandler struct {
	service services.ProductService
}

func New(productService services.ProductService) ProductHandler {
	return &productHandler{
		service: productService,
	}
}

func (ph *productHandler) GetProducts(ctx *gin.Context) {
	list, err := ph.service.List()
	if err != nil {
		if err.Error() == "slice is empty" {
			ctx.JSON(204, gin.H{
				"status": "slice is empty",
			})
			return
		}
	}

	ctx.JSON(200, list)
	return
}
func (ph *productHandler) GetProductByID(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := ph.service.GetByID(id)

	if err != nil {
		if err.Error() == "product not found" {
			ctx.JSON(404, gin.H{
				"error": "product not found",
			})
			return
		}
	}

	ctx.JSON(200, product)
}
func (ph *productHandler) CreateProduct(ctx *gin.Context) {
	var product models.Product

	ctx.BindJSON(&product)

	err := ph.service.Create(product)
	switch {
	case errors.Is(err, services.ErrInvalidPrice):
		ctx.JSON(400, gin.H{"error": err.Error()})
	case errors.Is(err, services.ErrProductExists):
		ctx.JSON(409, gin.H{"error": err.Error()})
	case err != nil:
		ctx.JSON(500, gin.H{"error": "internal error"})
	default:
		ctx.JSON(201, product)
	}

	ctx.JSON(201, gin.H{
		"status":  "created",
		"product": product,
	})
}
func (ph *productHandler) UpdateProduct(ctx *gin.Context) {
	var product models.Product

	id := ctx.Param("id")
	ctx.BindJSON(&product)

	err := ph.service.Update(id, product)
	if err != nil {
		if err.Error() == "Product doesn't exist" {
			ctx.JSON(404, gin.H{
				"error": "product not found",
			})
			return
		}
	}

	ctx.JSON(201, gin.H{
		"status":  "updated",
		"product": product,
	})
}
func (ph *productHandler) DeleteProduct(ctx *gin.Context) {

}
