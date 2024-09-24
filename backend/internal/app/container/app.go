package container

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/tosaken1116/spino_cup_2024/backend/pkg/database"
)

type App struct {
	echo *echo.Echo
	db   *database.DB
}

func New(e *echo.Echo, db *database.DB) *App {
	return &App{
		echo: e,
		db:   db,
	}
}

func (a *App) Start() error {
	return a.echo.Start(":8080")
}

func (a *App) Close() error {
	return errors.Join(
		a.db.Close(),
	)
}
