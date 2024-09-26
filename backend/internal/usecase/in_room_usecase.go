package usecase

import (
	"context"
	"sync"

	"github.com/tosaken1116/spino_cup_2024/backend/internal/domain/model"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/domain/repository"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/domain/service"
)

type SendPointerReq struct {
	RoomID    string
	UserID    string
	X         float64
	Y         float64
	Color     string
	IsClicked bool
	PenSize   int
}

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatarUrl"`
}

type Pointer struct {
	User      *User   `json:"user"`
	IsClicked bool    `json:"isClicked"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Color     string  `json:"color"`
	PenSize   int     `json:"penSize"`
}

type JoinRoomResp struct {
	UserID       string  `json:"userId"`
	OwnerID      string  `json:"ownerId"`
	ScreenHeight float64 `json:"height"`
	ScreenWidth  float64 `json:"width"`
}

type ActiveRoomUsecase interface {
	JoinRoom(ctx context.Context, userID, roomID string) error
	SendPointer(ctx context.Context, req *SendPointerReq) error
	ChangeScreenSize(ctx context.Context, roomID string, height, width int) error
	LeaveRoom(ctx context.Context, roomID, userID string) error
}

type activeRoomUsecase struct {
	msgSender service.MessageSender
	repo      repository.ActiveRoomRepo
	rRepo     repository.RoomRepository
	uRepo     repository.UserRepository
}

func NewActiveRoomUsecase(
	msgSender service.MessageSender,
	repo repository.ActiveRoomRepo,
	rRepo repository.RoomRepository,
	uRepo repository.UserRepository,
) ActiveRoomUsecase {
	return &activeRoomUsecase{
		msgSender: msgSender,
		repo:      repo,
		rRepo:     rRepo,
		uRepo:     uRepo,
	}
}

// JoinRoom implements InRoomUsecase.
func (i *activeRoomUsecase) JoinRoom(ctx context.Context, userID, roomID string) error {
	room, err := i.repo.Get(ctx, roomID)
	if err != nil {
		id, err := model.ParseRoomID(roomID)
		if err != nil {
			return err
		}

		_room, err := i.rRepo.GetRoom(ctx, id)
		if err != nil {
			return err
		}

		room = &model.AcitveRoom{
			Room:    _room,
			OwnerID: userID,
			Lock:    sync.RWMutex{},
		}
	}

	user, err := i.uRepo.GetUser(ctx, userID)
	if err != nil {
		return err
	}

	room.AddUser(user)
	if err := i.repo.Store(ctx, room); err != nil {
		return err
	}

	room.Lock.RLock()
	defer room.Lock.RUnlock()

	msg := &JoinRoomResp{
		UserID:       user.ID,
		OwnerID:      room.OwnerID,
		ScreenHeight: float64(room.ScreenHeight),
		ScreenWidth:  float64(room.ScreenWidth),
	}

	if err := i.msgSender.Send(ctx, userID, map[string]interface{}{
		"type":    "JoinRoom",
		"payload": msg,
	}); err != nil {
		return err
	}

	return nil
}

// SendPointer implements InRoomUsecase.
func (i *activeRoomUsecase) SendPointer(ctx context.Context, req *SendPointerReq) error {
	room, err := i.repo.Get(ctx, req.RoomID)
	if err != nil {
		return err
	}

	room.Lock.RLock()
	defer room.Lock.RUnlock()

	user, err := i.uRepo.GetUser(ctx, req.UserID)
	if err != nil {
		return err
	}

	msg := &Pointer{
		User:      &User{ID: user.ID, Name: user.Name, AvatarURL: user.AvatarURL},
		IsClicked: req.IsClicked,
		X:         req.X,
		Y:         req.Y,
		Color:     req.Color,
		PenSize:   req.PenSize,
	}

	if err := i.msgSender.Send(ctx, room.OwnerID, map[string]interface{}{
		"type":    "ChangeUserPosition",
		"payload": []*Pointer{msg},
	}); err != nil {
		return err
	}

	return nil
}

func (i *activeRoomUsecase) ChangeScreenSize(ctx context.Context, roomID string, height, width int) error {
	room, err := i.repo.Get(ctx, roomID)
	if err != nil {
		return err
	}
	room.Lock.Lock()
	defer room.Lock.Unlock()

	room.ScreenHeight = height
	room.ScreenWidth = width
	if err := i.repo.Store(ctx, room); err != nil {
		return err
	}

	msg := map[string]interface{}{
		"type": "ChangeCurrentScreen",
		"payload": map[string]interface{}{
			"height": height,
			"width":  width,
		},
	}
	for _, user := range room.Users {
		if err := i.msgSender.Send(ctx, user.ID, msg); err != nil {
			return err
		}
	}
	return nil
}

func (i *activeRoomUsecase) LeaveRoom(ctx context.Context, roomID, userID string) error {
	room, err := i.repo.Get(ctx, roomID)
	if err != nil {
		return err
	}

	room.Lock.Lock()
	defer room.Lock.Unlock()

	for i, user := range room.Users {
		if user.ID == userID {
			room.Users = append(room.Users[:i], room.Users[i+1:]...)
			break
		}
	}
	if err := i.repo.Store(ctx, room); err != nil {
		return err
	}

	msg := map[string]interface{}{
		"type": "LeaveRoom",
		"payload": map[string]interface{}{
			"userId": userID,
		},
	}
	if err := i.msgSender.Send(ctx, room.OwnerID, msg); err != nil {
		return err
	}

	return nil
}
