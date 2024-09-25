package repository

import (
	"context"

	"github.com/tosaken1116/spino_cup_2024/backend/internal/domain/model"
)

type RoomRepository interface {
	CreateRoom(ctx context.Context, room *model.Room) error
	GetRoom(ctx context.Context, id model.RoomID) (*model.Room, error)
	ListRoom(ctx context.Context) ([]*model.Room, error)
	UpdateRoom(ctx context.Context, room *model.Room) error
}
