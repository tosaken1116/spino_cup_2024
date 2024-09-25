package usecase

import (
	"context"

	"github.com/tosaken1116/spino_cup_2024/backend/internal/domain/model"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/domain/repository"
)

type UserDTO struct {
	ID        string
	Name      string
	AvatarURL string
}

func NewUserDTOFromEntity(user *model.User) *UserDTO {
	return &UserDTO{
		ID:        user.ID,
		Name:      user.Name,
		AvatarURL: user.AvatarURL,
	}
}

type UserUsecase interface {
	SignUp(ctx context.Context, id, name, avatarURL string) (*UserDTO, error)
}

type userUsecae struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecae{
		repo: repo,
	}
}

// SignUp implements UserUsecase.
func (u *userUsecae) SignUp(ctx context.Context, id string, name string, avatarURL string) (*UserDTO, error) {
	user := model.NewUser(id, name, avatarURL)

	if err := u.repo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return NewUserDTOFromEntity(user), nil
}
