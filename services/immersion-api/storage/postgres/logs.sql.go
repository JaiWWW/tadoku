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

const checkIfLogCanBeDeleted = `-- name: CheckIfLogCanBeDeleted :one
select (not(true = any(
  select
    (contests.contest_end < $1)
  from contest_logs
  inner join contests on (contests.id = contest_logs.contest_id)
  where
    contest_logs.log_id = $2
)))::boolean as can_be_deleted
`

type CheckIfLogCanBeDeletedParams struct {
	Now   time.Time
	LogID uuid.UUID
}

func (q *Queries) CheckIfLogCanBeDeleted(ctx context.Context, arg CheckIfLogCanBeDeletedParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkIfLogCanBeDeleted, arg.Now, arg.LogID)
	var can_be_deleted bool
	err := row.Scan(&can_be_deleted)
	return can_be_deleted, err
}

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

const deleteLog = `-- name: DeleteLog :exec
update logs
set deleted_at = now()
where
  id = $1
  and deleted_at is null
`

func (q *Queries) DeleteLog(ctx context.Context, logID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteLog, logID)
	return err
}

const fetchScoresForProfile = `-- name: FetchScoresForProfile :many
select
  language_code,
  sum(score)::real as score,
  languages.name as language_name
from logs
inner join languages on (languages.code = logs.language_code)
where
  user_id = $1
  and year = $2
  and deleted_at is null
group by language_code, languages.name
order by score desc
`

type FetchScoresForProfileParams struct {
	UserID uuid.UUID
	Year   int16
}

type FetchScoresForProfileRow struct {
	LanguageCode string
	Score        float32
	LanguageName string
}

