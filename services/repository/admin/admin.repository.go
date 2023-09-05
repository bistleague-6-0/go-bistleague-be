package admin

import (
	"bistleague-be/model/config"
	"bistleague-be/model/entity"
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

func (r *Repository) RegisterNewAdmin(ctx context.Context, newAdmin entity.AdminEntity) (*entity.AdminEntity, error) {
	resp := newAdmin
	tx, err := r.db.BeginTxx(ctx, nil)
	query := r.qb.Insert("admins").Rows(goqu.Record{
		"username":  newAdmin.Username,
		"password":  newAdmin.Password,
		"full_name": newAdmin.FullName,
	}).Returning("uid")
	sql, _, err := query.ToSQL()
	if err != nil {
		return nil, err
	}
	err = tx.GetContext(ctx, &resp, sql)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	fmt.Println("WOI MINT!")
	err = tx.Commit()
	return &resp, err
}

func (r *Repository) LoginAdmin(ctx context.Context, username string) (*entity.AdminEntity, error) {
	resp := entity.AdminEntity{}
	query := r.qb.
		Select("uid", "password", "username", "full_name").
		From("admins").
		Where(goqu.C("username").Eq(username)).Limit(1)
	sql, _, err := query.ToSQL()
	if err != nil {
		return nil, err
	}
	err = r.db.GetContext(ctx, &resp, sql)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
