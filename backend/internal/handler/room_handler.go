package handler

import "github.com/labstack/echo/v4"

type RoomHandler interface {
	CreateRoom(c echo.Context) error
	GetRoom(c echo.Context) error
	UpdateRoom(c echo.Context) error
}

type roomHandler struct {
}

func NewRoomHandler() RoomHandler {
	return &roomHandler{}
}

// CreateRoom implements RoomHandler.
func (r *roomHandler) CreateRoom(c echo.Context) error {
	panic("unimplemented")
}

// GetRoom implements RoomHandler.
func (r *roomHandler) GetRoom(c echo.Context) error {
	panic("unimplemented")
}

// UpdateRoom implements RoomHandler.
func (r *roomHandler) UpdateRoom(c echo.Context) error {
	panic("unimplemented")
}
