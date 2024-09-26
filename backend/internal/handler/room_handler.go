package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/domain/model"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/handler/schema/api/room/rpc"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/usecase"
	"github.com/tosaken1116/spino_cup_2024/backend/pkg/auth"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatarUrl"`
}

type RoomResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerID     string `json:"ownerId"`
	Owner       *User  `json:"owner,omitempty"`
}

func NewRoomResponseFromDTO(dto *usecase.RoomDTO) *RoomResponse {
	resp := &RoomResponse{
		ID:          dto.ID,
		Name:        dto.Name,
		Description: dto.Description,
		OwnerID:     dto.OwnerID,
	}
	if dto.Owner != nil {
		resp.Owner = &User{
			ID:        dto.Owner.ID,
			Name:      dto.Owner.Name,
			AvatarURL: dto.Owner.AvatarURL,
		}
	}

	return resp
}

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
	var req rpc.CreateRoomRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err).SetInternal(err)
	}

	userID, err := auth.GetUserID(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	roomDTO, err := r.roomUsecase.CreateRoom(c.Request().Context(), &usecase.RoomDTO{
		Name:        req.Name,
		Description: req.Description,
		OwnerID:     userID,
	})
	if err != nil {
		switch {
		case errors.Is(err, model.ErrRoomNameRequired):
			return echo.NewHTTPError(http.StatusBadRequest, err).SetInternal(err)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
		}
	}

	roomResponse := NewRoomResponseFromDTO(roomDTO)

	return c.JSON(http.StatusOK, echo.Map{"room": roomResponse})
}

// GetRoom implements RoomHandler.
func (r *roomHandler) GetRoom(c echo.Context) error {
	id := c.Param("id")
	roomDTO, err := r.roomUsecase.GetRoom(c.Request().Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrRoomNotFound), errors.Is(err, model.ErrRoomIDInvalid):
			return echo.NewHTTPError(http.StatusNotFound, err).SetInternal(err)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
		}
	}

	roomResponse := NewRoomResponseFromDTO(roomDTO)

	return c.JSON(http.StatusOK, echo.Map{"room": roomResponse})
}

// ListRoom implements RoomHandler.
func (r *roomHandler) ListRoom(c echo.Context) error {
	rooms, err := r.roomUsecase.ListRoom(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	roomResponses := make([]*RoomResponse, 0, len(rooms))
	for _, room := range rooms {
		roomResponses = append(roomResponses, NewRoomResponseFromDTO(room))
	}

	return c.JSON(http.StatusOK, echo.Map{"rooms": roomResponses})
}

// UpdateRoom implements RoomHandler.
func (r *roomHandler) UpdateRoom(c echo.Context) error {
	id := c.Param("id")
	var req rpc.UpdateRoomRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err).SetInternal(err)
	}

	roomDTO, err := r.roomUsecase.UpdateRoom(c.Request().Context(), &usecase.RoomDTO{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		switch {
		case errors.Is(err, model.ErrRoomNotFound), errors.Is(err, model.ErrRoomIDInvalid):
			return echo.NewHTTPError(http.StatusNotFound, err).SetInternal(err)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
		}
	}

	roomResponse := NewRoomResponseFromDTO(roomDTO)

	return c.JSON(http.StatusOK, echo.Map{"room": roomResponse})
}
