package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/infra/ws"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/usecase"
)

type Base struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type ChangeCurrentPosition struct {
	Type    string       `json:"type"`
	Payload UserPosition `json:"payload"`
}

type UserPosition struct {
	ID        string  `json:"id"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Color     string  `json:"color"`
	IsClicked bool    `json:"isClicked"`
}

type ChangeCurrentScreen struct {
	Type    string     `json:"type"`
	Payload ScreenSize `json:"payload"`
}

type ScreenSize struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type WSHandler interface {
	Join(c echo.Context) error
}

type wsHandler struct {
	upgrader  *websocket.Upgrader
	uc        usecase.ActiveRoomUsecase
	msgSender *ws.MsgSender
}

func NewWSHandler(uc usecase.ActiveRoomUsecase, msgSender *ws.MsgSender) WSHandler {
	return &wsHandler{
		upgrader: &websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		uc:        uc,
		msgSender: msgSender,
	}
}

// Join implements WSHandler.
func (w *wsHandler) Join(c echo.Context) (err error) {
	roomID := c.Param("id")
	userID := uuid.New().String()

	ws, err := w.upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	errCh := make(chan error)
	defer close(errCh)

	w.msgSender.Register(userID, ws, errCh)
	defer w.msgSender.Unregister(userID)

	ctx := c.Request().Context()
	if err := w.uc.JoinRoom(ctx, userID, roomID); err != nil {
		fmt.Printf("err: %v\n", err)
		return nil
	}

	defer func() {
		if _err := w.uc.LeaveRoom(ctx, roomID, userID); _err != nil {
			errors.Join(err, _err)
		}
	}()

	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			break
		}

		var msg Base
		if err := json.Unmarshal(p, &msg); err != nil {
			break
		}

		switch msg.Type {
		case "ChangeCurrentPosition":
			var msg ChangeCurrentPosition
			if err := json.Unmarshal(p, &msg); err != nil {
				fmt.Printf("err: %v\n", err)
			}

			if err := w.uc.SendPointer(ctx, &usecase.SendPointerReq{
				UserID:    userID,
				RoomID:    roomID,
				X:         msg.Payload.X,
				Y:         msg.Payload.Y,
				Color:     msg.Payload.Color,
				IsClicked: msg.Payload.IsClicked,
			}); err != nil {
				fmt.Printf("err: %v\n", err)
			}
		case "ChangeCurrentScreen":
			var msg ChangeCurrentScreen
			if err := json.Unmarshal(p, &msg); err != nil {
				fmt.Printf("err: %v\n", err)
			}

			if err := w.uc.ChangeScreenSize(ctx, roomID, msg.Payload.Height, msg.Payload.Width); err != nil {
				fmt.Printf("err: %v\n", err)
			}
		}
	}

	return nil
}
