package user

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mariobac1/backend_webpages/model"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	storage Storage
}

func New(s Storage) User {
	return User{storage: s}
}

func (u User) Create(m *model.User) error {
	ID, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("%s %w", "uuid.NewUUID", err)
	}

	m.ID = ID
	password, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("%s %w", "bcrypt.GenerateFromPassword", err)
	}
	m.Password = string(password)
	if m.Details == nil {
		m.Details = []byte("{}")
	}
	m.CreatedAt = time.Now().Unix()

	err = u.storage.Create(m)
	if err != nil {
		return fmt.Errorf("%s %w", "storage.Create", err)
	}

	m.Password = ""

	return nil
}