func (q *Queries) FetchScoresForProfile(ctx context.Context, arg FetchScoresForProfileParams) ([]FetchScoresForProfileRow, error) {
	rows, err := q.db.QueryContext(ctx, fetchScoresForProfile, arg.UserID, arg.Year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FetchScoresForProfileRow
	for rows.Next() {
		var i FetchScoresForProfileRow
		if err := rows.Scan(&i.LanguageCode, &i.Score, &i.LanguageName); err != nil {
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

const findAttachedContestRegistrationsForLog = `-- name: FindAttachedContestRegistrationsForLog :many
select
  contest_logs.contest_id,
  contests.title,
  contest_registrations.id,
  contests.contest_end
from contest_logs
inner join contests on (contests.id = contest_logs.contest_id)
inner join logs on (logs.id = contest_logs.log_id)
inner join contest_registrations on (
  contest_registrations.contest_id = contest_logs.contest_id
  and contest_registrations.user_id = logs.user_id
)
where log_id = $1
`

type FindAttachedContestRegistrationsForLogRow struct {
	ContestID  uuid.UUID
	Title      string
	ID         uuid.UUID
	ContestEnd time.Time
}

func (q *Queries) FindAttachedContestRegistrationsForLog(ctx context.Context, id uuid.UUID) ([]FindAttachedContestRegistrationsForLogRow, error) {
	rows, err := q.db.QueryContext(ctx, findAttachedContestRegistrationsForLog, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindAttachedContestRegistrationsForLogRow
	for rows.Next() {
		var i FindAttachedContestRegistrationsForLogRow
		if err := rows.Scan(
			&i.ContestID,
			&i.Title,
			&i.ID,
			&i.ContestEnd,
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

const findLogByID = `-- name: FindLogByID :one
select
  logs.id,
  logs.user_id,
  users.display_name as user_display_name,
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
from logs
inner join languages on (languages.code = logs.language_code)
inner join log_activities on (log_activities.id = logs.log_activity_id)
inner join log_units on (log_units.id = logs.unit_id)
inner join users on (users.id = logs.user_id)
where
  ($1::boolean or deleted_at is null)
  and logs.id = $2
`

type FindLogByIDParams struct {
	IncludeDeleted bool
	ID             uuid.UUID
}

type FindLogByIDRow struct {
	ID              uuid.UUID
	UserID          uuid.UUID
	UserDisplayName string
	LanguageCode    string
	LanguageName    string
	ActivityID      int16
	ActivityName    string
	UnitName        string
	Description     sql.NullString
	Tags            []string
	Amount          float32
	Modifier        float32
	Score           float32
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       sql.NullTime
}

func (q *Queries) FindLogByID(ctx context.Context, arg FindLogByIDParams) (FindLogByIDRow, error) {
	row := q.db.QueryRowContext(ctx, findLogByID, arg.IncludeDeleted, arg.ID)
	var i FindLogByIDRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.UserDisplayName,
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
	)
	return i, err
}

const listLogsForContest = `-- name: ListLogsForContest :many
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
    logs.deleted_at,
    contest_registrations.user_display_name
  from contest_logs
  inner join logs on (logs.id = contest_logs.log_id)
  inner join languages on (languages.code = logs.language_code)
  inner join log_activities on (log_activities.id = logs.log_activity_id)
  inner join log_units on (log_units.id = logs.unit_id)
  inner join contest_registrations on (
    contest_registrations.contest_id = $3
    and contest_registrations.user_id = logs.user_id
  )
  where
    ($4::boolean or logs.deleted_at is null)
    and (logs.user_id = $5 or $5 is null)
    and contest_logs.contest_id = $3
)
select
  id, user_id, language_code, language_name, activity_id, activity_name, unit_name, description, tags, amount, modifier, score, created_at, updated_at, deleted_at, user_display_name,
  (select count(eligible_logs.id) from eligible_logs) as total_size
from eligible_logs
order by created_at desc
limit $2
offset $1
`

type ListLogsForContestParams struct {
	StartFrom      int32
	PageSize       int32
	ContestID      uuid.UUID
	IncludeDeleted bool
	UserID         uuid.NullUUID
}

type ListLogsForContestRow struct {
	ID              uuid.UUID
	UserID          uuid.UUID
	LanguageCode    string
	LanguageName    string
	ActivityID      int16
	ActivityName    string
	UnitName        string
	Description     sql.NullString
	Tags            []string
	Amount          float32
	Modifier        float32
	Score           float32
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       sql.NullTime
	UserDisplayName string
	TotalSize       int64
}

func (q *Queries) ListLogsForContest(ctx context.Context, arg ListLogsForContestParams) ([]ListLogsForContestRow, error) {
	rows, err := q.db.QueryContext(ctx, listLogsForContest,
		arg.StartFrom,
		arg.PageSize,
		arg.ContestID,
		arg.IncludeDeleted,
		arg.UserID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListLogsForContestRow
	for rows.Next() {
		var i ListLogsForContestRow
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
			&i.UserDisplayName,
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

const listLogsForUser = `-- name: ListLogsForUser :many
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
  from logs
  inner join languages on (languages.code = logs.language_code)
  inner join log_activities on (log_activities.id = logs.log_activity_id)
  inner join log_units on (log_units.id = logs.unit_id)
  where
    ($3::boolean or deleted_at is null)
    and logs.user_id = $4
)
select
  id, user_id, language_code, language_name, activity_id, activity_name, unit_name, description, tags, amount, modifier, score, created_at, updated_at, deleted_at,
  (select count(eligible_logs.id) from eligible_logs) as total_size
from eligible_logs
order by created_at desc
limit $2
offset $1
`

type ListLogsForUserParams struct {
	StartFrom      int32
	PageSize       int32
	IncludeDeleted bool
	UserID         uuid.UUID
}

type ListLogsForUserRow struct {
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

func (q *Queries) ListLogsForUser(ctx context.Context, arg ListLogsForUserParams) ([]ListLogsForUserRow, error) {
	rows, err := q.db.QueryContext(ctx, listLogsForUser,
		arg.StartFrom,
		arg.PageSize,
		arg.IncludeDeleted,
		arg.UserID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListLogsForUserRow
	for rows.Next() {
		var i ListLogsForUserRow
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

const yearlyActivityForUser = `-- name: YearlyActivityForUser :many
select
  sum(score)::real as score,
  count(id) as update_count,
  created_at::date as "date"
from logs
where
  user_id = $1
  and year = $2
  and deleted_at is null
group by "date"
order by date asc
`

type YearlyActivityForUserParams struct {
	UserID uuid.UUID
	Year   int16
}

type YearlyActivityForUserRow struct {
	Score       float32
	UpdateCount int64
	Date        time.Time
}

func (q *Queries) YearlyActivityForUser(ctx context.Context, arg YearlyActivityForUserParams) ([]YearlyActivityForUserRow, error) {
	rows, err := q.db.QueryContext(ctx, yearlyActivityForUser, arg.UserID, arg.Year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []YearlyActivityForUserRow
	for rows.Next() {
		var i YearlyActivityForUserRow
		if err := rows.Scan(&i.Score, &i.UpdateCount, &i.Date); err != nil {
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

const yearlyActivitySplitForUser = `-- name: YearlyActivitySplitForUser :many
select
  sum(logs.score)::real as score,
  logs.log_activity_id,
  log_activities.name as log_activity_name
from logs
inner join log_activities on (log_activities.id = logs.log_activity_id)
where
  user_id = $1
  and year = $2
  and deleted_at is null
group by logs.log_activity_id, log_activities.name
order by score desc
`

type YearlyActivitySplitForUserParams struct {
	UserID uuid.UUID
	Year   int16
}

type YearlyActivitySplitForUserRow struct {
	Score           float32
	LogActivityID   int16
	LogActivityName string
}

func (q *Queries) YearlyActivitySplitForUser(ctx context.Context, arg YearlyActivitySplitForUserParams) ([]YearlyActivitySplitForUserRow, error) {
	rows, err := q.db.QueryContext(ctx, yearlyActivitySplitForUser, arg.UserID, arg.Year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []YearlyActivitySplitForUserRow
	for rows.Next() {
		var i YearlyActivitySplitForUserRow
		if err := rows.Scan(&i.Score, &i.LogActivityID, &i.LogActivityName); err != nil {
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
