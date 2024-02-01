package apperrors

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleDbError(err error) error {
	if !errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Not Found (DB error: %s)", err.Error()))
	}

	return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error (DB error: %s)"+err.Error()))
}

func HandleBindError(err error) error {
	return echo.NewHTTPError(http.StatusBadRequest, "Bad Request (Bind error: %s)", err.Error())
}

func HandleFileError(err error) error {
	return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error (Failed to get file: %s)", err.Error())
}
