package variablevalue

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"github.com/mariobac1/backend_webpages/domain/variablevalue"
	"github.com/mariobac1/backend_webpages/infrastructure/handler/middle"
	storageVariableValue "github.com/mariobac1/backend_webpages/infrastructure/postgres/variablevalue"
)

func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	authMiddleware := middle.New()
	sendImage(e, h)
	publicRoutes(e, h)
	privateRoutes(e, h, authMiddleware.IsValid)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	storage := storageVariableValue.New(dbPool)
	useCase := variablevalue.New(storage)

	return newHandler(useCase)
}

func privateRoutes(e *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	g := e.Group("/api/v1/private/variablevalue", middlewares...)

	g.POST("", h.Create)
	g.PUT("/:id", h.Update)

}

func publicRoutes(e *echo.Echo, h handler) {
	g := e.Group("/api/v1/public/variablevalue")

	g.GET("", h.GetAll)
	g.GET("/:id", h.GetByID)
}

func sendImage(e *echo.Echo, h handler) {
	g := e.Group("/api/v1/img/variablevalue")

	g.GET("/:id", h.GetImage)
}
