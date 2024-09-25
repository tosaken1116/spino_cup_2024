package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/domain/model"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/handler/schema/api/room/rpc"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/usecase"
)

type RoomHandler interface {
	CreateRoom(c echo.Context) error
	GetRoom(c echo.Context) error
	ListRoom(c echo.Context) error
	UpdateRoom(c echo.Context) error
}

type roomHandler struct {
	roomUsecase usecase.RoomUsecase
}

func NewRoomHandler(roomUsecase usecase.RoomUsecase) RoomHandler {
	return &roomHandler{roomUsecase}
}

// CreateRoom implements RoomHandler.
func (r *roomHandler) CreateRoom(c echo.Context) error {
	req := &rpc.CreateRoomRequest{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err).SetInternal(err)
	}

	room, err := r.roomUsecase.CreateRoom(c.Request().Context(), &usecase.RoomDTO{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err).SetInternal(err)
	}

	return c.JSON(http.StatusOK, room)
}

// GetRoom implements RoomHandler.
func (r *roomHandler) GetRoom(c echo.Context) error {
	id := c.Param("id")
	room, err := r.roomUsecase.GetRoom(c.Request().Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrRoomNotFound), errors.Is(err, model.ErrRoomIDInvalid):
			return echo.NewHTTPError(http.StatusNotFound, err).SetInternal(err)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err).SetInternal(err)
		}
	}

	return c.JSON(http.StatusOK, room)
}

// ListRoom implements RoomHandler.
func (r *roomHandler) ListRoom(c echo.Context) error {
	rooms, err := r.roomUsecase.ListRoom(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err).SetInternal(err)
	}

	return c.JSON(http.StatusOK, rooms)
}

// UpdateRoom implements RoomHandler.
func (r *roomHandler) UpdateRoom(c echo.Context) error {
	id := c.Param("id")
	req := &rpc.UpdateRoomRequest{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err).SetInternal(err)
	}

	room, err := r.roomUsecase.UpdateRoom(c.Request().Context(), &usecase.RoomDTO{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		switch {
		case errors.Is(err, model.ErrRoomNotFound), errors.Is(err, model.ErrRoomIDInvalid):
			return echo.NewHTTPError(http.StatusNotFound, err).SetInternal(err)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err).SetInternal(err)
		}
	}

	return c.JSON(http.StatusOK, room)
}
