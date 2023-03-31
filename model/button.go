package model

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Button struct {
	ID        uuid.UUID       `json:"id" form:"id"`
	Name      string          `json:"name" form:"name"`
	Color     string          `json:"color" form:"color"`
	Shape     string          `json:"shape" form:"shape"`
	Details   json.RawMessage `json:"details" form:"details"`
	CreatedAt int64           `json:"created_at" form:"created_at"`
	UpdatedAt int64           `json:"updated_at" form:"updated_at"`
}

type Buttons []Button
