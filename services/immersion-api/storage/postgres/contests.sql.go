// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: contests.sql

package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const cancelContest = `-- name: CancelContest :one
update contests
set deleted_at = now()
where
  id = $1
  and deleted_at is null
returning id
`

func (q *Queries) CancelContest(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, cancelContest, id)
	err := row.Scan(&id)
	return id, err
}

const createContest = `-- name: CreateContest :one
insert into contests (
  owner_user_id,
  owner_user_display_name,
  official,
  "private",
  contest_start,
  contest_end,
  registration_start,
  registration_end,
  "description",
  language_code_allow_list,
  activity_type_id_allow_list
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
  $10,
  $11
) returning id
`

type CreateContestParams struct {
	OwnerUserID             uuid.UUID
	OwnerUserDisplayName    string
	Official                bool
	Private                 bool
	ContestStart            time.Time
	ContestEnd              time.Time
	RegistrationStart       time.Time
	RegistrationEnd         time.Time
	Description             string
	LanguageCodeAllowList   []string
	ActivityTypeIDAllowList []int32
}

func (q *Queries) CreateContest(ctx context.Context, arg CreateContestParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createContest,
		arg.OwnerUserID,
		arg.OwnerUserDisplayName,
		arg.Official,
		arg.Private,
		arg.ContestStart,
		arg.ContestEnd,
		arg.RegistrationStart,
		arg.RegistrationEnd,
		arg.Description,
		pq.Array(arg.LanguageCodeAllowList),
		pq.Array(arg.ActivityTypeIDAllowList),
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const findContestById = `-- name: FindContestById :one
select
  id,
  owner_user_id,
  owner_user_display_name,
  "private",
  contest_start,
  contest_end,
  registration_start,
  registration_end,
  "description",
  language_code_allow_list,
  activity_type_id_allow_list,
  official,
  created_at,
  updated_at,
  deleted_at
from contests
where
  id = $1
  and deleted_at is null
order by created_at desc
`

func (q *Queries) FindContestById(ctx context.Context, id uuid.UUID) (Contest, error) {
	row := q.db.QueryRowContext(ctx, findContestById, id)
	var i Contest
	err := row.Scan(
		&i.ID,
		&i.OwnerUserID,
		&i.OwnerUserDisplayName,
		&i.Private,
		&i.ContestStart,
		&i.ContestEnd,
		&i.RegistrationStart,
		&i.RegistrationEnd,
		&i.Description,
		pq.Array(&i.LanguageCodeAllowList),
		pq.Array(&i.ActivityTypeIDAllowList),
		&i.Official,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listContests = `-- name: ListContests :many
select
  id,
  owner_user_id,
  owner_user_display_name,
  "private",
  contest_start,
  contest_end,
  registration_start,
  registration_end,
  "description",
  language_code_allow_list,
  activity_type_id_allow_list,
  official,
  created_at,
  updated_at,
  deleted_at
from contests
where
  ($1::boolean or deleted_at is null)
  and (owner_user_id = $2 or $2 is null)
  and (official = $3)
order by created_at desc
limit $5
offset $4
`

type ListContestsParams struct {
	IncludeDeleted bool
	UserID         uuid.NullUUID
	Official       bool
	StartFrom      int32
	PageSize       int32
}

func (q *Queries) ListContests(ctx context.Context, arg ListContestsParams) ([]Contest, error) {
	rows, err := q.db.QueryContext(ctx, listContests,
		arg.IncludeDeleted,
		arg.UserID,
		arg.Official,
		arg.StartFrom,
		arg.PageSize,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Contest
	for rows.Next() {
		var i Contest
		if err := rows.Scan(
			&i.ID,
			&i.OwnerUserID,
			&i.OwnerUserDisplayName,
			&i.Private,
			&i.ContestStart,
			&i.ContestEnd,
			&i.RegistrationStart,
			&i.RegistrationEnd,
			&i.Description,
			pq.Array(&i.LanguageCodeAllowList),
			pq.Array(&i.ActivityTypeIDAllowList),
			&i.Official,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
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

const updateContest = `-- name: UpdateContest :one
update contests
set
  "private" = $1,
  contest_start = $2,
  contest_end = $3,
  registration_start = $4,
  registration_end = $5,
  "description" = $6,
  updated_at = now()
where
  id = $7
  and deleted_at is null
returning id
`

type UpdateContestParams struct {
	Private           bool
	ContestStart      time.Time
	ContestEnd        time.Time
	RegistrationStart time.Time
	RegistrationEnd   time.Time
	Description       string
	ID                uuid.UUID
}

func (q *Queries) UpdateContest(ctx context.Context, arg UpdateContestParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, updateContest,
		arg.Private,
		arg.ContestStart,
		arg.ContestEnd,
		arg.RegistrationStart,
		arg.RegistrationEnd,
		arg.Description,
		arg.ID,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}
