package main

import (
	"context"
	"log"

	"github.com/tosaken1116/spino_cup_2024/backend/internal/app"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Printf("failed to create server: %v\n", err)
		return
	}
	defer func() {
		if err := app.Close(); err != nil {
			log.Printf("failed to close server: %v\n", err)
		}
	}()

	if err := app.Migrate(context.Background()); err != nil {
		log.Fatalf("failed to migrate: %v\n", err)
	}

	if err := app.Start(); err != nil {
		log.Printf("failed to start server: %v\n", err)
		return
	}
}
