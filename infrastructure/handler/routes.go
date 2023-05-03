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
	"github.com/mariobac1/backend_webpages/infrastructure/handler/variablevalue"
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
	//M
	// P
	product.NewRouter(e, dbPool)
	// R
	// S
	// sendImage(e)
	// T
	// U
	user.NewRouter(e, dbPool)
	// V
	variablevalue.NewRouter(e, dbPool)
	// W
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
