package button

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"github.com/mariobac1/backend_webpages/domain/button"
	"github.com/mariobac1/backend_webpages/infrastructure/handler/middle"
	storageButton "github.com/mariobac1/backend_webpages/infrastructure/postgres/button"
)

func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	authMiddleware := middle.New()
	publicRoutes(e, h)
	privateRoutes(e, h, authMiddleware.IsValid)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	storage := storageButton.New(dbPool)
	useCase := button.New(storage)

	return newHandler(useCase)
}

func privateRoutes(e *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	g := e.Group("/api/v1/private/button", middlewares...)

	g.POST("", h.Create)
	g.PUT("", h.Update)

}

func publicRoutes(e *echo.Echo, h handler) {
	g := e.Group("/api/v1/public/button")

	g.GET("", h.GetAll)
	g.GET("/:id", h.GetByID)
}
