package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/tosaken1116/spino_cup_2024/backend/pkg/database"
)

func main() {
	db, err := database.New(&database.Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		panic(err)
	}
	defer db.Close()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
