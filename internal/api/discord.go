package api

import (
	"encoding/json"
	"errors"
	"github.com/traPtitech/game3-back/openapi/models"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
)

type DiscordUserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type TokenResponse struct {
	TokenType    *string `json:"token_type"`
	AccessToken  *string `json:"access_token"`
	ExpiresIn    *int    `json:"expires_in"`
	RefreshToken *string `json:"refresh_token"`
	Scope        *string `json:"scope"`
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

func GetDiscordUserToken(params models.OauthCallbackParams) (*TokenResponse, error) {
	formData := url.Values{}
	formData.Set("client_id", "1188893707215315045")
	formData.Set("client_secret", "HNmgqBqvYE2EowiFr88vSqq8gXAA5gWV")
	formData.Set("grant_type", "authorization_code")
	formData.Set("code", params.Code)
	formData.Set("redirect_uri", "http://localhost:8080/api/auth/callback")
	formData.Set("scope", "identify")

	req, err := http.NewRequest("POST", "https://discordapp.com/api/oauth2/token", strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Discord OAuth failed: status: " + resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var tokenResponse TokenResponse
	if err = json.Unmarshal(body, &tokenResponse); err != nil {
		return nil, err
	}

	return &tokenResponse, nil
}
