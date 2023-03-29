package login

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"github.com/mariobac1/backend_webpages/domain/login"
	"github.com/mariobac1/backend_webpages/domain/user"
	userStorage "github.com/mariobac1/backend_webpages/infrastructure/postgres/user"
)

// NewRouter returns a router to handle model.Login requests
func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	publicRoutes(e, h)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	useCaseUser := user.New(userStorage.New(dbPool))
	useCase := login.New(useCaseUser)
	return newHandler(useCase)
}

// publicRoutes handle the routes that not requires a validation of any kind to be use
func publicRoutes(e *echo.Echo, h handler) {
	route := e.Group("/api/v1/public/login")

	route.POST("", h.Login)
}
