package handler

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"github.com/mariobac1/backend_webpages/infrastructure/handler/button"
	"github.com/mariobac1/backend_webpages/infrastructure/handler/imagehome"
	"github.com/mariobac1/backend_webpages/infrastructure/handler/login"
	"github.com/mariobac1/backend_webpages/infrastructure/handler/product"
	"github.com/mariobac1/backend_webpages/infrastructure/handler/user"
)

const pathIcon = "./public/icons/social_network/"

func InitRoutes(e *echo.Echo, dbPool *pgxpool.Pool) {
	health(e)
	// A
	// B
	button.NewRouter(e, dbPool)
	// C
	//...
	// E
	// F
	facebookIcon(e)
	facebookGrayIcon(e)
	// L
	login.NewRouter(e, dbPool)
	// H
	imagehome.NewRouter(e, dbPool)
	//I
	instagramIcon(e)
	instagramGrayIcon(e)
	// image(e)
	//M
	msnFBIcon(e)
	msnFBGrayIcon(e)
	// P
	product.NewRouter(e, dbPool)
	// R
	// S
	// sendImage(e)
	// T
	twitterIcon(e)
	twitterGrayIcon(e)
	// U
	user.NewRouter(e, dbPool)
	// V
	// W
	whatsappIcon(e)
	whatsappGrayIcon(e)
	// Y
	// Z
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

func facebookIcon(e *echo.Echo) {
	e.GET("/api/v1/facebookIcon", func(c echo.Context) error {
		var imagePath string
		files, err := ioutil.ReadDir(pathIcon)
		if err != nil {
			return err
		}
		for _, f := range files {
			if strings.HasPrefix(f.Name(), "facebook") {
				imagePath = pathIcon + f.Name()
				break
			}
		}
		return c.File(imagePath)
	})
}

func facebookGrayIcon(e *echo.Echo) {
	e.GET("/api/v1/facebookGrayIcon", func(c echo.Context) error {
		var imagePath string
		files, err := ioutil.ReadDir(pathIcon)
		if err != nil {
			return err
		}
		for _, f := range files {
			if strings.HasPrefix(f.Name(), "facebook_gray") {
				imagePath = pathIcon + f.Name()
				break
			}
		}
		return c.File(imagePath)
	})
}

func twitterIcon(e *echo.Echo) {
	e.GET("/api/v1/TwitterIcon", func(c echo.Context) error {
		var imagePath string
		files, err := ioutil.ReadDir(pathIcon)
		if err != nil {
			return err
		}
		for _, f := range files {
			if strings.HasPrefix(f.Name(), "twitter") {
				imagePath = pathIcon + f.Name()
				break
			}
		}
		return c.File(imagePath)
	})
}

func twitterGrayIcon(e *echo.Echo) {
	e.GET("/api/v1/TwitterGrayIcon", func(c echo.Context) error {
		var imagePath string
		files, err := ioutil.ReadDir(pathIcon)
		if err != nil {
			return err
		}
		for _, f := range files {
			if strings.HasPrefix(f.Name(), "twitter_gray") {
				imagePath = pathIcon + f.Name()
				break
			}
		}
		return c.File(imagePath)
	})
}

func instagramIcon(e *echo.Echo) {
	e.GET("/api/v1/instagramIcon", func(c echo.Context) error {
		var imagePath string
		files, err := ioutil.ReadDir(pathIcon)
		if err != nil {
			return err
		}
		for _, f := range files {
			if strings.HasPrefix(f.Name(), "instagram") {
				imagePath = pathIcon + f.Name()
				break
			}
		}
		return c.File(imagePath)
	})
}

func instagramGrayIcon(e *echo.Echo) {
	e.GET("/api/v1/instagramGrayIcon", func(c echo.Context) error {
		var imagePath string
		files, err := ioutil.ReadDir(pathIcon)
		if err != nil {
			return err
		}
		for _, f := range files {
			if strings.HasPrefix(f.Name(), "instagram_gray") {
				imagePath = pathIcon + f.Name()
				break
			}
		}
		return c.File(imagePath)
	})
}

func whatsappIcon(e *echo.Echo) {
	e.GET("/api/v1/whatsappIcon", func(c echo.Context) error {
		var imagePath string
		files, err := ioutil.ReadDir(pathIcon)
		if err != nil {
			return err
		}
		for _, f := range files {
			if strings.HasPrefix(f.Name(), "whatsapp") {
				imagePath = pathIcon + f.Name()
				break
			}
		}
		return c.File(imagePath)
	})
}

func whatsappGrayIcon(e *echo.Echo) {
	e.GET("/api/v1/whatsappGrayIcon", func(c echo.Context) error {
		var imagePath string
		files, err := ioutil.ReadDir(pathIcon)
		if err != nil {
			return err
		}
		for _, f := range files {
			if strings.HasPrefix(f.Name(), "whatsapp_gray") {
				imagePath = pathIcon + f.Name()
				break
			}
		}
		return c.File(imagePath)
	})
}

func msnFBIcon(e *echo.Echo) {
	e.GET("/api/v1/msnfbIcon", func(c echo.Context) error {
		var imagePath string
		files, err := ioutil.ReadDir(pathIcon)
		if err != nil {
			return err
		}
		for _, f := range files {
			if strings.HasPrefix(f.Name(), "fb_messenger") {
				imagePath = pathIcon + f.Name()
				break
			}
		}
		return c.File(imagePath)
	})
}

func msnFBGrayIcon(e *echo.Echo) {
	e.GET("/api/v1/msnfbGrayIcon", func(c echo.Context) error {
		var imagePath string
		files, err := ioutil.ReadDir(pathIcon)
		if err != nil {
			return err
		}
		for _, f := range files {
			if strings.HasPrefix(f.Name(), "fb_messenger") {
				imagePath = pathIcon + f.Name()
				break
			}
		}
		return c.File(imagePath)
	})
}
