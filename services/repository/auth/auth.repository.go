package auth

import (
	"bistleague-be/model/config"
	"context"
	"fmt"
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

func (r *Repository) RegisterNewUser(ctx context.Context) (string, error) {
	query := r.qb.Insert("users").Prepared(true).Rows(goqu.Record{
		"mim": "mun",
	}, goqu.Record{
		"mim": "mun",
	}).Returning("id")
	sql, params, err := query.ToSQL()
	if err != nil {
		return "", err
	}
	fmt.Println(sql, params)
	//r.db.ExecContext(ctx, sql, params)
	return sql, err
}
