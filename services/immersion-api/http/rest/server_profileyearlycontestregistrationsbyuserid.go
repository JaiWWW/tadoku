package rest

import (
	"errors"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/labstack/echo/v4"
	"github.com/tadoku/tadoku/services/immersion-api/domain/query"
	"github.com/tadoku/tadoku/services/immersion-api/http/rest/openapi"
)

// Fetches the contest registrations of a user for a given year
// (GET /users/{userId}/contest-registrations/{year})
func (s *Server) ProfileYearlyContestRegistrationsByUserID(ctx echo.Context, userId types.UUID, year int) error {
	regs, err := s.queryService.YearlyContestRegistrationsForUser(ctx.Request().Context(), &query.YearlyContestRegistrationsForUserRequest{
		UserID: userId,
		Year:   year,
	})
	if err != nil {
		if errors.Is(err, query.ErrUnauthorized) {
			return ctx.NoContent(http.StatusUnauthorized)
		}

		ctx.Echo().Logger.Error(err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	res := &openapi.ContestRegistrations{
		TotalSize:     regs.TotalSize,
		NextPageToken: regs.NextPageToken,
		Registrations: make([]openapi.ContestRegistration, len(regs.Registrations)),
	}

	for i, it := range regs.Registrations {
		it := it
		res.Registrations[i] = *contestRegistrationToAPI(&it)
	}

	return ctx.JSON(http.StatusOK, res)
}
