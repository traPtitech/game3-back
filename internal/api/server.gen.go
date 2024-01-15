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
	// ゲームを登録
	// (POST /games)
	PostGame(ctx echo.Context) error
	// ゲーム情報を取得
	// (GET /games/{gameId})
	GetGame(ctx echo.Context, gameId GameIdInPath) error
	// ゲーム情報を変更
	// (PATCH /games/{gameId})
	PatchGame(ctx echo.Context, gameId GameIdInPath) error
	// ゲームの画像を取得
	// (GET /games/{gameId}/image)
	GetGameImage(ctx echo.Context, gameId GameIdInPath) error
	// サーバーの生存確認
	// (GET /ping)
	PingServer(ctx echo.Context) error
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
	router.POST(baseURL+"/games", wrapper.PostGame)
	router.GET(baseURL+"/games/:gameId", wrapper.GetGame)
	router.PATCH(baseURL+"/games/:gameId", wrapper.PatchGame)
	router.GET(baseURL+"/games/:gameId/image", wrapper.GetGameImage)
	router.GET(baseURL+"/ping", wrapper.PingServer)
	router.GET(baseURL+"/test", wrapper.Test)
	router.GET(baseURL+"/users/me", wrapper.GetMe)
	router.GET(baseURL+"/users/me/games", wrapper.GetMeGames)
	router.GET(baseURL+"/users/:userId", wrapper.GetUser)
	router.GET(baseURL+"/users/:userId/games", wrapper.GetUserGames)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"arF3Ggdos1Oz/w8AAP//DPhr+Wo2AAA=",
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
