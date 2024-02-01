package handler

import (
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime/types"
	"github.com/traPtitech/game3-back/internal/api"
	"github.com/traPtitech/game3-back/internal/apperrors"
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

// respondWithError returns an error response based on the error type.
func respondWithError(c echo.Context, err error) error {
	//TODO err.Error()をそのまま返すのはセキュリティ的に問題があるので、エラータイプによって適切なエラーメッセージを返すように修正する
	var notFoundError *apperrors.NotFoundError
	var badRequestError *apperrors.BadRequestError
	var forbiddenError *apperrors.ForbiddenError
	var unauthorizedError *apperrors.UnauthorizedError
	var internalServerError *apperrors.InternalServerError
	switch {
	case errors.As(err, &notFoundError):
		return c.String(http.StatusNotFound, err.Error())
	case errors.As(err, &badRequestError):
		return c.String(http.StatusBadRequest, err.Error())
	case errors.As(err, &forbiddenError):
		return c.String(http.StatusForbidden, err.Error())
	case errors.As(err, &unauthorizedError):
		return c.String(http.StatusUnauthorized, err.Error())
	case errors.As(err, &internalServerError):
		return c.String(http.StatusInternalServerError, err.Error())
	default:
		return c.String(http.StatusInternalServerError, err.Error())
	}
}

// handleFile processes a file from the form data and returns a *types.File.
func handleFile(c echo.Context, formFileName string) (*types.File, error) {
	fileHeader, err := c.FormFile(formFileName)
	if err != nil {
		if errors.Is(err, http.ErrMissingFile) {
			return nil, nil
		}
		return nil, err
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
