package auth

import (
	"bistleague-be/model/config"
	"bistleague-be/model/entity"
	"context"
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	cfg *config.Config
	db  *sqlx.DB
	qb  *goqu.DialectWrapper
}

func New(cfg *config.Config, db *sqlx.DB, qb *goqu.DialectWrapper) *Repository {
	return &Repository{
		cfg: cfg,
		db:  db,
		qb:  qb,
	}
}

func (r *Repository) RegisterNewUser(ctx context.Context, newUser entity.UserEntity) (*entity.UserEntity, error) {
	resp := newUser
	query := r.qb.Insert("users").Rows(newUser.GetRecord()).Returning("uid")
	sql, _, err := query.ToSQL()
	if err != nil {
		return nil, err
	}
	err = r.db.GetContext(ctx, &resp, sql)
	if err != nil {
		return nil, err
	}
	return &resp, err
}
