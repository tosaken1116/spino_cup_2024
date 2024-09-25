package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRoomID(t *testing.T) {
	roomID, err := NewRoomID()
	assert.NoError(t, err)
	assert.NotEmpty(t, roomID)
}
