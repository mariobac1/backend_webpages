package imagehome

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

type ImageHome struct {
	storage Storage
}

func New(s Storage) ImageHome {
	return ImageHome{storage: s}
}

func (i ImageHome) Create(m *model.ImageHome) error {
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
	fmt.Println(m)
	err = i.storage.Create(m)
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

func (i ImageHome) Update(m *model.ImageHome) error {
	m.UpdatedAt = time.Now().Unix()

	err := i.storage.Update(m)
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

func (i ImageHome) GetByID(ID uuid.UUID) (model.ImageHome, error) {
	ImageHome, err := i.storage.GetByID(ID)
	if err != nil {
		return model.ImageHome{}, fmt.Errorf("ImageHome: %w", err)
	}

	return ImageHome, nil
}

func (i ImageHome) GetAll() (model.ImageHomes, error) {
	ImageHomes, err := i.storage.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%s %w", "storage.GetAll()", err)
	}

	return ImageHomes, nil
}

func (i ImageHome) GetImage(ID uuid.UUID) (string, error) {
	var imagePath string
	path := os.Getenv("IMAGES_DIR") + "imagehome/"
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
	dst, err := os.Create(os.Getenv("IMAGES_DIR") + "imagehome/" + ID.String() + "." + extensionArchivo)
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

// Erase an file with the same name
func eraseFile(nameFile string) error {
	dirpath := os.Getenv("IMAGES_DIR") + "imagehome/"

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
	if ext == "jpg" || ext == "jpeg" || ext == "png" {
		return nil
	}
	return fmt.Errorf("Archivo no es del tipo requerido")
}
