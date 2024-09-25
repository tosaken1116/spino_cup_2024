package router

import (
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/handler"
	"github.com/tosaken1116/spino_cup_2024/backend/pkg/auth"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho" //nolint
)

func New(
	roomHandler handler.RoomHandler,
	activeRoomHandler handler.WSHandler,
	userHandler handler.UserHandler,
	authClient *auth.AuthClient,
) *echo.Echo {
	e := echo.New()
	setup(e)
	registerRoutes(e, roomHandler, authClient)
	e.GET("/rooms/:id/join", activeRoomHandler.Join)

	e.POST("/v1/signup", userHandler.SignUp)

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
		if c.Path() == "/" {
			return true
		}
		if strings.Contains(c.Request().Header.Get(echo.HeaderConnection), "Upgrade") {
			return true
		}
		return false
	})))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper: func(c echo.Context) bool {
			return c.Request().URL.Path == "/"
		},
		AllowMethods: []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowOrigins: []string{"http://localhost:5173", "https://localhost:5173", os.Getenv("ALLOW_ORIGIN")},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))
}
