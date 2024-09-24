package model

import (
	"crypto/rand"
	"errors"
	"time"

	"github.com/oklog/ulid"
)

var ErrRoomNotFound = errors.New("room not found")
var ErrRoomIDInvalid = errors.New("room id is invalid")

type RoomID ulid.ULID
type Room struct {
	ID          RoomID
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewRoomID() (RoomID, error) {
	id, err := ulid.New(uint64(time.Now().UnixMilli()), rand.Reader)
	if err != nil {
		return RoomID{}, err
	}
	return RoomID(id), nil
}
func ParseRoomID(s string) (RoomID, error) {
	id, err := ulid.Parse(s)
	if err != nil {
		return RoomID{}, ErrRoomIDInvalid
	}
	return RoomID(id), nil
}

func (r RoomID) String() string {
	return ulid.ULID(r).String()
}

func NewRoom(name, description string) (*Room, error) {
	id, err := NewRoomID()
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

func NewRoomFromData(id, name, description string, createdAt, updatedAt time.Time) (*Room, error) {
	_id, err := ParseRoomID(id)
	if err != nil {
		return nil, err
	}

	return &Room{
		ID:          _id,
		Name:        name,
		Description: description,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}, nil
}
