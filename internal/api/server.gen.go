// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	. "github.com/traPtitech/game3-back/internal/api/models"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// OAuth認証コールバック
	// (GET /auth/callback)
	OauthCallback(ctx echo.Context, params OauthCallbackParams) error
	// ログイン
	// (POST /auth/login)
	Login(ctx echo.Context) error
	// ログアウト
	// (POST /auth/logout)
	Logout(ctx echo.Context) error
	// コンタクト用メールを送信
	// (POST /contacts)
	PostContacts(ctx echo.Context) error
	// イベントのリストを取得
	// (GET /events)
	GetEvents(ctx echo.Context) error
	// イベントを登録
	// (POST /events)
	PostEvent(ctx echo.Context) error
	// 開催中のイベントを取得
	// (GET /events/now)
	GetCurrentEvent(ctx echo.Context) error
	// イベントの情報を取得
	// (GET /events/{eventId})
	GetEvent(ctx echo.Context, eventId EventIdInPath) error
	// イベントの情報を変更
	// (PATCH /events/{eventId})
	PatchEvent(ctx echo.Context, eventId EventIdInPath) error
	// イベントとイベントに登録されているゲームの情報をCSV形式で取得
	// (GET /events/{eventId}/csv)
	GetEventCsv(ctx echo.Context, eventId EventIdInPath) error
	// イベントに登録されているゲームのリストを取得
	// (GET /events/{eventId}/games)
	GetEventGames(ctx echo.Context, eventId EventIdInPath) error
	// イベントの画像を取得
	// (GET /events/{eventId}/image)
	GetEventImage(ctx echo.Context, eventId EventIdInPath) error
	// イベントに登録されているタームのリストを取得
	// (GET /events/{eventId}/terms)
	GetEventTerms(ctx echo.Context, eventId EventIdInPath) error
	// ゲームのリストを取得 GET /games?termId=X&eventId=X&userId=X&include=unpublished
	// (GET /games)
	GetGames(ctx echo.Context, params GetGamesParams) error
	// ゲームを登録
	// (POST /games)
	PostGame(ctx echo.Context) error
	// ゲーム情報を取得
	// (GET /games/{gameId})
	GetGame(ctx echo.Context, gameId GameIdInPath) error
	// ゲーム情報を変更
	// (PATCH /games/{gameId})
	PatchGame(ctx echo.Context, gameId GameIdInPath) error
	// ゲームのアイコン画像を取得
	// (GET /games/{gameId}/icon)
	GetGameIcon(ctx echo.Context, gameId GameIdInPath) error
	// ゲームの画像を取得
	// (GET /games/{gameId}/image)
	GetGameImage(ctx echo.Context, gameId GameIdInPath) error
	// サーバーの生存確認
	// (GET /ping)
	PingServer(ctx echo.Context) error
	// イベントに登録されているタームのリストを取得
	// (GET /terms)
	GetTerms(ctx echo.Context) error
	// タームを登録
	// (POST /terms)
	PostTerm(ctx echo.Context) error
	// ターム情報を取得
	// (GET /terms/{termId})
	GetTerm(ctx echo.Context, termId TermIdInPath) error
	// ターム情報を変更
	// (PATCH /terms/{termId})
	PatchTerm(ctx echo.Context, termId TermIdInPath) error
	// タームに登録されているゲームのリストを取得
	// (GET /terms/{termId}/games)
	GetTermGames(ctx echo.Context, termId TermIdInPath) error
	// テスト用
	// (GET /test)
	Test(ctx echo.Context) error
	// 自分のユーザー情報を取得
	// (GET /users/me)
	GetMe(ctx echo.Context) error
	// 自分が登録したゲームのリストを取得
	// (GET /users/me/games)
	GetMeGames(ctx echo.Context) error
	// ユーザー情報を取得
	// (GET /users/{userId})
	GetUser(ctx echo.Context, userId UserIdInPath) error
	// ユーザーが登録したゲームのリストを取得
	// (GET /users/{userId}/games)
	GetUserGames(ctx echo.Context, userId UserIdInPath) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// OauthCallback converts echo context to params.
func (w *ServerInterfaceWrapper) OauthCallback(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params OauthCallbackParams
	// ------------- Required query parameter "code" -------------

	err = runtime.BindQueryParameter("form", true, true, "code", ctx.QueryParams(), &params.Code)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter code: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.OauthCallback(ctx, params)
	return err
}

// Login converts echo context to params.
func (w *ServerInterfaceWrapper) Login(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Login(ctx)
	return err
}

// Logout converts echo context to params.
func (w *ServerInterfaceWrapper) Logout(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Logout(ctx)
	return err
}

// PostContacts converts echo context to params.
func (w *ServerInterfaceWrapper) PostContacts(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostContacts(ctx)
	return err
}

