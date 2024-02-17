package handler

import (
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime/types"
	"github.com/traPtitech/game3-back/internal/api"
	"github.com/traPtitech/game3-back/internal/pkg/apperrors"
	"github.com/traPtitech/game3-back/internal/pkg/enum"
	"github.com/traPtitech/game3-back/internal/pkg/util"
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
func _(c echo.Context, err error) error {
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

// handleImageFileAndConvertImageToPNGAndResizeImage reads the image file from the form and converts it to PNG or GIF format.
func handleImageFileAndConvertImageToPNGAndResizeImage(c echo.Context, formFileName string, maxWidth, maxHeight int) (*types.File, error) {
	fileHeader, err := c.FormFile(formFileName)
	if err != nil {
		if errors.Is(err, http.ErrMissingFile) {
			return nil, nil
		}

		return nil, err
	}

	fileSrc, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer fileSrc.Close()

	srcImg, format, err := util.DecodeImage(fileSrc)
	if err != nil {
		return nil, err
	}

	var resizedImageData []byte
	if format == "gif" {
		gifData, err := util.DecodeGifImage(fileSrc)
		if err != nil {
			return nil, err
		}

		// 画像をリサイズ
		resizedGifData, err := util.ResizeGifImageMaintainingAspectRatio(c, gifData, maxWidth, maxHeight)
		if err != nil {
			return nil, err
		}

		resizedImageData, err = util.EncodeGifImage(resizedGifData)
		if err != nil {
			return nil, err
		}
	} else {
		// 画像をPNG形式に変換
		pngData, err := util.EncodeImageToPNG(srcImg)
		if err != nil {
			return nil, err
		}

		// 画像をリサイズ
		resizedImageData, err = util.ResizePngImageMaintainingAspectRatio(pngData, maxWidth, maxHeight)
		if err != nil {
			return nil, err
		}
	}

	file := types.File{}
	file.InitFromBytes(resizedImageData, fileHeader.Filename)

	return &file, nil
}

// getSessionIdByCookie gets the session ID from the SessionToken cookie.
func getSessionIDByCookie(c echo.Context) (*uuid.UUID, error) {
	sessionByCookie, err := c.Cookie("SessionToken")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
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

func validateCacheAndUpdateHeader(c echo.Context, lastModifiedTime string) error {
	ifModifiedSince := c.Request().Header.Get("If-Modified-Since")
	if ifModifiedSince == lastModifiedTime {
		return c.NoContent(http.StatusNotModified)
	}
	c.Response().Header().Set("Last-Modified", lastModifiedTime)

	return nil
}
