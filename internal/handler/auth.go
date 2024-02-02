package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traPtitech/game3-back/internal/api"
	"github.com/traPtitech/game3-back/internal/apperrors"
	"github.com/traPtitech/game3-back/internal/domain"
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
      function sendPostRequest() {
        fetch('http://localhost:8080/api/auth/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({"redirect": "dekita"}),
          redirect: 'follow'
        })
          .then(response => {
            if (response.ok) {
              if (response.redirected) {
                window.location.href = response.url;
              } else {
                return response.json();
              }
            } else {
              throw new Error('Network response was not ok');
            }
          })
          .then(data => {
            if (data) {
              console.log(data);
            }
          })
          .catch(error => {
            console.error('Fetch error:', error);
          });
      }
    </script>
</head>
<body>
<button onClick='sendPostRequest()'>Submit</button>
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

func (h *Handler) Login(c echo.Context) error {
	req := new(models.LoginJSONBody)
	if err := c.Bind(req); err != nil {
		return apperrors.HandleBindError(err)
	}

	if req.Redirect == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body: redirect is required")
	}

	sessionToken := uuid.New()
	CreateSessionParams := &domain.Session{
		ID:       &sessionToken,
		Redirect: req.Redirect,
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

	return c.Redirect(http.StatusSeeOther, discordURL)
}

func (h *Handler) Logout(c echo.Context) error {
	panic("implement me")
}
