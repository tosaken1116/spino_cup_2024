package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/handler"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

func New(roomHandler handler.RoomHandler, activeRoomHandler handler.WSHandler) *echo.Echo {
	e := echo.New()
	setup(e)
	registerRoutes(e, roomHandler)
	e.GET("/rooms/:id/join", activeRoomHandler.Join)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK!")
	})

	return e
}

func setup(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			return c.Request().URL.Path == "/"
		},
	}))
	e.Use(otelecho.Middleware("api.spino.kurichi.dev", otelecho.WithSkipper(func(c echo.Context) bool {
		return c.Request().URL.Path == "/" || c.Request().URL.Path == "/rooms/:id/join"
	})))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper: func(c echo.Context) bool {
			return c.Request().URL.Path == "/"
		},
		AllowMethods: []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowOrigins: []string{"http://localhost:5173"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))
}
