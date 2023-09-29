package challenge

import (
	"bistleague-be/model/config"
	"bistleague-be/model/entity"
	"context"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	cfg                 *config.Config
	db                  *sqlx.DB
	insertChallengeStmt *sqlx.Stmt
}

func New(cfg *config.Config, db *sqlx.DB) (*Repository, error) {
	repo := Repository{
		cfg: cfg,
		db:  db,
	}
	if err := repo.prepare(); err != nil {
		return nil, err
	}
	return &repo, nil
}

func (r *Repository) prepare() error {
	ctx := context.Background()
	iChallengeStmt, err := r.db.PreparexContext(ctx, QueryInsertChallenge)
	if err != nil {
		return err
	}
	r.insertChallengeStmt = iChallengeStmt
	return nil
}

func (r *Repository) InsertUserChallenge(ctx context.Context, req entity.UserChallengeEntity) error {
	_, err := r.insertChallengeStmt.ExecContext(ctx, req.UID, req.IgUsername, req.IgContentURl, req.TiktokUsername, req.TiktokContentURl)
	return err
}
