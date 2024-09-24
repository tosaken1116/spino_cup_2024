package model

import "sync"

type AcitveRoom struct {
	Room         *Room
	OwnerID      string
	ScreenHeight int
	ScreenWidth  int
	Lock         sync.RWMutex
	Users        []string
}

func (a *AcitveRoom) AddUser(userID string) {
	a.Lock.Lock()
	defer a.Lock.Unlock()

	for _, u := range a.Users {
		if u == userID {
			return
		}
	}
	a.Users = append(a.Users, userID)
}
