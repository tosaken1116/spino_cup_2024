// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/tosaken1116/spino_cup_2024/backend/internal/app/config"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/app/container"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/handler"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/infra/db"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/infra/ws"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/router"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/usecase"
	"github.com/tosaken1116/spino_cup_2024/backend/pkg/auth"
	"github.com/tosaken1116/spino_cup_2024/backend/pkg/database"
)

// Injectors from wire.go:

func New() (*container.App, error) {
	dbConfig := config.NewDBConfig()
	databaseConfig := convertDBConfig(dbConfig)
	databaseDB, err := database.New(databaseConfig)
	if err != nil {
		return nil, err
	}
	roomRepository := db.NewRoomRepository(databaseDB)
	roomUsecase := usecase.NewRoomUsecase(roomRepository)
	roomHandler := handler.NewRoomHandler(roomUsecase)
	msgSender := ws.NewMsgSender()
	activeRoomRepo := db.NewActiveRoomRepository()
	userRepository := db.NewUserRepository(databaseDB)
	activeRoomUsecase := usecase.NewActiveRoomUsecase(msgSender, activeRoomRepo, roomRepository, userRepository)
	wsHandler := handler.NewWSHandler(activeRoomUsecase, msgSender)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)
	authClient, err := auth.New()
	if err != nil {
		return nil, err
	}
	echo := router.New(roomHandler, wsHandler, userHandler, authClient)
	app := container.New(echo, databaseDB)
	return app, nil
}

// wire.go:

func convertDBConfig(cfg *config.DBConfig) *database.Config {
	return &database.Config{
		DBHost:     cfg.DBHost,
		DBPort:     cfg.DBPort,
		DBName:     cfg.DBName,
		DBUser:     cfg.DBUser,
		DBPassword: cfg.DBPassword,
	}
}
