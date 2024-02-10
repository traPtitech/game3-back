package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traPtitech/game3-back/internal/pkg/apperrors"
	"github.com/traPtitech/game3-back/openapi/models"
	"net/http"
)

func (h *Handler) GetTerms(c echo.Context) error {
	terms, err := h.repo.GetTerms()
	if err != nil {
		return apperrors.HandleDbError(err)
	}

	return c.JSON(http.StatusOK, terms)
}

func (h *Handler) PostTerm(c echo.Context) error {
	if _, _, err := h.enforceAdminAccess(c); err != nil {
		return err
	}

	req := &models.PostTermRequest{}
	if err := c.Bind(req); err != nil {
		return apperrors.HandleBindError(err)
	}

	newTermID := uuid.New()
	if err := h.repo.PostTerm(newTermID, req); err != nil {
		return apperrors.HandleDbError(err)
	}

	term, err := h.repo.GetTerm(newTermID)
	if err != nil {
		return apperrors.HandleDbError(err)
	}

	return c.JSON(http.StatusCreated, term)
}

func (h *Handler) GetTerm(c echo.Context, termID models.TermIdInPath) error {
	term, err := h.repo.GetTerm(termID)
	if err != nil {
		return apperrors.HandleDbError(err)
	}

	return c.JSON(http.StatusOK, term)
}

func (h *Handler) PatchTerm(c echo.Context, termID models.TermIdInPath) error {
	if _, _, err := h.enforceAdminAccess(c); err != nil {
		return err
	}

	req := &models.PatchTermRequest{}
	if err := c.Bind(req); err != nil {
		return apperrors.HandleBindError(err)
	}

	if err := h.repo.PatchTerm(termID, req); err != nil {
		return apperrors.HandleDbError(err)
	}

	term, err := h.repo.GetTerm(termID)
	if err != nil {
		return apperrors.HandleDbError(err)
	}

	return c.JSON(http.StatusOK, term)

}

func (h *Handler) GetTermGames(c echo.Context, termID models.TermIdInPath) error {
	games, err := h.repo.GetGames(models.GetGamesParams{TermId: &termID})
	if err != nil {
		return apperrors.HandleDbError(err)
	}

	return c.JSON(http.StatusOK, games)
}