// GetEvents converts echo context to params.
func (w *ServerInterfaceWrapper) GetEvents(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEvents(ctx)
	return err
}

// PostEvent converts echo context to params.
func (w *ServerInterfaceWrapper) PostEvent(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostEvent(ctx)
	return err
}

// GetCurrentEvent converts echo context to params.
func (w *ServerInterfaceWrapper) GetCurrentEvent(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetCurrentEvent(ctx)
	return err
}

// GetEvent converts echo context to params.
func (w *ServerInterfaceWrapper) GetEvent(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "eventId" -------------
	var eventId EventIdInPath

	err = runtime.BindStyledParameterWithLocation("simple", false, "eventId", runtime.ParamLocationPath, ctx.Param("eventId"), &eventId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eventId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEvent(ctx, eventId)
	return err
}

// PatchEvent converts echo context to params.
func (w *ServerInterfaceWrapper) PatchEvent(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "eventId" -------------
	var eventId EventIdInPath

	err = runtime.BindStyledParameterWithLocation("simple", false, "eventId", runtime.ParamLocationPath, ctx.Param("eventId"), &eventId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eventId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchEvent(ctx, eventId)
	return err
}

// GetEventCsv converts echo context to params.
func (w *ServerInterfaceWrapper) GetEventCsv(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "eventId" -------------
	var eventId EventIdInPath

	err = runtime.BindStyledParameterWithLocation("simple", false, "eventId", runtime.ParamLocationPath, ctx.Param("eventId"), &eventId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eventId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEventCsv(ctx, eventId)
	return err
}

// GetEventGames converts echo context to params.
func (w *ServerInterfaceWrapper) GetEventGames(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "eventId" -------------
	var eventId EventIdInPath

	err = runtime.BindStyledParameterWithLocation("simple", false, "eventId", runtime.ParamLocationPath, ctx.Param("eventId"), &eventId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eventId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEventGames(ctx, eventId)
	return err
}

// GetEventImage converts echo context to params.
func (w *ServerInterfaceWrapper) GetEventImage(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "eventId" -------------
	var eventId EventIdInPath

	err = runtime.BindStyledParameterWithLocation("simple", false, "eventId", runtime.ParamLocationPath, ctx.Param("eventId"), &eventId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eventId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEventImage(ctx, eventId)
	return err
}

// GetEventTerms converts echo context to params.
func (w *ServerInterfaceWrapper) GetEventTerms(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "eventId" -------------
	var eventId EventIdInPath

	err = runtime.BindStyledParameterWithLocation("simple", false, "eventId", runtime.ParamLocationPath, ctx.Param("eventId"), &eventId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eventId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEventTerms(ctx, eventId)
	return err
}

// GetGames converts echo context to params.
func (w *ServerInterfaceWrapper) GetGames(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetGamesParams
	// ------------- Optional query parameter "termId" -------------

	err = runtime.BindQueryParameter("form", true, false, "termId", ctx.QueryParams(), &params.TermId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter termId: %s", err))
	}

	// ------------- Optional query parameter "eventId" -------------

	err = runtime.BindQueryParameter("form", true, false, "eventId", ctx.QueryParams(), &params.EventId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eventId: %s", err))
	}

	// ------------- Optional query parameter "userId" -------------

	err = runtime.BindQueryParameter("form", true, false, "userId", ctx.QueryParams(), &params.UserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	// ------------- Optional query parameter "include" -------------

	err = runtime.BindQueryParameter("form", true, false, "include", ctx.QueryParams(), &params.Include)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter include: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetGames(ctx, params)
	return err
}

// PostGame converts echo context to params.
func (w *ServerInterfaceWrapper) PostGame(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostGame(ctx)
	return err
}

// GetGame converts echo context to params.
func (w *ServerInterfaceWrapper) GetGame(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "gameId" -------------
	var gameId GameIdInPath

	err = runtime.BindStyledParameterWithLocation("simple", false, "gameId", runtime.ParamLocationPath, ctx.Param("gameId"), &gameId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter gameId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetGame(ctx, gameId)
	return err
}

// PatchGame converts echo context to params.
func (w *ServerInterfaceWrapper) PatchGame(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "gameId" -------------
	var gameId GameIdInPath

	err = runtime.BindStyledParameterWithLocation("simple", false, "gameId", runtime.ParamLocationPath, ctx.Param("gameId"), &gameId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter gameId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchGame(ctx, gameId)
	return err
}

// GetGameIcon converts echo context to params.
func (w *ServerInterfaceWrapper) GetGameIcon(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "gameId" -------------
	var gameId GameIdInPath

	err = runtime.BindStyledParameterWithLocation("simple", false, "gameId", runtime.ParamLocationPath, ctx.Param("gameId"), &gameId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter gameId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetGameIcon(ctx, gameId)
	return err
}

// GetGameImage converts echo context to params.
func (w *ServerInterfaceWrapper) GetGameImage(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "gameId" -------------
	var gameId GameIdInPath

	err = runtime.BindStyledParameterWithLocation("simple", false, "gameId", runtime.ParamLocationPath, ctx.Param("gameId"), &gameId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter gameId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetGameImage(ctx, gameId)
	return err
}

// PingServer converts echo context to params.
func (w *ServerInterfaceWrapper) PingServer(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PingServer(ctx)
	return err
}

// GetTerms converts echo context to params.
func (w *ServerInterfaceWrapper) GetTerms(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTerms(ctx)
	return err
}

// PostTerm converts echo context to params.
func (w *ServerInterfaceWrapper) PostTerm(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostTerm(ctx)
	return err
}

// GetTerm converts echo context to params.
func (w *ServerInterfaceWrapper) GetTerm(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "termId" -------------
	var termId TermIdInPath

	err = runtime.BindStyledParameterWithLocation("simple", false, "termId", runtime.ParamLocationPath, ctx.Param("termId"), &termId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter termId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTerm(ctx, termId)
	return err
}

// PatchTerm converts echo context to params.
func (w *ServerInterfaceWrapper) PatchTerm(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "termId" -------------
	var termId TermIdInPath

	err = runtime.BindStyledParameterWithLocation("simple", false, "termId", runtime.ParamLocationPath, ctx.Param("termId"), &termId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter termId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchTerm(ctx, termId)
	return err
}

// GetTermGames converts echo context to params.
func (w *ServerInterfaceWrapper) GetTermGames(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "termId" -------------
	var termId TermIdInPath

	err = runtime.BindStyledParameterWithLocation("simple", false, "termId", runtime.ParamLocationPath, ctx.Param("termId"), &termId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter termId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTermGames(ctx, termId)
	return err
}

// Test converts echo context to params.
func (w *ServerInterfaceWrapper) Test(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Test(ctx)
	return err
}

// GetMe converts echo context to params.
func (w *ServerInterfaceWrapper) GetMe(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetMe(ctx)
	return err
}

// GetMeGames converts echo context to params.
func (w *ServerInterfaceWrapper) GetMeGames(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetMeGames(ctx)
	return err
}

// GetUser converts echo context to params.
func (w *ServerInterfaceWrapper) GetUser(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId UserIdInPath

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUser(ctx, userId)
	return err
}

// GetUserGames converts echo context to params.
func (w *ServerInterfaceWrapper) GetUserGames(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId UserIdInPath

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUserGames(ctx, userId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/auth/callback", wrapper.OauthCallback)
	router.POST(baseURL+"/auth/login", wrapper.Login)
	router.POST(baseURL+"/auth/logout", wrapper.Logout)
	router.POST(baseURL+"/contacts", wrapper.PostContacts)
	router.GET(baseURL+"/events", wrapper.GetEvents)
	router.POST(baseURL+"/events", wrapper.PostEvent)
	router.GET(baseURL+"/events/now", wrapper.GetCurrentEvent)
	router.GET(baseURL+"/events/:eventId", wrapper.GetEvent)
	router.PATCH(baseURL+"/events/:eventId", wrapper.PatchEvent)
	router.GET(baseURL+"/events/:eventId/csv", wrapper.GetEventCsv)
	router.GET(baseURL+"/events/:eventId/games", wrapper.GetEventGames)
	router.GET(baseURL+"/events/:eventId/image", wrapper.GetEventImage)
	router.GET(baseURL+"/events/:eventId/terms", wrapper.GetEventTerms)
	router.GET(baseURL+"/games", wrapper.GetGames)
	router.POST(baseURL+"/games", wrapper.PostGame)
	router.GET(baseURL+"/games/:gameId", wrapper.GetGame)
	router.PATCH(baseURL+"/games/:gameId", wrapper.PatchGame)
	router.GET(baseURL+"/games/:gameId/icon", wrapper.GetGameIcon)
	router.GET(baseURL+"/games/:gameId/image", wrapper.GetGameImage)
	router.GET(baseURL+"/ping", wrapper.PingServer)
	router.GET(baseURL+"/terms", wrapper.GetTerms)
	router.POST(baseURL+"/terms", wrapper.PostTerm)
	router.GET(baseURL+"/terms/:termId", wrapper.GetTerm)
	router.PATCH(baseURL+"/terms/:termId", wrapper.PatchTerm)
	router.GET(baseURL+"/terms/:termId/games", wrapper.GetTermGames)
	router.GET(baseURL+"/test", wrapper.Test)
	router.GET(baseURL+"/users/me", wrapper.GetMe)
	router.GET(baseURL+"/users/me/games", wrapper.GetMeGames)
	router.GET(baseURL+"/users/:userId", wrapper.GetUser)
	router.GET(baseURL+"/users/:userId/games", wrapper.GetUserGames)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
