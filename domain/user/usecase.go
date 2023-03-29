package user

import (
	"fmt"
	"os"
	"strings"
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
	//create route of image
	m.Avatar = fmt.Sprintf(os.Getenv("IMAGES_DIR"), m.Name)

	//remove all blanck space in email
	m.Email = strings.Replace(m.Email, " ", "", -1)

	//validate Fields empty
	if m.Email == "" || m.Password == "" || len(m.Password) < 9 {
		return fmt.Errorf("%s %w", "Fields can't be empty or length ", err)
	}

	//validate de password
	password, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("%s %w", "bcrypt.GenerateFromPassword", err)
	}
	m.Password = string(password)
	if m.Details == nil {
		m.Details = []byte(`[]`)
	}
	m.CreatedAt = time.Now().Unix()

	err = u.storage.Create(m)
	if err != nil {
		return fmt.Errorf("%s %w", "storage.Create", err)
	}

	m.Password = ""

	return nil
}

func (u User) Update(m *model.User) error {

	if m.Password == "" || len(m.Password) < 9 {
		return fmt.Errorf("%s %w", "Fields can't be empty or length ", nil)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("%s %w", "bcrypt.GenerateFromPassword()", err)
	}

	m.Password = string(password)

	m.UpdatedAt = time.Now().Unix()

	err = u.storage.Update(m)
	if err != nil {
		return fmt.Errorf("%s %w", "storage.Update()", err)
	}

	return nil
}

// func (u User) AdminUpdate(m *model.User) error {

// 	if m.Password == "" || len(m.Password) < 9 {
// 		return fmt.Errorf("%s %w", "Fields can't be empty or length ", nil)
// 	}

// 	password, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return fmt.Errorf("%s %w", "bcrypt.GenerateFromPassword()", err)
// 	}

// 	m.Password = string(password)

// 	m.UpdatedAt = time.Now().Unix()

// 	err = u.storage.AdminUpdate(m)
// 	if err != nil {
// 		return fmt.Errorf("%s %w", "storage.Update()", err)
// 	}

// 	return nil
// }

func (u User) GetByID(ID uuid.UUID) (model.User, error) {
	user, err := u.storage.GetByID(ID)
	if err != nil {
		return model.User{}, fmt.Errorf("user: %w", err)
	}

	user.Password = ""

	return user, nil
}

func (u User) GetByEmail(email string) (model.User, error) {
	user, err := u.storage.GetByEmail(email)
	if err != nil {
		return model.User{}, fmt.Errorf("%s %w", "storage.GetByEmail()", err)
	}
	return user, nil
}

func (u User) GetAll() (model.Users, error) {
	users, err := u.storage.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%s %w", "storage.GetAll()", err)
	}

	return users, nil
}

func (u User) Login(email, password string) (model.User, error) {
	m, err := u.GetByEmail(email)
	if err != nil {
		return model.User{}, fmt.Errorf("%s %w", "user.GetByEmail()", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(password))
	if err != nil {
		return model.User{}, fmt.Errorf("%s %w", "bcrypt.CompareHashAndPassword()", err)
	}

	m.Password = ""

	return m, nil
}

// func (u User) Image() error {

// 	return nil
// }
