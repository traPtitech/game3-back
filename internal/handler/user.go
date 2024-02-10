package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/traPtitech/game3-back/internal/apperrors"
	"github.com/traPtitech/game3-back/openapi/models"
	"net/http"
)

func (h *Handler) GetMe(c echo.Context) error {
	user, role, err := h.enforceUserOrAboveAccess(c)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, models.User{
		ProfileImageUrl: fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s", user.ID, user.Avatar),
		Role:            role.ToModelsUserRole(),
		UserId:          user.ID,
		Username:        user.Username,
	})
}

func (h *Handler) GetMeGames(c echo.Context) error {
	user, _, err := h.enforceUserOrAboveAccess(c)
	if err != nil {
		return err
	}

	includeUnpublished := true
	games, err := h.repo.GetGames(models.GetGamesParams{UserId: &user.ID, IncludeUnpublished: &includeUnpublished})
	if err != nil {
		return apperrors.HandleDbError(err)
	}

	return c.JSON(http.StatusOK, games)
}

func (h *Handler) GetUser(c echo.Context, _ models.UserIdInPath) error {
	return c.JSON(http.StatusNotImplemented, nil)
}

func (h *Handler) GetUserGames(c echo.Context, userID models.UserIdInPath) error {
	games, err := h.repo.GetGames(models.GetGamesParams{UserId: &userID})
	if err != nil {
		return apperrors.HandleDbError(err)
	}

	return c.JSON(http.StatusOK, games)
}
