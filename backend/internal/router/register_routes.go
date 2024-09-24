package router

import (
	"github.com/labstack/echo/v4"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/handler"
)

func registerRoutes(
	e *echo.Echo,
	roomHandler handler.RoomHandler,
) {
	e.POST("/v1/rooms", roomHandler.CreateRoom)
	e.GET("/v1/rooms/:id", roomHandler.GetRoom)
	e.PUT("/v1/rooms/:id", roomHandler.UpdateRoom)
	e.POST("/v1/rooms/:id/join", roomHandler.JoinRoom)
}
