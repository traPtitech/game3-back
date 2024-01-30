package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/traPtitech/game3-back/internal/api/models"
)

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
func (h *Handler) GetEventTerms(ctx echo.Context, eventId models.EventIdInPath) error {
	//TODO implement me
	panic("implement me")
}
func (h *Handler) GetEventGames(c echo.Context, eventId models.EventIdInPath) error {
	panic("implement me")
}

func (h *Handler) GetEventImage(c echo.Context, eventId models.EventIdInPath) error {
	panic("implement me")
}
