package db

import (
	"context"
	"errors"
	"sync"

	"github.com/tosaken1116/spino_cup_2024/backend/internal/domain/model"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/domain/repository"
)

type activeRepoImpl struct {
	db *sync.Map
}

func NewActiveRoomRepository() repository.ActiveRoomRepo {
	return &activeRepoImpl{
		db: &sync.Map{},
	}
}

// Get implements repository.ActiveRoomRepo.
func (a *activeRepoImpl) Get(ctx context.Context, roomID string) (*model.AcitveRoom, error) {
	rawRoom, ok := a.db.Load(roomID)
	if !ok {
		return nil, errors.New("room not found")
	}

	room, ok := rawRoom.(*model.AcitveRoom)
	if !ok {
		return nil, errors.New("invalid room")
	}

	return room, nil
}

// Store implements repository.ActiveRoomRepo.
func (a *activeRepoImpl) Store(ctx context.Context, room *model.AcitveRoom) error {
	a.db.Store(room.Room.ID.String(), room)
	return nil
}
