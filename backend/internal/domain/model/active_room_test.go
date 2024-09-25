package model

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAcitveRoom_AddUser(t *testing.T) {
	t.Parallel()

	a := &AcitveRoom{
		Room:    &Room{},
		OwnerID: "",
		Lock:    sync.RWMutex{},
		Users:   []string{},
	}
	a.AddUser("user1")
	assert.Equal(t, 1, len(a.Users))

	a.AddUser("user1")
	assert.Equal(t, 1, len(a.Users))
}
