package models

import "time"

type Product struct {
	ID          string         `json:"ID"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Category    string         `json:"category"`
	Price       int64          `json:"price"`
	Sizes       []string       `json:"sizes"`
	Colors      []string       `json:"colors"`
	StockBySize map[string]int `json:"stockBySize"`
	Images      []string       `json:"images"`
	IsActive    bool           `json:"isActive"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
}
