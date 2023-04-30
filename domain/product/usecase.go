package product

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mariobac1/backend_webpages/model"
)

type Product struct {
	storage Storage
}

func New(s Storage) Product {
	return Product{storage: s}
}

func (p Product) Create(m *model.Product) error {
	ID, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("%s %w", "uuid.NewUUID()", err)
	}
	m.ID = ID

	if m.Name == "" || m.Description == "" {
		return fmt.Errorf("%s %w", "Fields can't be empty ", err)
	}

	if len(m.Details) == 0 {
		m.Details = []byte(`[]`)
	}

	m.CreatedAt = time.Now().Unix()

	err = p.storage.Create(m)
	if err != nil {
		return fmt.Errorf("%s %w", "storage.Create()", err)
	}

	if m.File != nil {
		err = saveImage(m.ID, m.File)
		if err != nil {
			return fmt.Errorf("%s %w", "saveImage()", err)
		}
	}

	return nil
}

func (p Product) Update(m *model.Product) error {

	m.UpdatedAt = time.Now().Unix()

	err := p.storage.Update(m)
	if err != nil {
		return fmt.Errorf("%s %w", "storage.Update()", err)
	}

	if m.File != nil {
		err = saveImage(m.ID, m.File)
		if err != nil {
			return fmt.Errorf("%s %w", "saveImage()", err)
		}
	}

	return nil
}

func (p Product) GetByID(ID uuid.UUID) (model.Product, error) {
	Product, err := p.storage.GetByID(ID)
	if err != nil {
		return model.Product{}, fmt.Errorf("Product: %w", err)
	}

	return Product, nil
}

func (p Product) GetAll() (model.Products, error) {
	Products, err := p.storage.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%s %w", "storage.GetAll()", err)
	}

	return Products, nil
}

func (p Product) Delete(ID uuid.UUID) error {
	err := p.storage.Delete(ID)
	if err != nil {
		return fmt.Errorf("Booking: %w", err)
	}

	return nil
}

func (p Product) GetImage(ID uuid.UUID) (string, error) {
	var imagePath string
	path := os.Getenv("IMAGES_DIR") + "products/"
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

func saveImage(ID uuid.UUID, file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	// leemos la extensión
	nombreArchivo := strings.Split(file.Filename, ".")
	extensionArchivo := nombreArchivo[len(nombreArchivo)-1]
	err = validateExt(extensionArchivo)
	if err != nil {
		return err
	}

	err = eraseFile(ID.String())
	if err != nil {
		return err
	}
	// Destination
	dst, err := os.Create(os.Getenv("IMAGES_DIR") + "products/" + ID.String() + "." + extensionArchivo)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}

func eraseFile(nameFile string) error {
	dirpath := os.Getenv("IMAGES_DIR") + "products/"

	err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if strings.TrimSuffix(info.Name(), filepath.Ext(info.Name())) == nameFile {
			err = os.Remove(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func validateExt(ext string) error {
	if ext == "jpg" || ext == "jpeg" || ext == "png" {
		return nil
	}
	return fmt.Errorf("Archivo no es del tipo requerido")
}
