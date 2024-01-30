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

func (h *Handler) PostGame(ctx echo.Context) error {
	req := &models.PostGameRequest{}
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	newGameId, err := h.repo.PostGame(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	game, err := h.repo.GetGame(newGameId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, game)
}

func (h *Handler) GetGame(ctx echo.Context, gameId models.GameIdInPath) error {
	gameUUID, err := uuid.Parse(gameId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	game, err := h.repo.GetGame(gameUUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, game)
}

func (h *Handler) PatchGame(ctx echo.Context, gameId models.GameIdInPath) error {
	panic("implement me")
}

func (h *Handler) GetGameIcon(ctx echo.Context, gameId models.GameIdInPath) error {
	//TODO implement me
	panic("implement me")
}
func (h *Handler) GetGameImage(ctx echo.Context, gameId models.GameIdInPath) error {
	panic("implement me")
}
