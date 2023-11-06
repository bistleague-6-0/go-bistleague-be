package auth

import (
	"bistleague-be/model/config"
	"bistleague-be/model/entity"
	"context"
	"errors"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/gofiber/fiber/v2/log"
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

func (r *Repository) GetUserInformationByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
	result := &entity.UserEntity{}
	err := r.db.GetContext(ctx, result, "SELECT uid, full_name FROM users WHERE email = $1 lIMIT 1", email)
	if err != nil {
		log.Error(err)
		return nil, errors.New("user with email not found")
	}
	return result, nil
}

func (r *Repository) UpdateUserPassword(ctx context.Context, uid string, newPass string) error {
	_, err := r.db.ExecContext(ctx, "UPDATE users SET password = $1 WHERE uid = $2", newPass, uid)
	if err != nil {
		log.Error(err)
		return errors.New("cannot update user password")
	}
	return nil
}

func (r *Repository) RegisterNewUser(ctx context.Context, newUser entity.UserEntity) (*entity.UserEntity, error) {
	resp := newUser
	tx, err := r.db.BeginTxx(ctx, nil)
	query := r.qb.Insert("users").Rows(goqu.Record{
		"email":     newUser.Email,
		"password":  newUser.Password,
		"full_name": newUser.FullName,
		"username":  newUser.Username,
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
	q2 := "INSERT INTO users_docs(uid) VALUES ($1)"
	_, err = tx.ExecContext(ctx, q2, resp.UID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	fmt.Println("WOI!")
	err = tx.Commit()
	return &resp, err
}

func (r *Repository) LoginUser(ctx context.Context, username string) (*entity.UserEntity, error) {
	resp := entity.UserEntity{}
	query := r.qb.
		Select("uid", "password", "username", "team_id").
		From("users").
		Where(goqu.C("username").Eq(username)).Limit(1)
	sql, _, err := query.ToSQL()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = r.db.GetContext(ctx, &resp, sql)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &resp, nil
}

func (r *Repository) GetUserInformation(ctx context.Context, uid string) (*entity.UserEntity, error) {
	resp := entity.UserEntity{}
	query := r.qb.
		Select("uid", "password", "username", "team_id", "email", "full_name").
		From("users").
		Where(goqu.C("uid").Eq(uid)).Limit(1)
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
