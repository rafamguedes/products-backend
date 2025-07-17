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

type ProductFilter struct {
	Name     string   `json:"name" form:"name"`
	Category string   `json:"category" form:"category"`
	MinPrice *float64 `json:"min_price" form:"min_price"`
	MaxPrice *float64 `json:"max_price" form:"max_price"`
	MinStock *int     `json:"min_stock" form:"min_stock"`
	MaxStock *int     `json:"max_stock" form:"max_stock"`
}

type NextTokenRequest struct {
	Row   int    `json:"row" form:"row"`
	Order string `json:"order" form:"order"`
	Limit int    `json:"limit" form:"limit"`
}

type ProductFilterRequest struct {
	Filter    ProductFilter    `json:"filter"`
	NextToken NextTokenRequest `json:"next_token"`
}

type ProductFilterResponse struct {
	Data      []Product         `json:"data"`
	Total     int               `json:"total"`
	NextToken *NextTokenRequest `json:"next_token,omitempty"`
	HasMore   bool              `json:"has_more"`
}
