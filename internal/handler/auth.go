package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traPtitech/game3-back/internal/api/models"
	"github.com/traPtitech/game3-back/internal/domains"
	"github.com/traPtitech/game3-back/internal/repository"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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
        fetch('http://localhost:8080/auth/login', {
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

func New(repo *repository.Repository) *Handler {
	return &Handler{
		repo: repo,
	}
}

type TokenResponse struct {
	TokenType    *string `json:"token_type"`
	AccessToken  *string `json:"access_token"`
	ExpiresIn    *int    `json:"expires_in"`
	RefreshToken *string `json:"refresh_token"`
	Scope        *string `json:"scope"`
}

func (h *Handler) OauthCallback(c echo.Context, params models.OauthCallbackParams) error {
	cookie, err := c.Cookie("SessionToken")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "SessionToken is not found")
	}
	sessionToken, err := uuid.Parse(cookie.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "SessionToken is invalid")
	}

	formData := url.Values{}
	formData.Set("client_id", "1188893707215315045")
	formData.Set("client_secret", "HNmgqBqvYE2EowiFr88vSqq8gXAA5gWV")
	formData.Set("grant_type", "authorization_code")
	formData.Set("code", params.Code)
	formData.Set("redirect_uri", "http://localhost:8080/auth/callback")
	formData.Set("scope", "identify")

	req, err := http.NewRequest("POST", "https://discordapp.com/api/oauth2/token", strings.NewReader(formData.Encode()))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var tokenResponse TokenResponse
	if err = json.Unmarshal(body, &tokenResponse); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	CreateSessionParams := &domains.Session{
		ID:           &sessionToken,
		AccessToken:  tokenResponse.AccessToken,
		RefreshToken: tokenResponse.RefreshToken,
		ExpiresIn:    tokenResponse.ExpiresIn,
	}
	if err = h.repo.UpdateSession(CreateSessionParams); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	session, err := h.repo.GetSession(sessionToken.String())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(http.StatusSeeOther, *session.Redirect)
}

func (h *Handler) Login(c echo.Context) error {
	req := new(models.LoginJSONBody)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	if req.Redirect == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	sessionToken, err := uuid.NewRandom()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate session token")
	}

	CreateSessionParams := &domains.Session{
		ID:       &sessionToken,
		Redirect: req.Redirect,
	}
	if err = h.repo.CreateSession(CreateSessionParams); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	//Discord OAuth URL
	discordURL := "https://discord.com/api/oauth2/authorize?client_id=1188893707215315045&response_type=code&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fauth%2Fcallback&scope=identify"

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
