package model

import (
	"encoding/json"
	"mime/multipart"

	"github.com/google/uuid"
)

// ImageHome model of table Image imagehomes
type ImageHome struct {
	ID          uuid.UUID             `json:"id" form:"id"`
	Name        string                `json:"name" form:"name"`
	Color       string                `json:"color" form:"color"`
	Description string                `json:"description" form:"description"`
	File        *multipart.FileHeader `json:"file" form:"file"`
	Details     json.RawMessage       `json:"details" form:"details"`
	CreatedAt   int64                 `json:"created_at" form:"created_at"`
	UpdatedAt   int64                 `json:"updated_at" form:"updated_at"`
}

type ImageHomes []ImageHome
