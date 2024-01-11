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
