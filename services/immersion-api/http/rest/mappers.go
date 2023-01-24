package rest

import (
	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/tadoku/tadoku/services/immersion-api/domain/query"
	"github.com/tadoku/tadoku/services/immersion-api/http/rest/openapi"
)

func logToAPI(log *query.Log) *openapi.Log {
	refs := make([]openapi.ContestRegistrationReference, len(log.Registrations))
	for i, it := range log.Registrations {
		refs[i] = openapi.ContestRegistrationReference{
			ContestId:      it.ContestID,
			ContestEnd:     types.Date{Time: it.ContestEnd},
			RegistrationId: it.RegistrationID,
			Title:          it.Title,
		}
	}

	return &openapi.Log{
		Id: log.ID,
		Activity: openapi.Activity{
			Id:   int32(log.ActivityID),
			Name: log.ActivityName,
		},
		Language: openapi.Language{
			Code: log.LanguageCode,
			Name: log.LanguageName,
		},
		Amount:          log.Amount,
		Modifier:        log.Modifier,
		Score:           log.Score,
		Tags:            log.Tags,
		UnitName:        log.UnitName,
		UserId:          log.UserID,
		UserDisplayName: log.UserDisplayName,
		CreatedAt:       log.CreatedAt,
		Deleted:         log.Deleted,
		Description:     log.Description,
		Registrations:   &refs,
	}
}

func contestRegistrationToAPI(r *query.ContestRegistration) *openapi.ContestRegistration {
	registration := openapi.ContestRegistration{
		ContestId:       r.ContestID,
		Id:              &r.ID,
		Languages:       make([]openapi.Language, len(r.Languages)),
		UserId:          r.UserID,
		UserDisplayName: r.UserDisplayName,
	}

	if r.Contest != nil {
		contest := openapi.ContestView{
			Id:                &r.ContestID,
			ContestStart:      types.Date{Time: r.Contest.ContestStart},
			ContestEnd:        types.Date{Time: r.Contest.ContestEnd},
			RegistrationEnd:   types.Date{Time: r.Contest.RegistrationEnd},
			Title:             r.Contest.Title,
			Description:       r.Contest.Description,
			Official:          r.Contest.Official,
			Private:           r.Contest.Private,
			AllowedLanguages:  []openapi.Language{},
			AllowedActivities: make([]openapi.Activity, len(r.Contest.AllowedActivities)),
		}

		for i, a := range r.Contest.AllowedActivities {
			contest.AllowedActivities[i] = openapi.Activity{
				Id:   a.ID,
				Name: a.Name,
			}
		}

		registration.Contest = &contest
	}

	for i, lang := range r.Languages {
		registration.Languages[i] = openapi.Language{
			Code: lang.Code,
			Name: lang.Name,
		}
	}

	return &registration
}
