package services

import (
	"errors"

	"github.com/Tedra-ez/AdvancedProgramming_Final/models"
	"github.com/Tedra-ez/AdvancedProgramming_Final/repository"
)

// also, for the purity of the code, it is better to pull the error in the var.
var (
	ErrInvalidPrice    = errors.New("invalid price")
	ErrProductExists   = errors.New("product already exists")
	ErrProductNotFound = errors.New("product not found")
)

type ProductService interface {
	List() ([]models.Product, error)
	GetByID(ID string) (models.Product, error)
	Create(product models.Product) error
	Update(id string, product models.Product) error
	Delete(id string) error
}

type productService struct {
	repository repository.ProductRepository
}

func New(r repository.ProductRepository) ProductService {
	return &productService{repository: r}
}

func (ps *productService) List() ([]models.Product, error) {
	return ps.repository.FindAll()
}
func (ps *productService) GetByID(ID string) (models.Product, error) {
	return ps.repository.FindByID(ID)
}
func (ps *productService) Create(product models.Product) error {
	if product.Price <= 0 {
		return ErrInvalidPrice
	}
	err := ps.repository.Insert(product)
	if err != nil {
		if errors.Is(err, repository.ErrAlreadyExists) {
			return ErrProductExists
		}
		return err // unknown error --> handler
	}
	return nil
}
func (ps *productService) Update(id string, product models.Product) error {
	if product.Price <= 0 {
		return ErrInvalidPrice
	}
	err := ps.repository.Update(id, product)
	if err != nil {
		if errors.Is(err, repository.ErrNotExist) {
			return ErrProductNotFound
		}
		return err // unknown error --> handler
	}

	return nil
}
func (ps *productService) Delete(id string) error {
	return ps.repository.Delete(id)
}
