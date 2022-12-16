// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package postgres

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Page struct {
	ID               uuid.UUID
	Slug             string
	CurrentContentID uuid.UUID
	PublishedAt      sql.NullTime
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        sql.NullTime
}

type PagesContent struct {
	ID        uuid.UUID
	PageID    uuid.UUID
	Title     string
	Html      string
	CreatedAt time.Time
}
