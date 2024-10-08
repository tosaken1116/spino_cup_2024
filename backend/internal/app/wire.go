//go:build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/app/config"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/app/container"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/domain/service"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/handler"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/infra/db"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/infra/ws"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/router"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/usecase"
	"github.com/tosaken1116/spino_cup_2024/backend/pkg/auth"
	"github.com/tosaken1116/spino_cup_2024/backend/pkg/database"
)

func New() (*container.App, error) {
	wire.Build(
		// DB initialize
		config.NewDBConfig,
		convertDBConfig,
		database.New,

		// User
		auth.New,
		db.NewUserRepository,
		usecase.NewUserUsecase,
		handler.NewUserHandler,

		// Room
		db.NewRoomRepository,
		usecase.NewRoomUsecase,
		handler.NewRoomHandler,

		// Active Room
		ws.NewMsgSender,
		db.NewActiveRoomRepository,
		usecase.NewActiveRoomUsecase,
		handler.NewWSHandler,
		wire.Bind(new(service.MessageSender), new(*ws.MsgSender)),

		// Router
		router.New,

		// App initialize
		container.New,
	)

	return nil, nil
}

func convertDBConfig(cfg *config.DBConfig) *database.Config {
	return &database.Config{
		DBHost:     cfg.DBHost,
		DBPort:     cfg.DBPort,
		DBName:     cfg.DBName,
		DBUser:     cfg.DBUser,
		DBPassword: cfg.DBPassword,
	}
}
