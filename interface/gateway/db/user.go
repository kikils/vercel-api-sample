package db

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/kikils/vercel-api-sample/domain/model"
	"golang.org/x/xerrors"
)

type UserRepository struct {
	Repository
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		Repository: Repository{
			db: db,
		},
	}
}

func (r *UserRepository) Store(ctx context.Context, m *model.User) error {
	if err := r.Exec(`INSERT INTO users(id, name) VALUES($1, $2);`, m.ID, m.Name); err != nil {
		return xerrors.Errorf("UserRepository Store: %w", err)
	}
	return nil
}

func (r *UserRepository) SearchByID(ctx context.Context, id string) (*model.User, error) {
	users := make([]model.User, 1)
	if err := r.Query(&users, `SELECT * FROM users WHERE id = $1 LIMIT 1;`, id); err != nil {
		return nil, xerrors.Errorf("UserRepository SearchByID: %w", err)
	}
	if len(users) == 0 {
		return nil, xerrors.Errorf("UserReposiotry SearchByID: no record")
	}
	return &users[0], nil
}

func (r *UserRepository) UpdateByID(ctx context.Context, id string, user *model.User) error {
	if err := r.Exec(`UPDATE users SET name = $1 WHERE id = $2;`, user.Name, id); err != nil {
		return xerrors.Errorf("UserRepository UpdateByID: %w", err)
	}
	return nil
}

func (r *UserRepository) DeleteByID(ctx context.Context, id string) error {
	if err := r.Exec(`DELETE FROM users WHERE id = $1;`, id); err != nil {
		return xerrors.Errorf("UserRepository DeleteByID: %w", err)
	}
	return nil
}
