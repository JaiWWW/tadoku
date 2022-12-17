// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: pages.sql

package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createPage = `-- name: CreatePage :one
insert into pages (
  id,
  slug,
  current_content_id,
  published_at
) values (
  $1,
  $2,
  $3,
  $4
) returning id
`

type CreatePageParams struct {
	ID               uuid.UUID
	Slug             string
	CurrentContentID uuid.UUID
	PublishedAt      sql.NullTime
}

func (q *Queries) CreatePage(ctx context.Context, arg CreatePageParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createPage,
		arg.ID,
		arg.Slug,
		arg.CurrentContentID,
		arg.PublishedAt,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const createPageContent = `-- name: CreatePageContent :one
insert into pages_content (
  id,
  page_id,
  title,
  html
) values (
  $1,
  $2,
  $3,
  $4
) returning id
`

type CreatePageContentParams struct {
	ID     uuid.UUID
	PageID uuid.UUID
	Title  string
	Html   string
}

func (q *Queries) CreatePageContent(ctx context.Context, arg CreatePageContentParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createPageContent,
		arg.ID,
		arg.PageID,
		arg.Title,
		arg.Html,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const findPageBySlug = `-- name: FindPageBySlug :one
select
  pages.id,
  slug,
  pages_content.title,
  pages_content.html,
  published_at
from pages
inner join pages_content
  on pages_content.id = pages.current_content_id
where
  deleted_at is null
  and slug = $1
`

type FindPageBySlugRow struct {
	ID          uuid.UUID
	Slug        string
	Title       string
	Html        string
	PublishedAt sql.NullTime
}

func (q *Queries) FindPageBySlug(ctx context.Context, slug string) (FindPageBySlugRow, error) {
	row := q.db.QueryRowContext(ctx, findPageBySlug, slug)
	var i FindPageBySlugRow
	err := row.Scan(
		&i.ID,
		&i.Slug,
		&i.Title,
		&i.Html,
		&i.PublishedAt,
	)
	return i, err
}
