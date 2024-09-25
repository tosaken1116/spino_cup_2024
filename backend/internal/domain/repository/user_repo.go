package repository

import (
	"context"

	"github.com/tosaken1116/spino_cup_2024/backend/internal/domain/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, id string) (*model.User, error)
	GetUsers(ctx context.Context, ids []string) ([]*model.User, error)
}
