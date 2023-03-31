package imagehome

import (
	"github.com/google/uuid"
	"github.com/mariobac1/backend_webpages/model"
)

type UseCase interface {
	Create(m *model.ImageHome) error
	Update(m *model.ImageHome) error
	GetByID(ID uuid.UUID) (model.ImageHome, error)
	GetAll() (model.ImageHomes, error)
	GetImage(ID uuid.UUID) (string, error)
}

type Storage interface {
	Create(m *model.ImageHome) error
	Update(m *model.ImageHome) error
	GetByID(ID uuid.UUID) (model.ImageHome, error)
	GetAll() (model.ImageHomes, error)
}
