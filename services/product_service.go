package services

import (
	"context"
	"errors"
	"time"

	"github.com/Tedra-ez/AdvancedProgramming_Final/models"
	"github.com/Tedra-ez/AdvancedProgramming_Final/repository"
)

type ProductService struct {
	repo repository.ProductStore
}

func NewProductService(repo repository.ProductStore) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) List(ctx context.Context) ([]*models.Product, error) {
	return s.repo.FindAll(ctx)
}

func (s *ProductService) GetByID(ctx context.Context, id string) (*models.Product, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *ProductService) Create(ctx context.Context, req *models.CreateProductRequest) (*models.Product, error) {
	now := time.Now()
	p := &models.Product{
		Name:        req.Name,
		Description: req.Description,
		Category:    req.Category,
		Price:       req.Price,
		Sizes:       req.Sizes,
		Colors:      req.Colors,
		StockBySize: req.StockBySize,
		Images:      req.Images,
		IsActive:    true,
		CreatedAt:   now,
		UpdateAt:    now,
	}
	if p.Sizes == nil {
		p.Sizes = []string{}
	}
	if p.Colors == nil {
		p.Colors = []string{}
	}
	if p.StockBySize == nil {
		p.StockBySize = make(map[string]int)
	}
	if p.Images == nil {
		p.Images = []string{}
	}
	if p.Price <= 0 {
		return nil, errors.New("price must be greater than 0")
	}
	return s.repo.Insert(ctx, p)
}

func (s *ProductService) Update(ctx context.Context, id string, p *models.Product) error {
	p.UpdateAt = time.Now()
	return s.repo.Update(ctx, id, p)
}

func (s *ProductService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
