package user

import "github.com/mariobac1/backend_webpages/model"

type UseCase interface {
	Create(m *model.User) error
	GetByEmail(email string) (model.User, error)
	GetAll() (model.Users, error)
	Update(m *model.User) error
}

type Storage interface {
	Create(m *model.User) error
	GetByEmail(email string) (model.User, error)
	GetAll() (model.User, error)
	Update(m *model.User) error
}
