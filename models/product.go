package models

import "time"

// it should be similar to the ERD scheme
type Product struct {
	ID          string         `json:"id" bson:"_id,omitempty"`
	Name        string         `json:"name" bson:"name"`
	Description string         `json:"description" bson:"description"`
	Category    string         `json:"category" bson:"category"`
	Price       float64        `json:"price" bson:"price"`
	Sizes       []string       `json:"sizes" bson:"sizes"`
	Colors      []string       `json:"colors" bson:"colors"`
	StockBySize map[string]int `json:"stock_by_size" bson:"stockBySize"`
	Images      []string       `json:"images" bson:"images"`
	IsActive    bool           `json:"is_active" bson:"isActive"`
	CreatedAt   time.Time      `json:"created_at" bson:"createdAt"`
	UpdateAt    time.Time      `json:"updated_at" bson:"updateAt"`
}

type CreateProductRequest struct {
	Name        string         `json:"name" binding:"required"`
	Description string         `json:"description"`
	Category    string         `json:"category"`
	Price       float64        `json:"price" binding:"required"`
	Sizes       []string       `json:"sizes"`
	Colors      []string       `json:"colors"`
	StockBySize map[string]int `json:"stock_by_size"`
	Images      []string       `json:"images"`
}
