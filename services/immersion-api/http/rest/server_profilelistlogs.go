package rest

import (
	"errors"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/labstack/echo/v4"
	"github.com/tadoku/tadoku/services/immersion-api/domain/query"
	"github.com/tadoku/tadoku/services/immersion-api/http/rest/openapi"
)

// Lists the logs of a user
// (GET /users/{user_id}/logs)
func (s *Server) ProfileListLogs(ctx echo.Context, userId types.UUID, params openapi.ProfileListLogsParams) error {
	req := &query.ListLogsForUserRequest{
		UserID:         userId,
		IncludeDeleted: false,
		PageSize:       0,
		Page:           0,
	}

	if params.PageSize != nil {
		req.PageSize = *params.PageSize
	}
	if params.Page != nil {
		req.Page = *params.Page
	}
	if params.IncludeDeleted != nil {
		req.IncludeDeleted = *params.IncludeDeleted
	}

	list, err := s.queryService.ListLogsForUser(ctx.Request().Context(), req)
	if err != nil {
		if errors.Is(err, query.ErrUnauthorized) {
			return ctx.NoContent(http.StatusForbidden)
		}

		ctx.Echo().Logger.Error(err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	res := openapi.Logs{
		Logs:          make([]openapi.Log, len(list.Logs)),
		NextPageToken: list.NextPageToken,
		TotalSize:     list.TotalSize,
	}

	for i, it := range list.Logs {
		it := it
		res.Logs[i] = openapi.Log{
			Id: it.ID,
			Activity: openapi.Activity{
				Id:   int32(it.ActivityID),
				Name: it.ActivityName,
			},
			Language: openapi.Language{
				Code: it.LanguageCode,
				Name: it.LanguageName,
			},
			Amount:      it.Amount,
			Modifier:    it.Modifier,
			Score:       it.Score,
			Tags:        it.Tags,
			UnitName:    it.UnitName,
			UserId:      it.UserID,
			CreatedAt:   it.CreatedAt,
			Deleted:     it.Deleted,
			Description: it.Description,
		}
	}

	return ctx.JSON(http.StatusOK, res)
}
