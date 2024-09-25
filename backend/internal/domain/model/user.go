package model

type User struct {
	ID        string
	Name      string
	AvatarURL string
}

func NewUser(id, name, avatarURL string) *User {
	return &User{
		ID:        id,
		Name:      name,
		AvatarURL: avatarURL,
	}
}
