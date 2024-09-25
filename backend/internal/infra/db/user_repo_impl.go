package db

import (
	"context"

	"github.com/tosaken1116/spino_cup_2024/backend/internal/domain/model"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/domain/repository"
	"github.com/tosaken1116/spino_cup_2024/backend/pkg/database"
	"github.com/uptrace/bun"
)

type userModel struct {
	bun.BaseModel `bun:"users"`

	ID        string `bun:",pk"`
	Name      string
	AvatarURL string
}

type userRepoImpl struct {
	db *database.DB
}

func NewUserRepository(db *database.DB) repository.UserRepository {
	return &userRepoImpl{
		db: db,
	}
}

// CreateUser implements repository.UserRepository.
func (u *userRepoImpl) CreateUser(ctx context.Context, user *model.User) error {
	data := &userModel{
		ID:        user.ID,
		Name:      user.Name,
		AvatarURL: user.AvatarURL,
	}

	if _, err := u.db.NewInsert().Model(data).Exec(ctx); err != nil {
		return err
	}

	return nil
}

// GetUser implements repository.UserRepository.
func (u *userRepoImpl) GetUser(ctx context.Context, id string) (*model.User, error) {
	var data userModel
	query := u.db.NewSelect().Model(&data).Where("id = ?", id)

	if err := query.Scan(ctx); err != nil {
		return nil, err
	}

	return model.NewUser(data.ID, data.Name, data.AvatarURL), nil
}

// GetUsers implements repository.UserRepository.
func (u *userRepoImpl) GetUsers(ctx context.Context, ids []string) ([]*model.User, error) {
	var data []*userModel
	query := u.db.NewSelect().Model(&data).Where("id IN (?)", bun.In(ids))
	if err := query.Scan(ctx); err != nil {
		return nil, err
	}

	users := make([]*model.User, 0, len(data))
	for _, d := range data {
		users = append(users, model.NewUser(d.ID, d.Name, d.AvatarURL))
	}

	return users, nil
}
