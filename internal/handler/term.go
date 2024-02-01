package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traPtitech/game3-back/openapi/models"
	"net/http"
)

func (h *Handler) GetTerms(c echo.Context) error {
	terms, err := h.repo.GetTerms()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, terms)
}

func (h *Handler) PostTerm(c echo.Context) error {
	req := &models.PostTermRequest{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	newTermID := uuid.New()
	if err := h.repo.PostTerm(newTermID, req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	term, err := h.repo.GetTerm(newTermID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, term)
}

func (h *Handler) GetTerm(c echo.Context, termID models.TermIdInPath) error {
	term, err := h.repo.GetTerm(termID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, term)
}

func (h *Handler) PatchTerm(c echo.Context, termID models.TermIdInPath) error {
	req := &models.PatchTermRequest{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.repo.PatchTerm(termID, req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	term, err := h.repo.GetTerm(termID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, term)

}

func (h *Handler) GetTermGames(c echo.Context, termID models.TermIdInPath) error {
	games, err := h.repo.GetGames(models.GetGamesParams{TermId: &termID})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, games)
}
