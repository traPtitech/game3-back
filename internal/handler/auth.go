package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traPtitech/game3-back/internal/api"
	"github.com/traPtitech/game3-back/internal/apperrors"
	"github.com/traPtitech/game3-back/internal/domain"
	"github.com/traPtitech/game3-back/internal/pkg/util"
	"github.com/traPtitech/game3-back/openapi/models"
	"net/http"
	"time"
)

func (h *Handler) Test(c echo.Context) error {
	//// すべてのクッキーを取得
	//cookies := c.Cookies()
	//
	//// クッキーの一覧を作成
	//var cookieList string
	//for _, cookie := range cookies {
	//	cookieList += "Name: " + cookie.Name + ", Value: " + cookie.Value + "\n"
	//}
	//
	//// クッキーの一覧をレスポンスとして返す
	//return c.String(http.StatusOK, cookieList)
	//discordURL := "https://discord.com/api/oauth2/authorize?client_id=1188893707215315045&response_type=code&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fauth%2Fcallback&scope=identify"
	//return c.Redirect(http.StatusFound, discordURL)
	html := `
<html>
<head>
    <title>リダイレクトテスト</title>
    <script>
      function sendGetRequest() {
        window.location.href = 'http://localhost:8080/api/auth/login?redirect=http://localhost:8080/api/ping';
      }
    </script>
</head>
<body>
<button onClick='sendGetRequest()'>Submit</button>
</body>
</html>

`

	return c.HTML(http.StatusOK, html)
}

func (h *Handler) OauthCallback(c echo.Context, params models.OauthCallbackParams) error {
	sessionID, err := getSessionIDByCookie(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	tokenResponse, err := api.GetDiscordUserToken(params)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	userServers, err := api.GetDiscordUserServers(tokenResponse.AccessToken)
	if err != nil {
		return err
	}

	isInGame3Server, err := checkUserIsInGame3Server(userServers)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if isInGame3Server == false {
		err = addUserToGame3Guild(tokenResponse.AccessToken)
		if err != nil {
			return err
		}
	}

	CreateSessionParams := &domain.Session{
		ID:           sessionID,
		AccessToken:  tokenResponse.AccessToken,
		RefreshToken: tokenResponse.RefreshToken,
		ExpiresIn:    tokenResponse.ExpiresIn,
	}

	if err = h.repo.UpdateSession(CreateSessionParams); err != nil {
		return apperrors.HandleDbError(err)
	}
	session, err := h.repo.GetSession(sessionID)
	if err != nil {
		return apperrors.HandleDbError(err)
	}

	return c.Redirect(http.StatusSeeOther, *session.Redirect)
}

func checkUserIsInGame3Server(userServers []api.GetDiscordUserGuildsResponse) (bool, error) {
	game3ServerID, err := util.GetEnvOrErr("DISCORD_SERVER_ID")
	if err != nil {
		return false, err
	}
	for _, server := range userServers {
		if server.ID == game3ServerID {
			return true, nil
		}
	}

	return false, nil
}

func addUserToGame3Guild(accessToken *string) error {
	game3ServerID, err := util.GetEnvOrErr("DISCORD_SERVER_ID")
	if err != nil {
		return err
	}
	user, err := api.GetDiscordUserInfo(accessToken)
	if err != nil {
		return err
	}
	err = api.AddUserToGuild(accessToken, game3ServerID, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) Login(c echo.Context, params models.LoginParams) error {
	sessionToken := uuid.New()
	CreateSessionParams := &domain.Session{
		ID:       &sessionToken,
		Redirect: &params.Redirect,
	}
	if err := h.repo.CreateSession(CreateSessionParams); err != nil {
		return apperrors.HandleDbError(err)
	}

	discordURL, err := api.GetDiscordOAuthRedirectURI()
	if err != nil {
		return apperrors.HandleInternalServerError(err)
	}

	c.SetCookie(&http.Cookie{
		Name:     "SessionToken",
		Value:    sessionToken.String(),
		Path:     "/",
		HttpOnly: true, // JavaScriptからのアクセスを防ぐ
		//Secure:   true, // HTTPSを通じてのみCookieを送信
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(24 * time.Hour),
		//Domain:   "localhost",
	})

	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	return c.Redirect(http.StatusSeeOther, discordURL)
}

func (h *Handler) Logout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:     "SessionToken",
		Value:    "",
		Path:     "/",
		HttpOnly: true, // JavaScriptからのアクセスを防ぐ
		//Secure:   true, // HTTPSを通じてのみCookieを送信
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(-1 * time.Hour),
		//Domain:   "localhost",
	})

	return c.Redirect(http.StatusSeeOther, "/")
}
