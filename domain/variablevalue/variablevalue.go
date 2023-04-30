package variablevalue

import (
	"github.com/google/uuid"
	"github.com/mariobac1/backend_webpages/model"
)

type UseCase interface {
	Create(m *model.VariableValue) error
	Update(m *model.VariableValue) error
	GetByID(ID uuid.UUID) (model.VariableValue, error)
	GetAll() (model.VariableValues, error)
	GetImage(ID uuid.UUID) (string, error)
}

type Storage interface {
	Create(m *model.VariableValue) error
	Update(m *model.VariableValue) error
	GetByID(ID uuid.UUID) (model.VariableValue, error)
	GetAll() (model.VariableValues, error)
}
