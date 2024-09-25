package auth

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

var ErrUserIDNotFound = errors.New("user ID not found")

func (a *AuthClient) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		if auth == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header is required")
		}

		auth = strings.TrimPrefix(auth, "Bearer ")
		if auth == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token must be Bearer")
		}

		ctx := c.Request().Context()
		idToken, err := a.client.VerifyIDToken(ctx, auth)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token").SetInternal(err)
		}

		ctx = SetUserID(ctx, idToken.UID)
		req := c.Request().Clone(ctx)
		c.SetRequest(req)

		return next(c)
	}
}

var authKey = struct{}{}

func SetUserID(c context.Context, uid string) context.Context {
	return context.WithValue(c, authKey, uid)
}

func GetUserID(c context.Context) (string, error) {
	if id, ok := c.Value(authKey).(string); ok {
		return id, nil
	}
	return "", ErrUserIDNotFound
}
