package model

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid"
)

type RoomID ulid.ULID
type Room struct {
	ID          RoomID
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewRoom(name, description string) (*Room, error) {
	id, err := ulid.New(uint64(time.Now().UnixMilli()), rand.Reader)
	if err != nil {
		return nil, err
	}

	return &Room{
		ID:          RoomID(id),
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
