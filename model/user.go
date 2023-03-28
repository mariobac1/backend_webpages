package model

import (
	"encoding/json"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID       `json:"id"`
	Name      string          `json:"name"`
	Email     string          `json:"email"`
	Password  string          `json:"password"`
	Avatar    string          `json:"avatar"`
	Details   json.RawMessage `json:"details"`
	CreatedAt int64           `json:"create_at"` //Unix time
	UpdatedAt int64           `json:"updated_at"`
}

type Users []User
