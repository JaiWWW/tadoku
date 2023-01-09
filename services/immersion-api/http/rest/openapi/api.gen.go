// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by unknown module path version unknown version DO NOT EDIT.
package openapi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/labstack/echo/v4"
)

const (
	CookieAuthScopes = "cookieAuth.Scopes"
)

// Activities defines model for Activities.
type Activities struct {
	Activities []Activity `json:"activities"`
}

// Activity defines model for Activity.
type Activity struct {
	Default *bool  `json:"default,omitempty"`
	Id      int32  `json:"id"`
	Name    string `json:"name"`
}

// Contest defines model for Contest.
type Contest struct {
	ActivityTypeIdAllowList []int32             `json:"activity_type_id_allow_list"`
	ContestEnd              openapi_types.Date  `json:"contest_end"`
	ContestStart            openapi_types.Date  `json:"contest_start"`
	CreatedAt               *time.Time          `json:"created_at,omitempty"`
	Deleted                 *bool               `json:"deleted,omitempty"`
	Description             string              `json:"description"`
	Id                      *openapi_types.UUID `json:"id,omitempty"`
	LanguageCodeAllowList   []string            `json:"language_code_allow_list"`
	Official                bool                `json:"official"`
	OwnerUserDisplayName    *string             `json:"owner_user_display_name,omitempty"`
	OwnerUserId             *openapi_types.UUID `json:"owner_user_id,omitempty"`
	Private                 bool                `json:"private"`
	RegistrationEnd         openapi_types.Date  `json:"registration_end"`
	UpdatedAt               *time.Time          `json:"updated_at,omitempty"`
}

// ContestBase defines model for ContestBase.
type ContestBase struct {
	ContestEnd           openapi_types.Date  `json:"contest_end"`
	ContestStart         openapi_types.Date  `json:"contest_start"`
	CreatedAt            *time.Time          `json:"created_at,omitempty"`
	Deleted              *bool               `json:"deleted,omitempty"`
	Description          string              `json:"description"`
	Id                   *openapi_types.UUID `json:"id,omitempty"`
	Official             bool                `json:"official"`
	OwnerUserDisplayName *string             `json:"owner_user_display_name,omitempty"`
	OwnerUserId          *openapi_types.UUID `json:"owner_user_id,omitempty"`
	Private              bool                `json:"private"`
	RegistrationEnd      openapi_types.Date  `json:"registration_end"`
	UpdatedAt            *time.Time          `json:"updated_at,omitempty"`
}

// ContestConfigurationOptions defines model for ContestConfigurationOptions.
type ContestConfigurationOptions struct {
	Activities             []Activity `json:"activities"`
	CanCreateOfficialRound bool       `json:"can_create_official_round"`
	Languages              []Language `json:"languages"`
}

// ContestRegistration defines model for ContestRegistration.
type ContestRegistration struct {
	Contest         *ContestView        `json:"contest,omitempty"`
	ContestId       openapi_types.UUID  `json:"contest_id"`
	Id              *openapi_types.UUID `json:"id,omitempty"`
	Languages       []Language          `json:"languages"`
	UserDisplayName string              `json:"user_display_name"`
	UserId          openapi_types.UUID  `json:"user_id"`
}

// ContestRegistrations defines model for ContestRegistrations.
type ContestRegistrations struct {
	// NextPageToken is empty if there's no next page
	NextPageToken string                `json:"next_page_token"`
	Registrations []ContestRegistration `json:"registrations"`
	TotalSize     int                   `json:"total_size"`
}

// ContestView defines model for ContestView.
type ContestView struct {
	AllowedActivities    []Activity          `json:"allowed_activities"`
	AllowedLanguages     []Language          `json:"allowed_languages"`
	ContestEnd           openapi_types.Date  `json:"contest_end"`
	ContestStart         openapi_types.Date  `json:"contest_start"`
	CreatedAt            *time.Time          `json:"created_at,omitempty"`
	Deleted              *bool               `json:"deleted,omitempty"`
	Description          string              `json:"description"`
	Id                   *openapi_types.UUID `json:"id,omitempty"`
	Official             bool                `json:"official"`
	OwnerUserDisplayName *string             `json:"owner_user_display_name,omitempty"`
	OwnerUserId          *openapi_types.UUID `json:"owner_user_id,omitempty"`
	Private              bool                `json:"private"`
	RegistrationEnd      openapi_types.Date  `json:"registration_end"`
	UpdatedAt            *time.Time          `json:"updated_at,omitempty"`
}

// Contests defines model for Contests.
type Contests struct {
	Contests []Contest `json:"contests"`

	// NextPageToken is empty if there's no next page
	NextPageToken string `json:"next_page_token"`
	TotalSize     int    `json:"total_size"`
}

// Language defines model for Language.
type Language struct {
	// Code In ISO-639-3 https://en.wikipedia.org/wiki/Wikipedia:WikiProject_Languages/List_of_ISO_639-3_language_codes_(2019)
	Code string `json:"code"`
	Name string `json:"name"`
}

// Languages defines model for Languages.
type Languages struct {
	Languages []Language `json:"languages"`
}

