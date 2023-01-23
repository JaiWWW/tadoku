package rest

import (
	"errors"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/labstack/echo/v4"
	"github.com/tadoku/tadoku/services/immersion-api/domain/query"
	"github.com/tadoku/tadoku/services/immersion-api/http/rest/openapi"
)

// Fetches the scores of a user profile in a contest
// (GET /contests/{id}/profile/{user_id}/scores)
func (s *Server) ContestProfileFetchScores(ctx echo.Context, id types.UUID, userId types.UUID) error {
	profile, err := s.queryService.ContestProfile(ctx.Request().Context(), &query.ContestProfileRequest{
		UserID:    userId,
		ContestID: id,
	})
	if err != nil {
		if errors.Is(err, query.ErrNotFound) {
			return ctx.NoContent(http.StatusNotFound)
		}
		ctx.Logger().Errorf("could not fetch profile: %w", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	scores := make([]openapi.Score, len(profile.Scores))
	for i, it := range profile.Scores {
		scores[i] = openapi.Score{
			LanguageCode: it.LanguageCode,
			Score:        it.Score,
		}
	}

	return ctx.JSON(http.StatusOK, &openapi.ContestProfileScores{
		OverallScore: profile.OverallScore,
		Registration: *contestRegistrationToAPI(profile.Registration),
		Scores:       scores,
	})
}
