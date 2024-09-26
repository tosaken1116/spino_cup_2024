package model

import (
	"crypto/rand"
	"errors"
	"time"

	"github.com/oklog/ulid"
)

var ErrRoomNotFound = errors.New("room not found")
var ErrRoomIDInvalid = errors.New("room id is invalid")
var ErrRoomNameRequired = errors.New("room name is required")

type RoomID ulid.ULID
type Room struct {
	ID          RoomID
	Name        string
	Description string
	OwnerID     string
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

func NewRoom(id RoomID, name, description, ownerID string) (*Room, error) {
	if name == "" {
		return nil, ErrRoomNameRequired
	}

	return &Room{
		ID:          id,
		Name:        name,
		Description: description,
		OwnerID:     ownerID,
	}, nil
}

func NewRoomFromData(id, name, description, ownerID string) (*Room, error) {
	_id, err := ParseRoomID(id)
	if err != nil {
		return nil, err
	}

	return &Room{
		ID:          _id,
		Name:        name,
		Description: description,
		OwnerID:     ownerID,
	}, nil
}
