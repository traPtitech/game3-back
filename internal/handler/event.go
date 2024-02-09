package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime/types"
	"github.com/traPtitech/game3-back/internal/apperrors"
	"github.com/traPtitech/game3-back/internal/pkg/constants"
	"github.com/traPtitech/game3-back/openapi/models"
	"net/http"
)

func (h *Handler) PostEvent(c echo.Context) (err error) {
	if _, _, err = h.enforceAdminAccess(c); err != nil {
		return err
	}

	req := &models.PostEventRequest{}
	if err = c.Bind(req); err != nil {
		return apperrors.HandleBindError(err)
	}

	req.Image, err = handleEventImage(c)
	if err != nil {
		return apperrors.HandleFileError(err)
	}

	if err = h.repo.PostEvent(req); err != nil {
		return apperrors.HandleDbError(err)
	}

	event, err := h.repo.GetEvent(req.Slug)
	if err != nil {
		return apperrors.HandleDbError(err)
	}

	if err = h.repo.CreateDefaultTerm(event.Slug); err != nil {
		return apperrors.HandleDbError(err)
	}

	return c.JSON(http.StatusCreated, event)
}

func (h *Handler) GetEvents(c echo.Context) error {
	events, err := h.repo.GetEvents()
	if err != nil {
		return apperrors.HandleDbError(err)
	}

	return c.JSON(http.StatusOK, events)
}

func (h *Handler) GetCurrentEvent(c echo.Context) error {
	event, err := h.repo.GetCurrentEvent()
	if err != nil {
		return apperrors.HandleDbError(err)
	}

	return c.JSON(http.StatusOK, event)
}

func (h *Handler) PatchEvent(c echo.Context, eventID models.EventSlugInPath) (err error) {
	if _, _, err = h.enforceAdminAccess(c); err != nil {
		return err
	}

	req := &models.PatchEventRequest{}
	if err = c.Bind(req); err != nil {
		return apperrors.HandleBindError(err)
	}

	req.Image, err = handleEventImage(c)
	if err != nil {
		return apperrors.HandleFileError(err)
	}

	if err = h.repo.PatchEvent(eventID, req); err != nil {
		return apperrors.HandleDbError(err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) GetEvent(c echo.Context, eventSlug models.EventSlugInPath) error {
	event, err := h.repo.GetEvent(eventSlug)
	if err != nil {
		return apperrors.HandleDbError(err)
	}

	return c.JSON(http.StatusOK, event)
}

func (h *Handler) GetEventImage(c echo.Context, eventID models.EventSlugInPath) error {
	image, err := h.repo.GetEventImage(eventID)
	if err != nil {
		return apperrors.HandleDbError(err)
	}

	return c.Blob(http.StatusOK, "image/png", image)
}

func (h *Handler) GetEventTerms(c echo.Context, eventID models.EventSlugInPath) error {
	events, err := h.repo.GetEventTerms(eventID)
	if err != nil {
		return apperrors.HandleDbError(err)
	}

	return c.JSON(http.StatusOK, events)
}

func (h *Handler) GetEventGames(c echo.Context, eventID models.EventSlugInPath) error {
	games, err := h.repo.GetGames(models.GetGamesParams{EventSlug: &eventID})
	if err != nil {
		return apperrors.HandleDbError(err)
	}

	return c.JSON(http.StatusOK, games)
}

func (h *Handler) GetEventCsv(c echo.Context, eventID models.EventSlugInPath) error {
	if _, _, err := h.enforceAdminAccess(c); err != nil {
		return err
	}

	return echo.NewHTTPError(http.StatusNotImplemented, "not implemented")
}

func handleEventImage(c echo.Context) (*types.File, error) {
	return handleImageFileAndConvertImageToPNGAndResizeImage(c, "image", constants.EventImageWidth, constants.EventImageHeight)
}
