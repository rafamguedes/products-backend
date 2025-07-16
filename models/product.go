package models

import (
	"time"
)

type Product struct {
	ID            int       `json:"id" db:"id"`
	Name          string    `json:"name" db:"name" binding:"required"`
	Description   string    `json:"description" db:"description"`
	Price         float64   `json:"price" db:"price" binding:"required"`
	Category      string    `json:"category" db:"category"`
	StockQuantity int       `json:"stock_quantity" db:"stock_quantity"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

type CreateProductRequest struct {
	Name          string  `json:"name" binding:"required"`
	Description   string  `json:"description"`
	Price         float64 `json:"price" binding:"required"`
	Category      string  `json:"category"`
	StockQuantity int     `json:"stock_quantity"`
}

type UpdateProductRequest struct {
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	Category      string  `json:"category"`
	StockQuantity int     `json:"stock_quantity"`
}
