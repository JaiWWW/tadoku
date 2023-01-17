// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: logs.sql

package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const createContestLogRelation = `-- name: CreateContestLogRelation :exec
insert into contest_logs (
  contest_id,
  log_id
) values (
  (select contest_id from contest_registrations where id = $1),
  $2
)
`

type CreateContestLogRelationParams struct {
	RegistrationID uuid.UUID
	LogID          uuid.UUID
}

func (q *Queries) CreateContestLogRelation(ctx context.Context, arg CreateContestLogRelationParams) error {
	_, err := q.db.ExecContext(ctx, createContestLogRelation, arg.RegistrationID, arg.LogID)
	return err
}

const createLog = `-- name: CreateLog :one
insert into logs (
  id,
  user_id,
  language_code,
  log_activity_id,
  unit_id,
  tags,
  amount,
  modifier,
  eligible_official_leaderboard,
  "description"
) values (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7,
  $8,
  $9,
  $10
) returning id
`

type CreateLogParams struct {
	ID                          uuid.UUID
	UserID                      uuid.UUID
	LanguageCode                string
	LogActivityID               int16
	UnitID                      uuid.UUID
	Tags                        []string
	Amount                      float32
	Modifier                    float32
	EligibleOfficialLeaderboard bool
	Description                 sql.NullString
}

func (q *Queries) CreateLog(ctx context.Context, arg CreateLogParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createLog,
		arg.ID,
		arg.UserID,
		arg.LanguageCode,
		arg.LogActivityID,
		arg.UnitID,
		pq.Array(arg.Tags),
		arg.Amount,
		arg.Modifier,
		arg.EligibleOfficialLeaderboard,
		arg.Description,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const listLogsForContestUser = `-- name: ListLogsForContestUser :many
with eligible_logs as (
  select
    logs.id,
    logs.user_id,
    logs.language_code,
    languages.name as language_name,
    logs.log_activity_id as activity_id,
    log_activities.name as activity_name,
    log_units.name as unit_name,
    logs.description,
    logs.tags,
    logs.amount,
    logs.modifier,
    logs.score,
    logs.created_at,
    logs.updated_at,
    logs.deleted_at
  from contest_logs
  inner join logs on (logs.id = contest_logs.log_id)
  inner join languages on (languages.code = logs.language_code)
  inner join log_activities on (log_activities.id = logs.log_activity_id)
  inner join log_units on (log_units.id = logs.unit_id)
  where
    ($3::boolean or deleted_at is null)
    and logs.user_id = $4
    and contest_logs.contest_id = $5
)
select
  id, user_id, language_code, language_name, activity_id, activity_name, unit_name, description, tags, amount, modifier, score, created_at, updated_at, deleted_at,
  (select count(eligible_logs.id) from eligible_logs) as total_size
from eligible_logs
order by created_at desc
limit $2
offset $1
`

type ListLogsForContestUserParams struct {
	StartFrom      int32
	PageSize       int32
	IncludeDeleted bool
	UserID         uuid.UUID
	ContestID      uuid.UUID
}

type ListLogsForContestUserRow struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	LanguageCode string
	LanguageName string
	ActivityID   int16
	ActivityName string
	UnitName     string
	Description  sql.NullString
	Tags         []string
	Amount       float32
	Modifier     float32
	Score        float32
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime
	TotalSize    int64
}

func (q *Queries) ListLogsForContestUser(ctx context.Context, arg ListLogsForContestUserParams) ([]ListLogsForContestUserRow, error) {
	rows, err := q.db.QueryContext(ctx, listLogsForContestUser,
		arg.StartFrom,
		arg.PageSize,
		arg.IncludeDeleted,
		arg.UserID,
		arg.ContestID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListLogsForContestUserRow
	for rows.Next() {
		var i ListLogsForContestUserRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.LanguageCode,
			&i.LanguageName,
			&i.ActivityID,
			&i.ActivityName,
			&i.UnitName,
			&i.Description,
			pq.Array(&i.Tags),
			&i.Amount,
			&i.Modifier,
			&i.Score,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.TotalSize,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
