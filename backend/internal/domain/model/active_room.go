package model

import "sync"

type AcitveRoom struct {
	Room         *Room
	OwnerID      string
	ScreenHeight int
	ScreenWidth  int
	Lock         sync.RWMutex
	Users        []*User
}

func (a *AcitveRoom) AddUser(user *User) {
	a.Lock.Lock()
	defer a.Lock.Unlock()

	for _, u := range a.Users {
		if u.ID == user.ID {
			return
		}
	}

	a.Users = append(a.Users, user)
}