// Leaderboard defines model for Leaderboard.
type Leaderboard struct {
	Entries []LeaderboardEntry `json:"entries"`

	// NextPageToken is empty if there's no next page
	NextPageToken string `json:"next_page_token"`
	TotalSize     int    `json:"total_size"`
}

// LeaderboardEntry defines model for LeaderboardEntry.
type LeaderboardEntry struct {
	IsTie           bool               `json:"is_tie"`
	Rank            int                `json:"rank"`
	Score           int                `json:"score"`
	UserDisplayName string             `json:"user_display_name"`
	UserId          openapi_types.UUID `json:"user_id"`
}

// LogConfigurationOptions defines model for LogConfigurationOptions.
type LogConfigurationOptions struct {
	Activities []Activity `json:"activities"`
	Languages  []Language `json:"languages"`
	Tags       []Tag      `json:"tags"`
	Units      []Unit     `json:"units"`
}

// PaginatedList defines model for PaginatedList.
type PaginatedList struct {
	// NextPageToken is empty if there's no next page
	NextPageToken string `json:"next_page_token"`
	TotalSize     int    `json:"total_size"`
}

// Tag defines model for Tag.
type Tag struct {
	Id            openapi_types.UUID `json:"id"`
	LogActivityId int                `json:"log_activity_id"`
	Name          string             `json:"name"`
}

// Tags defines model for Tags.
type Tags struct {
	Tags []Tag `json:"tags"`
}

// Unit defines model for Unit.
type Unit struct {
	Id            openapi_types.UUID `json:"id"`
	LanguageCode  *string            `json:"language_code,omitempty"`
	LogActivityId int                `json:"log_activity_id"`
	Modifier      float32            `json:"modifier"`
	Name          string             `json:"name"`
}

// Units defines model for Units.
type Units struct {
	Units []Unit `json:"units"`
}

// ContestListParams defines parameters for ContestList.
type ContestListParams struct {
	PageSize       *int                `form:"page_size,omitempty" json:"page_size,omitempty"`
	Page           *int                `form:"page,omitempty" json:"page,omitempty"`
	IncludeDeleted *bool               `form:"include_deleted,omitempty" json:"include_deleted,omitempty"`
	Official       *bool               `form:"official,omitempty" json:"official,omitempty"`
	UserId         *openapi_types.UUID `form:"user_id,omitempty" json:"user_id,omitempty"`
}

// ContestFetchLeaderboardParams defines parameters for ContestFetchLeaderboard.
type ContestFetchLeaderboardParams struct {
	PageSize     *int    `form:"page_size,omitempty" json:"page_size,omitempty"`
	Page         *int    `form:"page,omitempty" json:"page,omitempty"`
	LanguageCode *string `form:"language_code,omitempty" json:"language_code,omitempty"`
	ActivityId   *int    `form:"activity_id,omitempty" json:"activity_id,omitempty"`
}

// ContestRegistrationUpsertJSONBody defines parameters for ContestRegistrationUpsert.
type ContestRegistrationUpsertJSONBody struct {
	LanguageCodes []string `json:"language_codes"`
}

// LogCreateLogJSONBody defines parameters for LogCreateLog.
type LogCreateLogJSONBody struct {
	ActivityId      openapi_types.UUID   `json:"activity_id"`
	Amount          float32              `json:"amount"`
	LanguageCode    string               `json:"language_code"`
	RegistrationIds []openapi_types.UUID `json:"registration_ids"`
	Tags            []string             `json:"tags"`
	UnitId          openapi_types.UUID   `json:"unit_id"`
}

// ContestCreateJSONRequestBody defines body for ContestCreate for application/json ContentType.
type ContestCreateJSONRequestBody = Contest

// ContestRegistrationUpsertJSONRequestBody defines body for ContestRegistrationUpsert for application/json ContentType.
type ContestRegistrationUpsertJSONRequestBody ContestRegistrationUpsertJSONBody

