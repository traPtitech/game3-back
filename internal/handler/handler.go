package handler

import (
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime/types"
	"github.com/traPtitech/game3-back/internal/api"
	"github.com/traPtitech/game3-back/internal/apperrors"
	"github.com/traPtitech/game3-back/internal/enum"
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

// respondWithError returns an error response based on the error type. (使わなかったので消すかも)
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
		if errors.As(err, &http.ErrNoCookie) {
			return nil, apperrors.NewSessionTokenNotFoundError()
		}
		return nil, err
	}
	if sessionByCookie == nil {
		return nil, apperrors.NewSessionTokenNotFoundError()
	}
	sessionID, err := uuid.Parse(sessionByCookie.Value)
	if err != nil {
		return nil, err
	}

	return &sessionID, nil
}

func (h *Handler) getDiscordUserInfoByCookie(c echo.Context) (*api.GetDiscordUserInfoResponse, error) {
	sessionID, err := getSessionIDByCookie(c)
	if err != nil {
		return nil, err
	}

	session, err := h.repo.GetSession(sessionID)
	if err != nil {
		return nil, err
	}
	if session.AccessToken == nil {
		return nil, errors.New("session.AccessToken is not valid")
	}

	return api.GetDiscordUserInfo(session.AccessToken)
}

func (h *Handler) getDiscordUserInfoAndRoleByCookie(c echo.Context) (*api.GetDiscordUserInfoResponse, enum.UserRole, error) {
	user, err := h.getDiscordUserInfoByCookie(c)

	var notFoundErr *apperrors.SessionTokenNotFoundError
	if errors.As(err, &notFoundErr) {
		return nil, enum.Guest, nil
	}
	if err != nil {
		return nil, enum.Guest, err
	}

	role := api.GetDiscordUserRole(user.ID)

	return user, role, nil
}

func (h *Handler) enforceAdminAccess(c echo.Context) (user *api.GetDiscordUserInfoResponse, role enum.UserRole, err error) {
	user, role, err = h.getDiscordUserInfoAndRoleByCookie(c)
	if err != nil {
		err = apperrors.HandleInternalServerError(err)
	} else if role.IsGuest() {
		err = apperrors.HandleUnauthorized()
	} else if !role.IsAdmin() {
		err = apperrors.HandleAdminOnly()
	}

	return
}

func (h *Handler) enforceUserOrAboveAccess(c echo.Context) (user *api.GetDiscordUserInfoResponse, role enum.UserRole, err error) {
	user, role, err = h.getDiscordUserInfoAndRoleByCookie(c)
	if err != nil {
		err = apperrors.HandleInternalServerError(err)
	} else if role.IsGuest() {
		err = apperrors.HandleUnauthorized()
	} else if !role.IsUserOrAbove() {
		err = apperrors.HandleUserOrAbove()
	}

	return
}
