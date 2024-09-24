package usecase

import (
	"context"

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
}

type Pointer struct {
	UserID    string  `json:"id"`
	IsClicked bool    `json:"isClicked"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Color     string  `json:"color"`
}

type JoinRoomResp struct {
	UserID       string  `json:"userId"`
	OwnerID      string  `json:"ownerId"`
	ScreenHeight float64 `json:"height"`
	ScreenWidth  float64 `json:"width"`
}

type InRoomUsecase interface {
	JoinRoom(ctx context.Context, userID, roomID string) error
	SendPointer(ctx context.Context, req *SendPointerReq) error
}

type inRoomUsecase struct {
	msgSender service.MessageSender
	repo      repository.ActiveRoomRepo
	rRepo     repository.RoomRepository
}

func NewActiveRoomUsecase(msgSender service.MessageSender, repo repository.ActiveRoomRepo, rRepo repository.RoomRepository) InRoomUsecase {
	return &inRoomUsecase{
		msgSender: msgSender,
		repo:      repo,
		rRepo:     rRepo,
	}
}

// JoinRoom implements InRoomUsecase.
func (i *inRoomUsecase) JoinRoom(ctx context.Context, userID, roomID string) error {
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
		}

		if err := i.repo.Store(ctx, room); err != nil {
			return err
		}
	}

	msg := &JoinRoomResp{
		UserID:       userID,
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
func (i *inRoomUsecase) SendPointer(ctx context.Context, req *SendPointerReq) error {
	room, err := i.repo.Get(ctx, req.RoomID)
	if err != nil {
		return err
	}

	msg := &Pointer{
		UserID:    req.UserID,
		IsClicked: req.IsClicked,
		X:         req.X,
		Y:         req.Y,
		Color:     req.Color,
	}

	if err := i.msgSender.Send(ctx, room.OwnerID, map[string]interface{}{
		"type":    "ChangeUserPosition",
		"payload": []*Pointer{msg},
	}); err != nil {
		return err
	}

	return nil
}
