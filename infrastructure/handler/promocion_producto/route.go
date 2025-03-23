package promocionproducto

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	promocionproducto "github.com/ninosistemas10/delivery/domain/promocion_producto"
	"github.com/ninosistemas10/delivery/infrastructure/handler/middle"
	promocionProductoStorage "github.com/ninosistemas10/delivery/infrastructure/postgres/promocion_producto"
)

func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	authMiddleware := middle.New()

	adminRoutes(e, h, authMiddleware.IsValid, authMiddleware.IsAdmin)
	publicRoutes(e, h)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	useCase := promocionproducto.New(promocionProductoStorage.New(dbPool))
	return newHandler(useCase)
}

func adminRoutes(e *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	route := e.Group("ninosistemas/admin/promocion/", middlewares...)

	route.POST("", h.Create)
	route.PUT("/:id", h.Update)
	route.DELETE("/:id", h.Delete)

	// route.GET("", h.GetAll)
	route.GET("/:id", h.GetByID)
}

func publicRoutes(e *echo.Echo, h handler) {
	route := e.Group("ninosistemas/public/promocion/")

	route.POST("", h.Create)
	//route.GET("", h.GetAll)
}
