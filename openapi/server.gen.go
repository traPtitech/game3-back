// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package openapi

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
	. "github.com/traPtitech/game3-back/openapi/models"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// OAuth認証コールバック
	// (GET /auth/callback)
	OauthCallback(ctx echo.Context, params OauthCallbackParams) error
	// ログイン
	// (GET /auth/login)
	Login(ctx echo.Context, params LoginParams) error
	// ログアウト
	// (POST /auth/logout)
	Logout(ctx echo.Context) error
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
	// (GET /events/{eventSlug})
	GetEvent(ctx echo.Context, eventSlug EventSlugInPath) error
	// イベントの情報を変更
	// (PATCH /events/{eventSlug})
	PatchEvent(ctx echo.Context, eventSlug EventSlugInPath) error
	// イベントとイベントに登録されているゲームの情報をCSV形式で取得
	// (GET /events/{eventSlug}/csv)
	GetEventCsv(ctx echo.Context, eventSlug EventSlugInPath) error
	// イベントに登録されているゲームのリストを取得
	// (GET /events/{eventSlug}/games)
	GetEventGames(ctx echo.Context, eventSlug EventSlugInPath) error
	// イベントの画像を取得
	// (GET /events/{eventSlug}/image)
	GetEventImage(ctx echo.Context, eventSlug EventSlugInPath) error
	// イベントに登録されているタームのリストを取得
	// (GET /events/{eventSlug}/terms)
	GetEventTerms(ctx echo.Context, eventSlug EventSlugInPath) error
	// ゲームのリストを取得 GET /games?termId=X&eventSlug=X&userId=X&include=unpublished
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

	// Parameter object where we will unmarshal all parameters from the context
	var params LoginParams
	// ------------- Required query parameter "redirect" -------------

	err = runtime.BindQueryParameter("form", true, true, "redirect", ctx.QueryParams(), &params.Redirect)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter redirect: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Login(ctx, params)
	return err
}

// Logout converts echo context to params.
func (w *ServerInterfaceWrapper) Logout(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Logout(ctx)
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
	// ------------- Path parameter "eventSlug" -------------
	var eventSlug EventSlugInPath

	err = runtime.BindStyledParameterWithLocation("simple", false, "eventSlug", runtime.ParamLocationPath, ctx.Param("eventSlug"), &eventSlug)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eventSlug: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEvent(ctx, eventSlug)
	return err
}

// PatchEvent converts echo context to params.
func (w *ServerInterfaceWrapper) PatchEvent(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "eventSlug" -------------
	var eventSlug EventSlugInPath

	err = runtime.BindStyledParameterWithLocation("simple", false, "eventSlug", runtime.ParamLocationPath, ctx.Param("eventSlug"), &eventSlug)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eventSlug: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchEvent(ctx, eventSlug)
	return err
}

// GetEventCsv converts echo context to params.
func (w *ServerInterfaceWrapper) GetEventCsv(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "eventSlug" -------------
	var eventSlug EventSlugInPath

	err = runtime.BindStyledParameterWithLocation("simple", false, "eventSlug", runtime.ParamLocationPath, ctx.Param("eventSlug"), &eventSlug)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eventSlug: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEventCsv(ctx, eventSlug)
	return err
}

// GetEventGames converts echo context to params.
func (w *ServerInterfaceWrapper) GetEventGames(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "eventSlug" -------------
	var eventSlug EventSlugInPath

	err = runtime.BindStyledParameterWithLocation("simple", false, "eventSlug", runtime.ParamLocationPath, ctx.Param("eventSlug"), &eventSlug)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eventSlug: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEventGames(ctx, eventSlug)
	return err
}

// GetEventImage converts echo context to params.
func (w *ServerInterfaceWrapper) GetEventImage(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "eventSlug" -------------
	var eventSlug EventSlugInPath

	err = runtime.BindStyledParameterWithLocation("simple", false, "eventSlug", runtime.ParamLocationPath, ctx.Param("eventSlug"), &eventSlug)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eventSlug: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEventImage(ctx, eventSlug)
	return err
}

// GetEventTerms converts echo context to params.
func (w *ServerInterfaceWrapper) GetEventTerms(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "eventSlug" -------------
	var eventSlug EventSlugInPath

	err = runtime.BindStyledParameterWithLocation("simple", false, "eventSlug", runtime.ParamLocationPath, ctx.Param("eventSlug"), &eventSlug)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eventSlug: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEventTerms(ctx, eventSlug)
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

	// ------------- Optional query parameter "eventSlug" -------------

	err = runtime.BindQueryParameter("form", true, false, "eventSlug", ctx.QueryParams(), &params.EventSlug)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eventSlug: %s", err))
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
	router.GET(baseURL+"/auth/login", wrapper.Login)
	router.POST(baseURL+"/auth/logout", wrapper.Logout)
	router.GET(baseURL+"/events", wrapper.GetEvents)
	router.POST(baseURL+"/events", wrapper.PostEvent)
	router.GET(baseURL+"/events/now", wrapper.GetCurrentEvent)
	router.GET(baseURL+"/events/:eventSlug", wrapper.GetEvent)
	router.PATCH(baseURL+"/events/:eventSlug", wrapper.PatchEvent)
	router.GET(baseURL+"/events/:eventSlug/csv", wrapper.GetEventCsv)
	router.GET(baseURL+"/events/:eventSlug/games", wrapper.GetEventGames)
	router.GET(baseURL+"/events/:eventSlug/image", wrapper.GetEventImage)
	router.GET(baseURL+"/events/:eventSlug/terms", wrapper.GetEventTerms)
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
