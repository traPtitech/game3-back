package apperrors

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleDbError(err error) error {
	if errors.As(err, &sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Not Found (DB error: %s)", err.Error()))
	}

	return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error (DB error: %s)", err.Error()))
}

func HandleBindError(err error) error {
	return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Bad Request (Bind error: %s)", err.Error()))
}

func HandleFileError(err error) error {
	return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error (Failed to get file: %s)", err.Error()))
}

// HandleAdminOnly Adminのみ許可
func HandleAdminOnly() error {
	return echo.NewHTTPError(http.StatusForbidden, "Forbidden: Admin only")
}

// HandleUserOrAbove UserとAdminを許可
func HandleUserOrAbove() error {
	return echo.NewHTTPError(http.StatusForbidden, "Unauthorized: please login (You are Guest)")
}

// HandleUnauthorized please login
func HandleUnauthorized() error {
	return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized: please login")
}

// HandleInternalServerError サーバーエラー
func HandleInternalServerError(err error) error {
	return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error (error: %s)", err.Error()))
}
