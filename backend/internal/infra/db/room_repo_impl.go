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
	OwnerID     string
	Name        string
	Description string
	OwnerID     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type roomRepoImpl struct {
	db *database.DB
}

func NewRoomRepository(db *database.DB) repository.RoomRepository {
	return &roomRepoImpl{db: db}
}

// CreateRoom implements repository.RoomRepository.
func (r *roomRepoImpl) CreateRoom(ctx context.Context, room *model.Room) error {
	now := time.Now()
	data := &roomModel{
		ID:          room.ID.String(),
		Name:        room.Name,
		Description: room.Description,
		OwnerID:     room.OwnerID,
		CreatedAt:   now,
		UpdatedAt:   now,
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
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create room: %w", err)
	}

	return room, nil
}

// ListRoom implements repository.RoomRepository.
func (r *roomRepoImpl) ListRoom(ctx context.Context) ([]*model.Room, error) {
	var data []roomModel
	query := r.db.NewSelect().Model(&data).Order("created_at DESC").Limit(100)

	if err := query.Scan(ctx); err != nil {
		return nil, fmt.Errorf("failed to get rooms: %w", err)
	}

	rooms := make([]*model.Room, 0, len(data))
	for _, room := range data {
		room, err := model.NewRoomFromData(
			room.ID,
			room.Name,
			room.Description,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to create room: %w", err)
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

// UpdateRoom implements repository.RoomRepository.
func (r *roomRepoImpl) UpdateRoom(ctx context.Context, room *model.Room) error {
	data := &roomModel{
		ID:          room.ID.String(),
		Name:        room.Name,
		Description: room.Description,
		UpdatedAt:   time.Now(),
	}

	stmt := r.db.NewUpdate().Model(data).OmitZero().Where("id = ?", room.ID.String())
	if _, err := stmt.Exec(ctx); err != nil {
		return fmt.Errorf("failed to update room: %w", err)
	}
	return nil
}
