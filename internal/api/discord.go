package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DiscordUserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

func GetDiscordUserInfo(accessToken *string) (*DiscordUserResponse, error) {
	req, err := http.NewRequest("GET", "https://discordapp.com/api/users/@me", nil)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	req.Header.Set("Authorization", "Bearer "+*accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to get Discord user: status "+resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var discordUser DiscordUserResponse
	if err = json.Unmarshal(body, &discordUser); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return &discordUser, nil
}
