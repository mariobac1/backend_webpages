package model

import (
	"encoding/json"
	"mime/multipart"

	"github.com/google/uuid"
)

// Product model of table products
type Product struct {
	ID          uuid.UUID             `json:"id" form:"id"`
	Name        string                `json:"name" form:"name"`
	Price       float64               `json:"price" form:"price"`
	Promotion   bool                  `json:"promotion" form:"promotion"`
	Description string                `json:"description" form:"description"`
	File        *multipart.FileHeader `json:"file" form:"file"`
	Details     json.RawMessage       `json:"details" form:"details"`
	CreatedAt   int64                 `json:"created_at" form:"created_at"`
	UpdatedAt   int64                 `json:"updated_at" form:"updated_at"`
}

// Products slice of products
type Products []Product
