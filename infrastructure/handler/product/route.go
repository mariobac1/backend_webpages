package product

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"github.com/mariobac1/backend_webpages/domain/product"
	"github.com/mariobac1/backend_webpages/infrastructure/handler/middle"
	storageProduct "github.com/mariobac1/backend_webpages/infrastructure/postgres/product"
)

func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	authMiddleware := middle.New()
	sendImage(e, h)
	publicRoutes(e, h)
	privateRoutes(e, h, authMiddleware.IsValid)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	storage := storageProduct.New(dbPool)
	useCase := product.New(storage)

	return newHandler(useCase)
}

func privateRoutes(e *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	g := e.Group("/api/v1/private/product", middlewares...)

	g.POST("", h.Create)
	g.PUT("", h.Update)
	g.DELETE("/:id", h.Delete)

}

func publicRoutes(e *echo.Echo, h handler) {
	g := e.Group("/api/v1/public/product")

	g.GET("", h.GetAll)
	g.GET("/:id", h.GetByID)
}

func sendImage(e *echo.Echo, h handler) {
	g := e.Group("/api/v1/img/product")

	g.GET("/:id", h.GetImage)
}
