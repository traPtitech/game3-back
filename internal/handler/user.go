package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traPtitech/game3-back/internal/api"
	"github.com/traPtitech/game3-back/openapi/models"
	"net/http"
)

func (h *Handler) GetMe(c echo.Context) error {
	discordUser, err := h.getDiscordUserInfoByCookie(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, discordUser)
}

func (h *Handler) GetMeGames(c echo.Context) error {
	cookie, err := c.Cookie("SessionToken")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "SessionToken is not found")
	}
	sessionID, err := uuid.Parse(cookie.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "SessionToken is invalid")
	}

	session, err := h.repo.GetSession(&sessionID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	discordUser, err := api.GetDiscordUserInfo(session.AccessToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	games, err := h.repo.GetGames(models.GetGamesParams{UserId: &discordUser.ID})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, games)
}

func (h *Handler) GetUser(c echo.Context, userId models.UserIdInPath) error {
	return c.JSON(http.StatusNotImplemented, nil)
}

func (h *Handler) GetUserGames(c echo.Context, userId models.UserIdInPath) error {
	games, err := h.repo.GetGames(models.GetGamesParams{UserId: &userId})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, games)
}
