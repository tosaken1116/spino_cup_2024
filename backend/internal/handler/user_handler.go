package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/usecase"
)

type UserHandler interface {
	SignUp(c echo.Context) error
}

type userHandler struct {
	uc usecase.UserUsecase
}

func NewUserHandler(uc usecase.UserUsecase) UserHandler {
	return &userHandler{uc}
}

type SignUpRequest struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatarUrl"`
}

// SignUp implements UserHandler.
func (u *userHandler) SignUp(c echo.Context) error {
	var req SignUpRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err).SetInternal(err)
	}

	ctx := c.Request().Context()
	if _, err := u.uc.SignUp(ctx, req.ID, req.Name, req.AvatarURL); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	return c.NoContent(http.StatusCreated)
}
