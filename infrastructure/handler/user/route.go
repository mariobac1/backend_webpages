package user

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"github.com/mariobac1/backend_webpages/domain/user"
	"github.com/mariobac1/backend_webpages/infrastructure/handler/middle"
	storageUser "github.com/mariobac1/backend_webpages/infrastructure/postgres/user"
)

func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	authMiddleware := middle.New()
	privateRoutes(e, h, authMiddleware.IsValid)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	storage := storageUser.New(dbPool)
	useCase := user.New(storage)

	return newHandler(useCase)
}

func privateRoutes(e *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	g := e.Group("/api/v1/private/users", middlewares...)

	g.GET("", h.MySelf)
	g.GET("/:id", h.GetByID)
	g.PUT("/:id", h.Update)
}
