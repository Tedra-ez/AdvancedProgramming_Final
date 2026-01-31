package repository

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"sync"
	"time"

	"github.com/Tedra-ez/AdvancedProgramming_Final/models"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
	ErrNotExist      = errors.New("not exists")
)

type ProductRepository interface {
	FindAll() ([]models.Product, error)
	FindByID(ID string) (models.Product, error)
	Insert(product models.Product) error
	Update(id string, product models.Product) error
	Delete(id string) error
}

type DemoProductRepository struct {
	productsDB map[string]models.Product
	mu         sync.RWMutex
}

func New() *DemoProductRepository {
	return &DemoProductRepository{
		productsDB: make(map[string]models.Product, 10),
	}
}

func (pr *DemoProductRepository) FindAll() ([]models.Product, error) {
	pr.mu.Lock()
	defer pr.mu.Unlock()

	productSlice := make([]models.Product, 0, len(pr.productsDB))

	for _, product := range pr.productsDB {
		productSlice = append(productSlice, product)
	}
	return productSlice, nil
}
func (pr *DemoProductRepository) FindByID(ID string) (models.Product, error) {
	pr.mu.Lock()
	defer pr.mu.Unlock()
	product, found := pr.productsDB[ID]
	if !found {
		return models.Product{}, ErrNotFound
	}
	return product, nil
}
func (pr *DemoProductRepository) Insert(product models.Product) error {
	pr.mu.Lock()
	defer pr.mu.Unlock()
	_, isExist := pr.productsDB[product.ID]
	if isExist {
		return ErrAlreadyExists
	}
	if product.ID == "" {
		b := make([]byte, 8)
		if _, err := rand.Read(b); err == nil {
			product.ID = hex.EncodeToString(b)
		} else {
			product.ID = time.Now().Format("20060102150405.000")
		}
	}
	if product.CreatedAt.IsZero() {
		product.CreatedAt = time.Now()
	}
	pr.productsDB[product.ID] = product
	return nil
}

func (pr *DemoProductRepository) Update(id string, newProduct models.Product) error {
	pr.mu.Lock()
	defer pr.mu.Unlock()
	_, isExist := pr.productsDB[id]
	if !isExist {
		return ErrNotExist
	}

	pr.productsDB[id] = newProduct
	return nil
}

func (pr *DemoProductRepository) Delete(id string) error {
	pr.mu.Lock()
	defer pr.mu.Unlock()
	if _, ok := pr.productsDB[id]; !ok {
		return ErrNotExist
	}
	delete(pr.productsDB, id)
	return nil
}
