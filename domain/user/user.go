package user

import (
	"github.com/google/uuid"
	"github.com/mariobac1/backend_webpages/model"
)

type UseCase interface {
	Create(m *model.User) error
	GetByID(ID uuid.UUID) (model.User, error)
	GetByEmail(email string) (model.User, error)
	GetAll() (model.Users, error)
	Update(m *model.User) error
	GetImage(ID uuid.UUID) (string, error)
}

type Storage interface {
	Create(m *model.User) error
	GetByID(ID uuid.UUID) (model.User, error)
	GetByEmail(email string) (model.User, error)
	GetAll() (model.Users, error)
	Update(m *model.User) error
}
