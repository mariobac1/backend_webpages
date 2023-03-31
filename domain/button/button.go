package button

import (
	"github.com/google/uuid"
	"github.com/mariobac1/backend_webpages/model"
)

type UseCase interface {
	Create(m *model.Button) error
	Update(m *model.Button) error
	GetByID(ID uuid.UUID) (model.Button, error)
	GetAll() (model.Buttons, error)
}

type Storage interface {
	Create(m *model.Button) error
	Update(m *model.Button) error
	GetByID(ID uuid.UUID) (model.Button, error)
	GetAll() (model.Buttons, error)
}
