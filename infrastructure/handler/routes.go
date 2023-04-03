package handler

import (
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"github.com/mariobac1/backend_webpages/infrastructure/handler/button"
	"github.com/mariobac1/backend_webpages/infrastructure/handler/imagehome"
	"github.com/mariobac1/backend_webpages/infrastructure/handler/login"
	"github.com/mariobac1/backend_webpages/infrastructure/handler/product"
	"github.com/mariobac1/backend_webpages/infrastructure/handler/user"
)

func InitRoutes(e *echo.Echo, dbPool *pgxpool.Pool) {
	health(e)
	// A
	// B
	button.NewRouter(e, dbPool)
	// C
	//...
	// E
	// F
	// L
	login.NewRouter(e, dbPool)
	// H
	imagehome.NewRouter(e, dbPool)
	//I
	// image(e)
	// P
	product.NewRouter(e, dbPool)
	// R
	// S
	// sendImage(e)
	// T
	// U
	user.NewRouter(e, dbPool)
}

func health(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			map[string]string{
				"time":         time.Now().String(),
				"message":      "Hello World",
				"service_name": "",
			},
		)
	})
}

// func image(e *echo.Echo) {
// 	var m model.Imagen

// 	// e.Static("/imagenes", "public") //revisar si necesitamos esta línea
// 	e.POST("/api/v1/image", func(c echo.Context) error {
// 		if err := c.Bind(&m); err != nil {
// 			fmt.Printf("\n%v\n", m)
// 			fmt.Println(m)
// 			fmt.Println("No se pudo hacer el Bind", err)
// 			return c.JSON(http.StatusConflict, err)
// 		}

// 		fmt.Println(m.Nombre)
// 		// fmt.Println(m.Apellido)
// 		//-----------
// 		// Read file
// 		//-----------

// 		// Source
// 		file, err := c.FormFile("documento")
// 		if err != nil {
// 			return err
// 		}
// 		src, err := file.Open()
// 		if err != nil {
// 			return err
// 		}
// 		defer src.Close()
// 		// leemos la extensión
// 		nombreArchivo := strings.Split(file.Filename, ".")
// 		extensionArchivo := nombreArchivo[len(nombreArchivo)-1]
// 		eraseFile("logo")
// 		// Destination
// 		// dst, err := os.Create("./imagenes/" + file.Filename)
// 		dst, err := os.Create("./imagenes/logo." + extensionArchivo)
// 		if err != nil {
// 			return err
// 		}
// 		defer dst.Close()

// 		// Copy
// 		if _, err = io.Copy(dst, src); err != nil {
// 			return err
// 		}

// 		return c.JSON(http.StatusOK, "Archivo guardado exitosamente")
// 	})
// }

// func eraseFile(nameFile string) {
// 	dirpath := "."
// 	err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
// 		if info.Name() == nameFile {
// 			err = os.Remove(path)
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 			fmt.Println("Archivo eliminado: ", path)
// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		fmt.Printf("error walking the path %v\n", err)
// 	}
// }

// func sendImage(e *echo.Echo) {
// 	e.GET("/api/v1/image", func(c echo.Context) error {
// 		var imagePath string
// 		files, err := ioutil.ReadDir("imagenes")
// 		if err != nil {
// 			return err
// 		}
// 		for _, f := range files {
// 			if strings.HasPrefix(f.Name(), "logo") {
// 				imagePath = "imagenes/" + f.Name()
// 				break
// 			}
// 		}
// 		return c.File(imagePath)
// 	})
// }
