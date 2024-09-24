package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/tosaken1116/spino_cup_2024/backend/internal/domain/model"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/domain/repository"
	"github.com/tosaken1116/spino_cup_2024/backend/pkg/database"
	"github.com/uptrace/bun"
)

type roomModel struct {
	bun.BaseModel `bun:"rooms"`

	ID          string `bun:",pk"`
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type roomRepoImpl struct {
	db *database.DB
}

func NewRoomRepositroy(db *database.DB) repository.RoomRepository {
	return &roomRepoImpl{db: db}
}

// CreateRoom implements repository.RoomRepository.
func (r *roomRepoImpl) CreateRoom(ctx context.Context, room *model.Room) error {
	data := &roomModel{
		ID:          room.ID.String(),
		Name:        room.Name,
		Description: room.Description,
		CreatedAt:   room.CreatedAt,
		UpdatedAt:   room.UpdatedAt,
	}

	stmt := r.db.NewInsert().Model(data)
	if _, err := stmt.Exec(ctx); err != nil {
		return fmt.Errorf("failed to insert room: %w", err)
	}

	return nil
}

// GetRoom implements repository.RoomRepository.
func (r *roomRepoImpl) GetRoom(ctx context.Context, id model.RoomID) (*model.Room, error) {
	var data roomModel
	query := r.db.NewSelect().Model(&data).Where("id = ?", id.String()).Limit(1)

	if err := query.Scan(ctx); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, model.ErrRoomNotFound
		default:
			return nil, fmt.Errorf("failed to get room: %w", err)
		}
	}

	room, err := model.NewRoomFromData(
		data.ID,
		data.Name,
		data.Description,
		data.CreatedAt,
		data.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create room: %w", err)
	}

	return room, nil
}

// UpdateRoom implements repository.RoomRepository.
func (r *roomRepoImpl) UpdateRoom(ctx context.Context, room *model.Room) error {
	data := &roomModel{
		ID:          room.ID.String(),
		Name:        room.Name,
		Description: room.Description,
		CreatedAt:   room.CreatedAt,
		UpdatedAt:   room.UpdatedAt,
	}

	stmt := r.db.NewUpdate().Model(data).Where("id = ?", room.ID.String())
	if _, err := stmt.Exec(ctx); err != nil {
		return fmt.Errorf("failed to update room: %w", err)
	}
	return nil
}
