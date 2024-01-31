package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traPtitech/game3-back/internal/api/models"
	"net/http"
)

func (h *Handler) GetGames(c echo.Context, params models.GetGamesParams) error {
	games, err := h.repo.GetGames(params)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, games)
}

func (h *Handler) PostGame(c echo.Context) error {
	req := &models.PostGameRequest{}
	err := c.Bind(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	icon, err := h.handleFile(c, "icon")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get image file: "+err.Error())
	}
	req.Icon = *icon
	req.Image, err = h.handleFile(c, "image")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get image file: "+err.Error())
	}

	newGameID := uuid.New()
	if err := h.repo.PostGame(newGameID, req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	event, err := h.repo.GetCurrentEvent()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	term, err := h.repo.GetDefaultTerm(event.Slug)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err := h.repo.PatchGame(newGameID, &models.PatchGameRequest{TermId: term.Id}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	
	game, err := h.repo.GetGame(newGameID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, game)
}

func (h *Handler) GetGame(c echo.Context, gameID models.GameIdInPath) error {
	game, err := h.repo.GetGame(gameID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, game)
}

func (h *Handler) PatchGame(c echo.Context, gameID models.GameIdInPath) error {
	req := &models.PatchGameRequest{}
	err := c.Bind(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	req.Icon, err = h.handleFile(c, "icon")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get image file: "+err.Error())
	}
	req.Image, err = h.handleFile(c, "image")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get image file: "+err.Error())
	}

	if err := h.repo.PatchGame(gameID, req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	game, err := h.repo.GetGame(gameID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, game)
}

func (h *Handler) GetGameIcon(c echo.Context, gameID models.GameIdInPath) error {
	icon, err := h.repo.GetGameIcon(gameID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.Blob(http.StatusOK, "image/png", icon)
}
func (h *Handler) GetGameImage(c echo.Context, gameID models.GameIdInPath) error {
	image, err := h.repo.GetGameImage(gameID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.Blob(http.StatusOK, "image/png", image)
}
