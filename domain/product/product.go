package product

import (
	"github.com/google/uuid"

	"github.com/mariobac1/backend_webpages/model"
)

type UseCase interface {
	Create(m *model.Product) error
	Update(m *model.Product) error
	GetByID(ID uuid.UUID) (model.Product, error)
	GetAll() (model.Products, error)
	Delete(ID uuid.UUID) error
	GetImage(ID uuid.UUID) (string, error)
}

type Storage interface {
	Create(m *model.Product) error
	Update(m *model.Product) error
	GetByID(ID uuid.UUID) (model.Product, error)
	GetAll() (model.Products, error)
	Delete(ID uuid.UUID) error
}
