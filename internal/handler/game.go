package handler

import (
	"errors"
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
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	iconFile, err := c.FormFile("icon")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to get icon file: "+err.Error())
	}
	req.Icon.InitFromMultipart(iconFile)

	imageFile, err := c.FormFile("image")
	if errors.Is(err, http.ErrMissingFile) {
		// 画像がない場合はエラーにしない
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get image file: "+err.Error())
	} else {
		req.Image.InitFromMultipart(imageFile)
	}

	newGameID := uuid.New()
	if err := h.repo.PostGame(newGameID, req); err != nil {
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
	panic("implement me")
}

func (h *Handler) GetGameIcon(c echo.Context, gameID models.GameIdInPath) error {
	//TODO implement me
	panic("implement me")
}
func (h *Handler) GetGameImage(c echo.Context, gameID models.GameIdInPath) error {
	panic("implement me")
}
