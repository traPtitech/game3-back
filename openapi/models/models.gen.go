// Package models provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package models

import (
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Defines values for UserRole.
const (
	UserRoleAdmin UserRole = "admin"
	UserRoleGuest UserRole = "guest"
	UserRoleUser  UserRole = "user"
)

// Event defines model for Event.
type Event struct {
	// GameSubmissionPeriodEnd ゲーム展示の募集終了期間
	GameSubmissionPeriodEnd time.Time `db:"game_submission_period_end" json:"gameSubmissionPeriodEnd"`

	// GameSubmissionPeriodStart ゲーム展示の募集開始期間
	GameSubmissionPeriodStart time.Time `db:"game_submission_period_start" json:"gameSubmissionPeriodStart"`

	// Slug イベントslug (イベントのID的な立ち位置)
	Slug string `db:"slug" json:"slug"`

	// Title イベントのタイトル (例: 第18回)
	Title string `db:"title" json:"title"`
}

// Game defines model for Game.
type Game struct {
	// CreatorName ゲーム作成者
	CreatorName string `db:"creator_name" json:"creatorName"`

	// CreatorPageUrl ゲーム作成者のページのURL
	CreatorPageUrl *string `db:"creator_page_url" json:"creatorPageUrl,omitempty"`

	// Description ゲームの説明
	Description string `db:"description" json:"description"`

	// DiscordUserId DiscordのユーザーID
	DiscordUserId string `db:"discord_user_id" json:"discordUserId"`

	// GamePageUrl ゲームページのURL
	GamePageUrl *string `db:"game_page_url" json:"gamePageUrl,omitempty"`

	// Id ゲームID
	Id openapi_types.UUID `db:"id" json:"id"`

	// Place 展示場所
	Place *string `db:"place" json:"place,omitempty"`

	// TermId タームID
	TermId openapi_types.UUID `db:"term_id" json:"termId"`

	// Title 展示するゲームタイトル
	Title string `db:"title" json:"title"`
}

// GetEventCSVResponse defines model for GetEventCSVResponse.
type GetEventCSVResponse = string

// ImageResponse defines model for ImageResponse.
type ImageResponse = openapi_types.File

// PatchEventRequest defines model for PatchEventRequest.
type PatchEventRequest struct {
	// GameSubmissionPeriodEnd ゲーム展示の募集期間
	GameSubmissionPeriodEnd *time.Time `form:"gameSubmissionPeriodEnd" json:"gameSubmissionPeriodEnd,omitempty"`

	// GameSubmissionPeriodStart ゲーム展示の募集期間
	GameSubmissionPeriodStart *time.Time `form:"gameSubmissionPeriodStart" json:"gameSubmissionPeriodStart,omitempty"`

	// Image パンフレット用画像
	Image *openapi_types.File `form:"image" json:"image,omitempty"`

	// Slug slug (URL内で使用, unique) 18th等
	Slug *string `form:"slug" json:"slug,omitempty"`

	// Title 第18回
	Title *string `form:"title" json:"title,omitempty"`
}

// PatchGameRequest defines model for PatchGameRequest.
type PatchGameRequest struct {
	// CreatorName ゲーム作成者
	CreatorName *string `form:"creatorName" json:"creatorName,omitempty"`

	// CreatorPageUrl ゲーム作成者のページのURL
	CreatorPageUrl *string `form:"creatorPageUrl" json:"creatorPageUrl,omitempty"`

	// Description ゲームの説明
	Description *string `form:"description" json:"description,omitempty"`

	// DiscordUserId discordのユーザーID
	DiscordUserId *string `form:"discordUserId" json:"discordUserId,omitempty"`

	// GamePageUrl ゲームページのURL
	GamePageUrl *string `form:"gamePageUrl" json:"gamePageUrl,omitempty"`

	// Icon ゲームのアイコン画像
	Icon *openapi_types.File `form:"icon" json:"icon,omitempty"`

	// Image ゲームの画像
	Image *openapi_types.File `form:"image" json:"image,omitempty"`

	// Place 展示場所
	Place *string `form:"place" json:"place,omitempty"`

	// TermId タームID
	TermId *openapi_types.UUID `form:"termId" json:"termId,omitempty"`

	// Title 展示するゲームタイトル
	Title *string `form:"title" json:"title,omitempty"`
}

// PatchTermRequest defines model for PatchTermRequest.
type PatchTermRequest struct {
	// EndAt タームが終わる時間
	EndAt *time.Time `json:"endAt,omitempty"`

	// EventSlug イベントのslug
	EventSlug *string `json:"eventSlug,omitempty"`

	// IsDefault ゲーム登録時に割り当てられるTermならばTrue
	IsDefault *bool `json:"isDefault,omitempty"`

	// StartAt タームが始まる時間
	StartAt *time.Time `json:"startAt,omitempty"`
}

// PostEventRequest defines model for PostEventRequest.
type PostEventRequest struct {
	// GameSubmissionPeriodEnd ゲーム展示の募集期間
	GameSubmissionPeriodEnd time.Time `form:"gameSubmissionPeriodEnd" json:"gameSubmissionPeriodEnd"`

	// GameSubmissionPeriodStart ゲーム展示の募集期間
	GameSubmissionPeriodStart time.Time `form:"gameSubmissionPeriodStart" json:"gameSubmissionPeriodStart"`

	// Image パンフレット用画像
	Image *openapi_types.File `form:"image" json:"image,omitempty"`

	// Slug slug (URL内で使用, unique) 18th等
	Slug string `form:"slug" json:"slug"`

	// Title 第18回
	Title string `form:"title" json:"title"`
}

// PostGameRequest defines model for PostGameRequest.
type PostGameRequest struct {
	// CreatorName ゲーム作成
	CreatorName string `form:"creatorName" json:"creatorName"`

	// CreatorPageUrl ゲーム作成者のページのURL
	CreatorPageUrl *string `form:"creatorPageUrl" json:"creatorPageUrl,omitempty"`

	// Description ゲームの説明
	Description string `form:"description" json:"description"`

	// GamePageUrl ゲームページのURL
	GamePageUrl *string `form:"gamePageUrl" json:"gamePageUrl,omitempty"`

	// Icon ゲームのアイコン画像
	Icon openapi_types.File `form:"icon" json:"icon"`

	// Image ゲームの画像
	Image *openapi_types.File `form:"image" json:"image,omitempty"`

	// Title 展示するゲームタイトル
	Title string `form:"title" json:"title"`
}

// PostTermRequest defines model for PostTermRequest.
type PostTermRequest struct {
	// EndAt タームが終わる時間
	EndAt time.Time `json:"endAt"`

	// EventSlug イベントのslug
	EventSlug string `json:"eventSlug"`

	// StartAt タームが始まる時間
	StartAt time.Time `json:"startAt"`
}

// Term defines model for Term.
type Term struct {
	// EndAt タームが終わる時間
	EndAt time.Time `db:"end_at" json:"endAt"`

	// EventSlug イベントのslug
	EventSlug string `db:"event_slug" json:"eventSlug"`

	// Id タームのID
	Id openapi_types.UUID `db:"id" json:"id"`

	// IsDefault ゲーム登録時に割り当てられるTermならばTrue
	IsDefault bool `db:"is_default" json:"isDefault"`

	// StartAt タームが始まる時間
	StartAt time.Time `db:"start_at" json:"startAt"`
}

// User defines model for User.
type User struct {
	ProfileImageUrl string   `json:"profileImageUrl"`
	Role            UserRole `json:"role"`
	UserId          string   `json:"userId"`
	Username        string   `json:"username"`
}

// UserRole defines model for User.Role.
type UserRole string

// EventSlugInPath イベントSlug
type EventSlugInPath = string

// GameIdInPath ゲームID
type GameIdInPath = openapi_types.UUID

// TermIdInPath タームID
type TermIdInPath = openapi_types.UUID

// UserIdInPath ユーザーID
type UserIdInPath = string

// OauthCallbackParams defines parameters for OauthCallback.
type OauthCallbackParams struct {
	// Code Discordからの認証コード
	Code string `form:"code" json:"code"`
}

// LoginJSONBody defines parameters for Login.
type LoginJSONBody struct {
	// Redirect リダイレクト先のURL
	Redirect *string `json:"redirect,omitempty"`
}

// GetGamesParams defines parameters for GetGames.
type GetGamesParams struct {
	// TermId タームID
	TermId *openapi_types.UUID `form:"termId,omitempty" json:"termId,omitempty"`

	// EventSlug イベントID
	EventSlug *string `form:"eventSlug,omitempty" json:"eventSlug,omitempty"`

	// UserId ユーザーID
	UserId *string `form:"userId,omitempty" json:"userId,omitempty"`

	// Include 未公開のゲームを含むかどうか。includeに指定できる値は'unpublished'のみで、これが無ければ常にpublishedなGameのみ返す
	Include *string `form:"include,omitempty" json:"include,omitempty"`
}

// LoginJSONRequestBody defines body for Login for application/json ContentType.
type LoginJSONRequestBody LoginJSONBody

// PostEventMultipartRequestBody defines body for PostEvent for multipart/form-data ContentType.
type PostEventMultipartRequestBody = PostEventRequest

// PatchEventMultipartRequestBody defines body for PatchEvent for multipart/form-data ContentType.
type PatchEventMultipartRequestBody = PatchEventRequest

// PostGameMultipartRequestBody defines body for PostGame for multipart/form-data ContentType.
type PostGameMultipartRequestBody = PostGameRequest

// PatchGameMultipartRequestBody defines body for PatchGame for multipart/form-data ContentType.
type PatchGameMultipartRequestBody = PatchGameRequest

// PostTermJSONRequestBody defines body for PostTerm for application/json ContentType.
type PostTermJSONRequestBody = PostTermRequest

// PatchTermJSONRequestBody defines body for PatchTerm for application/json ContentType.
type PatchTermJSONRequestBody = PatchTermRequest
