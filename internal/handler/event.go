package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/traPtitech/game3-back/internal/api/models"
	"net/http"
)

func (h *Handler) GetEvents(c echo.Context) error {
	events, err := h.repo.GetEvents()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, events)
}

func (h *Handler) PostEvent(c echo.Context) error {
	req := &models.PostEventRequest{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.repo.PostEvent(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, nil)
}

func (h *Handler) GetCurrentEvent(c echo.Context) error {
	event, err := h.repo.GetCurrentEvent()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, event)
}

func (h *Handler) GetEvent(c echo.Context, eventId models.EventIdInPath) error {
	event, err := h.repo.GetEvent(eventId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, event)
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
