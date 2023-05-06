package db

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Exec(q string, params ...any) error {
	_, err := r.db.Exec(q, params...)
	return err
}

func (r *Repository) Query(t any, q string, params ...any) error {
	err := r.db.Select(t, q, params...)
	return err
}
