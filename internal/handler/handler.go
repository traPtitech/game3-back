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
func (h *Handler) handleFile(c echo.Context, formFileName string) (*types.File, error) {
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