// LogCreateLogJSONRequestBody defines body for LogCreateLog for application/json ContentType.
type LogCreateLogJSONRequestBody LogCreateLogJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Lists all the contests, paginated
	// (GET /contests)
	ContestList(ctx echo.Context, params ContestListParams) error
	// Creates a new contest
	// (POST /contests)
	ContestCreate(ctx echo.Context) error
	// Fetches the configuration options for a new contest
	// (GET /contests/configuration-options)
	ContestGetConfigurations(ctx echo.Context) error
	// Fetches all the ongoing contest registrations of the logged in user, always in a single page
	// (GET /contests/ongoing-registrations)
	ContestFindOngoingRegistrations(ctx echo.Context) error
	// Fetches a contest by id
	// (GET /contests/{id})
	ContestFindByID(ctx echo.Context, id openapi_types.UUID) error
	// Fetches the leaderboard for a contest
	// (GET /contests/{id}/leaderboard)
	ContestFetchLeaderboard(ctx echo.Context, id openapi_types.UUID, params ContestFetchLeaderboardParams) error
	// Fetches a contest registration if it exists
	// (GET /contests/{id}/registration)
	ContestFindRegistration(ctx echo.Context, id openapi_types.UUID) error
	// Creates or updates a registration for a contest
	// (POST /contests/{id}/registration)
	ContestRegistrationUpsert(ctx echo.Context, id openapi_types.UUID) error
	// Submits a new log
	// (POST /logs)
	LogCreateLog(ctx echo.Context) error
	// Fetches the configuration options for a log
	// (GET /logs/configuration-options)
	LogGetConfigurations(ctx echo.Context) error
	// Checks if service is responsive
	// (GET /ping)
	Ping(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ContestList converts echo context to params.
func (w *ServerInterfaceWrapper) ContestList(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ContestListParams
	// ------------- Optional query parameter "page_size" -------------

	err = runtime.BindQueryParameter("form", true, false, "page_size", ctx.QueryParams(), &params.PageSize)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page_size: %s", err))
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "include_deleted" -------------

	err = runtime.BindQueryParameter("form", true, false, "include_deleted", ctx.QueryParams(), &params.IncludeDeleted)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter include_deleted: %s", err))
	}

	// ------------- Optional query parameter "official" -------------

	err = runtime.BindQueryParameter("form", true, false, "official", ctx.QueryParams(), &params.Official)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter official: %s", err))
	}

	// ------------- Optional query parameter "user_id" -------------

	err = runtime.BindQueryParameter("form", true, false, "user_id", ctx.QueryParams(), &params.UserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter user_id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ContestList(ctx, params)
	return err
}

// ContestCreate converts echo context to params.
func (w *ServerInterfaceWrapper) ContestCreate(ctx echo.Context) error {
	var err error

	ctx.Set(CookieAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ContestCreate(ctx)
	return err
}

// ContestGetConfigurations converts echo context to params.
func (w *ServerInterfaceWrapper) ContestGetConfigurations(ctx echo.Context) error {
	var err error

	ctx.Set(CookieAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ContestGetConfigurations(ctx)
	return err
}

// ContestFindOngoingRegistrations converts echo context to params.
func (w *ServerInterfaceWrapper) ContestFindOngoingRegistrations(ctx echo.Context) error {
	var err error

	ctx.Set(CookieAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ContestFindOngoingRegistrations(ctx)
	return err
}

// ContestFindByID converts echo context to params.
func (w *ServerInterfaceWrapper) ContestFindByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ContestFindByID(ctx, id)
	return err
}

// ContestFetchLeaderboard converts echo context to params.
func (w *ServerInterfaceWrapper) ContestFetchLeaderboard(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(CookieAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ContestFetchLeaderboardParams
	// ------------- Optional query parameter "page_size" -------------

	err = runtime.BindQueryParameter("form", true, false, "page_size", ctx.QueryParams(), &params.PageSize)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page_size: %s", err))
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "language_code" -------------

	err = runtime.BindQueryParameter("form", true, false, "language_code", ctx.QueryParams(), &params.LanguageCode)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter language_code: %s", err))
	}

	// ------------- Optional query parameter "activity_id" -------------

	err = runtime.BindQueryParameter("form", true, false, "activity_id", ctx.QueryParams(), &params.ActivityId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter activity_id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ContestFetchLeaderboard(ctx, id, params)
	return err
}

// ContestFindRegistration converts echo context to params.
func (w *ServerInterfaceWrapper) ContestFindRegistration(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(CookieAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ContestFindRegistration(ctx, id)
	return err
}

// ContestRegistrationUpsert converts echo context to params.
func (w *ServerInterfaceWrapper) ContestRegistrationUpsert(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(CookieAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ContestRegistrationUpsert(ctx, id)
	return err
}

// LogCreateLog converts echo context to params.
func (w *ServerInterfaceWrapper) LogCreateLog(ctx echo.Context) error {
	var err error

	ctx.Set(CookieAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.LogCreateLog(ctx)
	return err
}

// LogGetConfigurations converts echo context to params.
func (w *ServerInterfaceWrapper) LogGetConfigurations(ctx echo.Context) error {
	var err error

	ctx.Set(CookieAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.LogGetConfigurations(ctx)
	return err
}

// Ping converts echo context to params.
func (w *ServerInterfaceWrapper) Ping(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Ping(ctx)
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

	router.GET(baseURL+"/contests", wrapper.ContestList)
	router.POST(baseURL+"/contests", wrapper.ContestCreate)
	router.GET(baseURL+"/contests/configuration-options", wrapper.ContestGetConfigurations)
	router.GET(baseURL+"/contests/ongoing-registrations", wrapper.ContestFindOngoingRegistrations)
	router.GET(baseURL+"/contests/:id", wrapper.ContestFindByID)
	router.GET(baseURL+"/contests/:id/leaderboard", wrapper.ContestFetchLeaderboard)
	router.GET(baseURL+"/contests/:id/registration", wrapper.ContestFindRegistration)
	router.POST(baseURL+"/contests/:id/registration", wrapper.ContestRegistrationUpsert)
	router.POST(baseURL+"/logs", wrapper.LogCreateLog)
	router.GET(baseURL+"/logs/configuration-options", wrapper.LogGetConfigurations)
	router.GET(baseURL+"/ping", wrapper.Ping)

}
