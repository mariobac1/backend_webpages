package model

import (
	"encoding/json"
	"mime/multipart"

	"github.com/google/uuid"
)

// VariableValue model of table Image VariableValues
type VariableValue struct {
	ID          uuid.UUID             `json:"id" form:"id"`
	Name        string                `json:"name" form:"name"`
	Title       string                `json:"title" form:"title"`
	Paragraph   string                `json:"paragraph" form:"paragraph"`
	Color       string                `json:"color" form:"color"`
	BgColor     string                `json:"bgcolor" form:"bgcolor"`
	Font        string                `json:"font" form:"font"`
	Icon        string                `json:"icon" form:"icon"`
	Description string                `json:"description" form:"description"`
	File        *multipart.FileHeader `json:"file" form:"file"`
	Details     json.RawMessage       `json:"details" form:"details"`
	CreatedAt   int64                 `json:"created_at" form:"created_at"`
	UpdatedAt   int64                 `json:"updated_at" form:"updated_at"`
}

type VariableValues []VariableValue
