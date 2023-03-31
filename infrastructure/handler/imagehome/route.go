package imagehome

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"github.com/mariobac1/backend_webpages/domain/imagehome"
	"github.com/mariobac1/backend_webpages/infrastructure/handler/middle"
	storageImageHome "github.com/mariobac1/backend_webpages/infrastructure/postgres/imagehome"
)

func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	authMiddleware := middle.New()
	sendImage(e, h)
	publicRoutes(e, h)
	privateRoutes(e, h, authMiddleware.IsValid)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	storage := storageImageHome.New(dbPool)
	useCase := imagehome.New(storage)

	return newHandler(useCase)
}

func privateRoutes(e *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	g := e.Group("/api/v1/private/imagehome", middlewares...)

	g.POST("", h.Create)
	g.PUT("", h.Update)

}

func publicRoutes(e *echo.Echo, h handler) {
	g := e.Group("/api/v1/public/imagehome")

	g.GET("", h.GetAll)
	g.GET("/:id", h.GetByID)
}

func sendImage(e *echo.Echo, h handler) {
	g := e.Group("/api/v1/img/imagehome")

	g.GET("/:id", h.GetImage)
}
