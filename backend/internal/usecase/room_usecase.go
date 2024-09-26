package usecase

import (
	"context"

	"github.com/tosaken1116/spino_cup_2024/backend/internal/domain/model"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/domain/repository"
)

type RoomDTO struct {
	ID          string
	Name        string
	Description string
	OwnerID     string
	Owner       *User
}

func NewRoomDTOFromModel(m *model.Room) *RoomDTO {
	return &RoomDTO{
		ID:          m.ID.String(),
		Name:        m.Name,
		Description: m.Description,
		OwnerID:     m.OwnerID,
	}
}

type RoomUsecase interface {
	CreateRoom(ctx context.Context, dto *RoomDTO) (*RoomDTO, error)
	GetRoom(ctx context.Context, id string) (*RoomDTO, error)
	ListRoom(ctx context.Context) ([]*RoomDTO, error)
	UpdateRoom(ctx context.Context, dto *RoomDTO) (*RoomDTO, error)
}

type roomUsecase struct {
	repo  repository.RoomRepository
	uRepo repository.UserRepository
}

func NewRoomUsecase(repo repository.RoomRepository, uRepo repository.UserRepository) RoomUsecase {
	return &roomUsecase{repo: repo, uRepo: uRepo}
}

// CreateRoom implements RoomUsecase.
func (r *roomUsecase) CreateRoom(ctx context.Context, dto *RoomDTO) (*RoomDTO, error) {
	id, err := model.NewRoomID()
	if err != nil {
		return nil, err
	}

	room, err := model.NewRoom(id, dto.Name, dto.Description, dto.OwnerID)
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

	userIDs := make([]string, 0, len(rooms))
	roomDTOs := make([]*RoomDTO, 0, len(rooms))
	for _, room := range rooms {
		roomDTOs = append(roomDTOs, NewRoomDTOFromModel(room))
		userIDs = append(userIDs, room.OwnerID)
	}

	// Get users
	userIDmap := make(map[string]*model.User)
	users, err := r.uRepo.GetUsers(ctx, userIDs)
	for _, user := range users {
		userIDmap[user.ID] = user
	}

	// Set user to room
	for _, room := range roomDTOs {
		if user, ok := userIDmap[room.OwnerID]; ok {
			room.Owner = &User{
				ID:        user.ID,
				Name:      user.Name,
				AvatarURL: user.AvatarURL,
			}
		}
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
