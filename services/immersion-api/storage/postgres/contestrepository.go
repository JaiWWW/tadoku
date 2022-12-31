package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/tadoku/tadoku/services/immersion-api/domain/contestcommand"
	"github.com/tadoku/tadoku/services/immersion-api/domain/contestquery"
)

type ContestRepository struct {
	psql *sql.DB
	q    *Queries
}

func NewContestRepository(psql *sql.DB) *ContestRepository {
	return &ContestRepository{
		psql: psql,
		q:    &Queries{psql},
	}
}

// COMMANDS

func (r *ContestRepository) CreateContest(ctx context.Context, req *contestcommand.ContestCreateRequest) (*contestcommand.ContestCreateResponse, error) {
	tx, err := r.psql.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create contest: %w", err)
	}

	qtx := r.q.WithTx(tx)

	id, err := qtx.CreateContest(ctx, CreateContestParams{
		OwnerUserID:             req.OwnerUserID,
		OwnerUserDisplayName:    req.OwnerUserDisplayName,
		Official:                req.Official,
		Private:                 req.Private,
		ContestStart:            req.ContestStart,
		ContestEnd:              req.ContestEnd,
		RegistrationStart:       req.RegistrationStart,
		RegistrationEnd:         req.RegistrationEnd,
		Description:             req.Description,
		LanguageCodeAllowList:   req.LanguageCodeAllowList,
		ActivityTypeIDAllowList: req.ActivityTypeIDAllowList,
	})

	if err != nil {
		_ = tx.Rollback()
		return nil, fmt.Errorf("could not create contest: %w", err)
	}

	contest, err := qtx.FindContestById(ctx, id)
	if err != nil {
		_ = tx.Rollback()
		return nil, fmt.Errorf("could not create contest: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("could not create contest: %w", err)
	}

	return &contestcommand.ContestCreateResponse{
		ID:                      contest.ID,
		ContestStart:            contest.ContestStart,
		ContestEnd:              contest.ContestEnd,
		RegistrationStart:       contest.RegistrationStart,
		RegistrationEnd:         contest.RegistrationEnd,
		Description:             contest.Description,
		OwnerUserID:             contest.OwnerUserID,
		OwnerUserDisplayName:    contest.OwnerUserDisplayName,
		Official:                contest.Official,
		Private:                 contest.Private,
		LanguageCodeAllowList:   contest.LanguageCodeAllowList,
		ActivityTypeIDAllowList: contest.ActivityTypeIDAllowList,
		CreatedAt:               contest.CreatedAt,
		UpdatedAt:               contest.UpdatedAt,
	}, nil
}

// QUERIES

func (r *ContestRepository) FetchContestConfigurationOptions(ctx context.Context) (*contestquery.FetchContestConfigurationOptionsResponse, error) {
	langs, err := r.q.ListLanguages(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not fetch contest configuration options: %w", err)
	}

	acts, err := r.q.ListActivities(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not fetch contest configuration options: %w", err)
	}

	options := contestquery.FetchContestConfigurationOptionsResponse{
		Languages:  make([]contestquery.Language, len(langs)),
		Activities: make([]contestquery.Activity, len(acts)),
	}

	for i, l := range langs {
		options.Languages[i] = contestquery.Language{
			Code: l.Code,
			Name: l.Name,
		}
	}

	for i, a := range acts {
		options.Activities[i] = contestquery.Activity{
			ID:      a.ID,
			Name:    a.Name,
			Default: a.Default,
		}
	}

	return &options, err
}

func (r *ContestRepository) ListContests(ctx context.Context, req *contestquery.ContestListRequest) (*contestquery.ContestListResponse, error) {
	tx, err := r.psql.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("could not list contests: %w", err)
	}

	qtx := r.q.WithTx(tx)

	meta, err := qtx.ContestsMetadata(ctx, ContestsMetadataParams{
		IncludeDeleted: req.IncludeDeleted,
		UserID:         req.UserID,
		Official:       req.OfficialOnly,
	})
	if err != nil {
		_ = tx.Rollback()
		return nil, fmt.Errorf("could not lists contests: %w", err)
	}

	contests, err := qtx.ListContests(ctx, ListContestsParams{
		StartFrom:      int32(req.Page * req.PageSize),
		PageSize:       int32(req.PageSize),
		IncludeDeleted: req.IncludeDeleted,
		UserID:         req.UserID,
		Official:       req.OfficialOnly,
	})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		_ = tx.Rollback()
		return nil, fmt.Errorf("could not list contests: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("could not list contests: %w", err)
	}

	res := make([]contestquery.ContestListEntry, len(contests))
	for i, c := range contests {
		res[i] = contestquery.ContestListEntry{
			ID:                      c.ID,
			ContestStart:            c.ContestStart,
			ContestEnd:              c.ContestEnd,
			RegistrationStart:       c.RegistrationStart,
			RegistrationEnd:         c.RegistrationEnd,
			Description:             c.Description,
			OwnerUserID:             c.OwnerUserID,
			OwnerUserDisplayName:    c.OwnerUserDisplayName,
			Official:                c.Official,
			Private:                 c.Private,
			LanguageCodeAllowList:   c.LanguageCodeAllowList,
			ActivityTypeIDAllowList: c.ActivityTypeIDAllowList,
			CreatedAt:               c.CreatedAt,
			UpdatedAt:               c.UpdatedAt,
			Deleted:                 c.DeletedAt.Valid,
		}
	}

	nextPageToken := ""
	if (req.Page*req.PageSize)+req.PageSize < int(meta.TotalSize) {
		nextPageToken = fmt.Sprint(req.Page + 1)
	}

	return &contestquery.ContestListResponse{
		Contests:      res,
		TotalSize:     int(meta.TotalSize),
		NextPageToken: nextPageToken,
	}, nil
}
