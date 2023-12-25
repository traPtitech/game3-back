package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/traPtitech/game3-back/internal/api/models"
	"github.com/traPtitech/game3-back/internal/repository"
)

type Handler struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) Login(c echo.Context) error {
	panic("implement me")
}

func (h *Handler) Logout(c echo.Context) error {
	panic("implement me")
}

func (h *Handler) PostContacts(c echo.Context) error {
	panic("implement me")
}

func (h *Handler) GetEvents(c echo.Context) error {
	panic("implement me")
}

func (h *Handler) PostEvent(c echo.Context) error {
	panic("implement me")
}

func (h *Handler) GetCurrentEvent(c echo.Context) error {
	panic("implement me")
}

func (h *Handler) GetEvent(c echo.Context, eventId models.EventIdInPath) error {
	panic("implement me")
}

func (h *Handler) PatchEvent(c echo.Context, eventId models.EventIdInPath) error {
	panic("implement me")
}

func (h *Handler) GetEventCsv(c echo.Context, eventId models.EventIdInPath) error {
	panic("implement me")
}

func (h *Handler) GetEventGames(c echo.Context, eventId models.EventIdInPath) error {
	panic("implement me")
}

func (h *Handler) GetEventImage(c echo.Context, eventId models.EventIdInPath) error {
	panic("implement me")
}

func (h *Handler) PostGame(ctx echo.Context) error {
	panic("implement me")
}

func (h *Handler) GetGame(ctx echo.Context, gameId models.GameIdInPath) error {
	panic("implement me")
}

func (h *Handler) PatchGame(ctx echo.Context, gameId models.GameIdInPath) error {
	panic("implement me")
}

func (h *Handler) GetGameImage(ctx echo.Context, gameId models.GameIdInPath) error {
	panic("implement me")
}

func (h *Handler) GetMe(ctx echo.Context) error {
	panic("implement me")
}

func (h *Handler) GetMeGames(ctx echo.Context) error {
	panic("implement me")
}

func (h *Handler) GetUser(ctx echo.Context, userId models.UserIdInPath) error {
	panic("implement me")
}

func (h *Handler) GetUserGames(ctx echo.Context, userId models.UserIdInPath) error {
	panic("implement me")
}
