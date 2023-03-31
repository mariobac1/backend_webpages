package model

import (
	"encoding/json"
	"mime/multipart"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID             `json:"id"`
	Name      string                `json:"name"`
	Email     string                `json:"email"`
	Password  string                `json:"password"`
	File      *multipart.FileHeader `json:"file" form:"file"`
	Details   json.RawMessage       `json:"details"`
	CreatedAt int64                 `json:"create_at"` //Unix time
	UpdatedAt int64                 `json:"updated_at"`
}

type Users []User
