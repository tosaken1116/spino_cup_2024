package router

import (
	"github.com/labstack/echo/v4"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/handler"
	"github.com/tosaken1116/spino_cup_2024/backend/pkg/auth"
)

func registerRoutes(
	e *echo.Echo,
	roomHandler handler.RoomHandler,
	authClient *auth.AuthClient,
) {
	// e.POST("/v1/rooms", roomHandler.CreateRoom, authClient.Middleware)
	e.POST("/v1/rooms", roomHandler.CreateRoom)
	e.GET("/v1/rooms/:id", roomHandler.GetRoom)
	e.PUT("/v1/rooms/:id", roomHandler.UpdateRoom)
	e.GET("/v1/rooms", roomHandler.ListRoom)
}
