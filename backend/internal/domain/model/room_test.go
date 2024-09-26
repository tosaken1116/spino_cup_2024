package model

import (
	"testing"

	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

func TestNewRoomID(t *testing.T) {
	roomID, err := NewRoomID()
	assert.NoError(t, err)
	assert.NotEmpty(t, roomID)
}

func TestParseRoomID(t *testing.T) {
	t.Parallel()

	id := ulid.MustNew(ulid.Now(), nil)
	type args struct {
		s string
	}
	tests := []struct {
		name      string
		args      args
		want      RoomID
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				s: id.String(),
			},
			want:      RoomID(id),
			assertion: assert.NoError,
		},
		{
			name: "failed",
			args: args{
				s: "invalid",
			},
			want: RoomID{},
			assertion: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
				return assert.Error(t, err, ErrRoomIDInvalid)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseRoomID(tt.args.s)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRoomID_String(t *testing.T) {
	t.Parallel()

	id := ulid.MustNew(ulid.Now(), nil)
	tests := []struct {
		name string
		r    RoomID
		want string
	}{
		{
			name: "success",
			r:    RoomID(id),
			want: id.String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.r.String())
		})
	}
}

func TestNewRoom(t *testing.T) {
	t.Parallel()

	id, _ := NewRoomID()
	type args struct {
		id          RoomID
		name        string
		description string
		ownerID     string
	}
	tests := []struct {
		name      string
		args      args
		want      *Room
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				id:          id,
				name:        "room1",
				description: "room1 description",
				ownerID:     "user1",
			},
			want: &Room{
				ID:          id,
				Name:        "room1",
				Description: "room1 description",
				OwnerID:     "user1",
			},
			assertion: assert.NoError,
		},
		{
			name: "filed",
			args: args{
				id:          id,
				name:        "",
				description: "",
				ownerID:     "",
			},
			want: nil,
			assertion: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
				return assert.ErrorIs(t, err, ErrRoomNameRequired)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewRoom(tt.args.id, tt.args.name, tt.args.description, tt.args.ownerID)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNewRoomFromData(t *testing.T) {
	t.Parallel()

	id := ulid.MustNew(ulid.Now(), nil)
	type args struct {
		id          string
		name        string
		description string
		ownerID     string
	}
	tests := []struct {
		name      string
		args      args
		want      *Room
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				id:          id.String(),
				name:        "room1",
				description: "room1 description",
				ownerID:     "user1",
			},
			want: &Room{
				ID:          RoomID(id),
				Name:        "room1",
				Description: "room1 description",
				OwnerID:     "user1",
			},
			assertion: assert.NoError,
		},
		{
			name: "failed",
			args: args{
				id:          "invalid id",
				name:        "room1",
				description: "room1 description",
				ownerID:     "user1",
			},
			want: nil,
			assertion: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
				return assert.ErrorIs(t, err, ErrRoomIDInvalid)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewRoomFromData(tt.args.id, tt.args.name, tt.args.description)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
