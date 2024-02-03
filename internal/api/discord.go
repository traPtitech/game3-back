package api

import (
	"encoding/json"
	"errors"
	"github.com/traPtitech/game3-back/internal/enum"
	"github.com/traPtitech/game3-back/internal/pkg/util"
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
	clientId, err := util.GetEnvOrErr("DISCORD_CLIENT_ID")
	if err != nil {
		return nil, err
	}
	clientSecret, err := util.GetEnvOrErr("DISCORD_CLIENT_SECRET")
	if err != nil {
		return nil, err
	}
	clientRedirectURI, err := util.GetEnvOrErr("DISCORD_CLIENT_REDIRECT_URI")
	if err != nil {
		return nil, err
	}

	formData := url.Values{}
	formData.Set("client_id", clientId)
	formData.Set("client_secret", clientSecret)
	formData.Set("grant_type", "authorization_code")
	formData.Set("code", params.Code)
	formData.Set("redirect_uri", clientRedirectURI)
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

func GetDiscordUserRole(discordUserID string) enum.UserRole {
	//TODO Discordのユーザーのロールを取得する BOTを使う必要がありそう、とりあえずハードコーディング
	adminUserIDs := []string{"310006192917315585", "222725046987128837", "707176617210019850", "1088448230662099024", "457179092127711232", "855464089496453171", "350623253141782528", "818846297535676456"}
	for _, id := range adminUserIDs {
		if id == discordUserID {
			return enum.Admin
		}
	}

	return enum.User
}

func GetDiscordOAuthRedirectURI() (string, error) {
	clientID, err := util.GetEnvOrErr("DISCORD_CLIENT_ID")
	if err != nil {
		return "", err
	}
	clientRedirectURI, err := util.GetEnvOrErr("DISCORD_CLIENT_REDIRECT_URI")
	if err != nil {
		return "", err
	}

	discordOAuthBaseURL := "https://discord.com/oauth2/authorize"
	u, err := url.Parse(discordOAuthBaseURL)
	if err != nil {
		return "", err
	}
	params := url.Values{}
	params.Add("client_id", clientID)
	params.Add("response_type", "code")
	params.Add("redirect_uri", clientRedirectURI)
	params.Add("scope", "identify")
	u.RawQuery = params.Encode()

	return u.String(), nil
}
