package product

import (
	"fmt"
	"io"
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

	err = saveImage(m.ID, m.File)
	if err != nil {
		return fmt.Errorf("%s %w", "saveImage()", err)
	}

	err = p.storage.Create(m)
	if err != nil {
		return fmt.Errorf("%s %w", "storage.Create()", err)
	}

	return nil
}

func (p Product) Update(m *model.Product) error {

	m.UpdatedAt = time.Now().Unix()

	err := p.storage.Update(m)
	if err != nil {
		return fmt.Errorf("%s %w", "storage.Update()", err)
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
		if info.Name() == nameFile {
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
	if ext == "jpg" || ext == "jepg" || ext == "png" {
		return nil
	}
	return fmt.Errorf("Archivo no es del tipo requerido")
}
