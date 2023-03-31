package button

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/mariobac1/backend_webpages/model"
)

type Button struct {
	storage Storage
}

func New(s Storage) Button {
	return Button{storage: s}
}

func (b Button) Create(m *model.Button) error {
	ID, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("%s %w", "uuid.NewUUID()", err)
	}
	m.ID = ID

	if m.Name == "" || m.Color == "" || m.Shape == "" {
		return fmt.Errorf("%s %w", "Fields can't be empty ", err)
	}

	if len(m.Details) == 0 {
		m.Details = []byte(`[]`)
	}

	m.CreatedAt = time.Now().Unix()

	err = b.storage.Create(m)
	if err != nil {
		return fmt.Errorf("%s %w", "storage.Create()", err)
	}

	return nil
}

func (b Button) Update(m *model.Button) error {

	m.UpdatedAt = time.Now().Unix()

	err := b.storage.Update(m)
	if err != nil {
		return fmt.Errorf("%s %w", "storage.Update()", err)
	}

	return nil
}

func (p Button) GetByID(ID uuid.UUID) (model.Button, error) {
	Button, err := p.storage.GetByID(ID)
	if err != nil {
		return model.Button{}, fmt.Errorf("Button: %w", err)
	}

	return Button, nil
}

func (b Button) GetAll() (model.Buttons, error) {
	Buttons, err := b.storage.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%s %w", "storage.GetAll()", err)
	}

	return Buttons, nil
}

func (b Button) GetImage(ID uuid.UUID) (string, error) {
	var imagePath string
	path := os.Getenv("IMAGES_DIR") + "Buttons/"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return "", err
	}
	for _, f := range files {
		if strings.HasPrefix(f.Name(), ID.String()) {
			imagePath = path + f.Name()
			break
		}
	}
	return imagePath, nil

}
