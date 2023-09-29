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
	updateChallengeStmt *sqlx.Stmt
	getChallengeStmt    *sqlx.Stmt
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
	iChallengeStmt, err := r.db.PreparexContext(ctx, queryInsertChallenge)
	r.insertChallengeStmt = iChallengeStmt
	if err != nil {
		return err
	}
	uChallengeStmt, err := r.db.PreparexContext(ctx, queryUpdateChallenge)
	if err != nil {
		return err
	}
	r.updateChallengeStmt = uChallengeStmt
	gChallengeStmt, err := r.db.PreparexContext(ctx, queryGetChallenge)
	if err != nil {
		return err
	}
	r.getChallengeStmt = gChallengeStmt
	return nil
}

func (r *Repository) InsertUserChallenge(ctx context.Context, req entity.UserChallengeEntity) error {
	_, err := r.insertChallengeStmt.ExecContext(ctx, req.UID, req.IgUsername, req.IgContentURl, req.TiktokUsername, req.TiktokContentURl)
	return err
}

func (r *Repository) UpdateUserChallenge(ctx context.Context, req entity.UserChallengeEntity) error {
	_, err := r.updateChallengeStmt.ExecContext(ctx, req.IgUsername, req.IgContentURl, req.TiktokUsername, req.TiktokContentURl, req.UID)
	return err
}

func (r *Repository) GetUserChallenge(ctx context.Context, userID string) (*entity.UserChallengeEntity, error) {
	resp := entity.UserChallengeEntity{UID: userID}
	err := r.getChallengeStmt.GetContext(ctx, &resp, userID)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
