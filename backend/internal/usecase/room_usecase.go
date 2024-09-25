package usecase

import (
	"context"
	"time"

	"github.com/tosaken1116/spino_cup_2024/backend/internal/domain/model"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/domain/repository"
)

type RoomDTO struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewRoomDTOFromModel(m *model.Room) *RoomDTO {
	return &RoomDTO{
		ID:          m.ID.String(),
		Name:        m.Name,
		Description: m.Description,
	}
}

type RoomUsecase interface {
	CreateRoom(ctx context.Context, dto *RoomDTO) (*RoomDTO, error)
	GetRoom(ctx context.Context, id string) (*RoomDTO, error)
	ListRoom(ctx context.Context) ([]*RoomDTO, error)
	UpdateRoom(ctx context.Context, dto *RoomDTO) (*RoomDTO, error)
}

type roomUsecase struct {
	repo repository.RoomRepository
}

func NewRoomUsecase(repo repository.RoomRepository) RoomUsecase {
	return &roomUsecase{repo: repo}
}

// CreateRoom implements RoomUsecase.
func (r *roomUsecase) CreateRoom(ctx context.Context, dto *RoomDTO) (*RoomDTO, error) {
	id, err := model.NewRoomID()
	if err != nil {
		return nil, err
	}

	room, err := model.NewRoom(id, dto.Name, dto.Description)
	if err != nil {
		return nil, err
	}

	if err := r.repo.CreateRoom(ctx, room); err != nil {
		return nil, err
	}

	return NewRoomDTOFromModel(room), nil
}

// GetRoom implements RoomUsecase.
func (r *roomUsecase) GetRoom(ctx context.Context, rawid string) (*RoomDTO, error) {
	id, err := model.ParseRoomID(rawid)
	if err != nil {
		return nil, model.ErrRoomIDInvalid
	}

	room, err := r.repo.GetRoom(ctx, id)
	if err != nil {
		return nil, err
	}

	return NewRoomDTOFromModel(room), nil
}

// ListRoom implements RoomUsecase.
func (r *roomUsecase) ListRoom(ctx context.Context) ([]*RoomDTO, error) {
	rooms, err := r.repo.ListRoom(ctx)
	if err != nil {
		return nil, err
	}

	roomDTOs := make([]*RoomDTO, 0, len(rooms))
	for _, room := range rooms {
		roomDTOs = append(roomDTOs, NewRoomDTOFromModel(room))
	}

	return roomDTOs, nil
}

// UpdateRoom implements RoomUsecase.
func (r *roomUsecase) UpdateRoom(ctx context.Context, dto *RoomDTO) (*RoomDTO, error) {
	id, err := model.ParseRoomID(dto.ID)
	if err != nil {
		return nil, model.ErrRoomIDInvalid
	}

	room, err := r.repo.GetRoom(ctx, id)
	if err != nil {
		return nil, err
	}

	room.Name = dto.Name
	room.Description = dto.Description
	if err := r.repo.UpdateRoom(ctx, room); err != nil {
		return nil, err
	}

	return NewRoomDTOFromModel(room), nil
}
