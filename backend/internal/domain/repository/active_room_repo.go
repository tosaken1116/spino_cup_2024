package repository

import (
	"context"

	"github.com/tosaken1116/spino_cup_2024/backend/internal/domain/model"
)

type ActiveRoomRepo interface {
	Get(ctx context.Context, roomID string) (*model.AcitveRoom, error)
	Store(ctx context.Context, room *model.AcitveRoom) error
}
