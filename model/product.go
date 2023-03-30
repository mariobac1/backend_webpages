package model

import (
	"encoding/json"
	"mime/multipart"

	"github.com/google/uuid"
)

// Product model of table products
type Product struct {
	ID          uuid.UUID             `json:"id"`
	Name        string                `json:"name"`
	Price       float64               `json:"price"`
	Description string                `json:"despcription"`
	File        *multipart.FileHeader `json:"file" form:"file" `
	Details     json.RawMessage       `json:"details"`
	CreatedAt   int64                 `json:"created_at"`
	UpdatedAt   int64                 `json:"updated_at"`
}

// Products slice of products
type Products []Product
