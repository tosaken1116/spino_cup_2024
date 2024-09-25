package auth

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

type AuthClient struct {
	client *auth.Client
}

func New() *AuthClient {
	app, err := firebase.NewApp(context.Background(), nil)

	client, err := app.Auth(context.Background())
	if err != nil {
		panic(err)
	}

	return &AuthClient{
		client: client,
	}
}
