package login

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"github.com/ninosistemas10/delivery/domain/login"
	"github.com/ninosistemas10/delivery/domain/user"

	userStorage "github.com/ninosistemas10/delivery/infrastructure/postgres/user_r"
)

// NewRouter returns a router to handle model.Login requests
func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	// build middlewares to validate permissions on the routes

	publicRoutes(e, h)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	useCaseUser := user.New(userStorage.New(dbPool))
	useCase := login.New(useCaseUser)
	return newHandler(useCase)
}

// publicRoutes handle the routes that not requires a validation of any kind to be use
func publicRoutes(e *echo.Echo, h handler) {
	route := e.Group("/ninosistemas/public/login")

	route.POST("", h.Login)
}
