package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime/types"
	"github.com/traPtitech/game3-back/internal/repository"
	"net/http"
)

type Handler struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Handler {
	return &Handler{
		repo: repo,
	}
}

// handleFile processes a file from the form data and returns a *types.File.
func handleFile(c echo.Context, formFileName string) (*types.File, error) {
	fileHeader, err := c.FormFile(formFileName)
	if err != nil {
		if errors.Is(err, http.ErrMissingFile) {
			return nil, nil
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to get file: "+err.Error())
	}

	file := types.File{}
	file.InitFromMultipart(fileHeader)

	return &file, nil
}

// getSessionIdByCookie gets the session ID from the SessionToken cookie.
func getSessionIDByCookie(c echo.Context) (*uuid.UUID, error) {
	sessionByCookie, err := c.Cookie("SessionToken")
	if err != nil {
		return nil, err
	}
	sessionID, err := uuid.Parse(sessionByCookie.Value)
	if err != nil {
		return nil, err
	}

	return &sessionID, nil
}

func (h *Handler) getDiscordUserInfoByCookie(c echo.Context) (*api.DiscordUserResponse, error) {
	sessionID, err := getSessionIDByCookie(c)
	if err != nil {
		return nil, err
	}

	session, err := h.repo.GetSession(sessionID)
	if err != nil {
		return nil, err
	}

	return api.GetDiscordUserInfo(session.AccessToken)
}
