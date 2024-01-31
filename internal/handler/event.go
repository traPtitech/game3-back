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

func (h *Handler) PostEvent(c echo.Context) (err error) {
	req := &models.PostEventRequest{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "test "+err.Error())
	}

	req.Image, err = h.handleFile(c, "image")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get image file: "+err.Error())
	}

	if err := h.repo.PostEvent(req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	event, err := h.repo.GetEvent(req.Slug)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := h.repo.CreateDefaultTerm(event.Slug); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, event)
}

func (h *Handler) GetCurrentEvent(c echo.Context) error {
	event, err := h.repo.GetCurrentEvent()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, event)
}

func (h *Handler) GetEvent(c echo.Context, eventSlug models.EventSlugInPath) error {
	event, err := h.repo.GetEvent(eventSlug)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, event)
}

func (h *Handler) PatchEvent(c echo.Context, eventID models.EventSlugInPath) (err error) {
	req := &models.PatchEventRequest{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	req.Image, err = h.handleFile(c, "image")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get image file: "+err.Error())
	}

	if err := h.repo.PatchEvent(eventID, req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) GetEventCsv(c echo.Context, eventID models.EventSlugInPath) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "not implemented")
}
func (h *Handler) GetEventTerms(ctx echo.Context, eventID models.EventSlugInPath) error {
	events, err := h.repo.GetEventTerms(eventID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, events)
}
func (h *Handler) GetEventGames(c echo.Context, eventID models.EventSlugInPath) error {
	games, err := h.repo.GetEventGames(eventID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, games)
}

func (h *Handler) GetEventImage(c echo.Context, eventID models.EventSlugInPath) error {
	image, err := h.repo.GetEventImage(eventID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.Blob(http.StatusOK, "image/png", image)
}
