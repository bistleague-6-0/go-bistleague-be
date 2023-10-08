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

func (r *Repository) GetUserChallenges(ctx context.Context, page uint64, limit uint64) ([]entity.AdminUserChallengeEntity, error) {
	offset := (page - 1) * limit
	q := `select 
		u.uid, u.username, u.email, u.full_name, 
		uc.ig_content_url, uc.ig_username, uc.tiktok_content_url, 
		uc.tiktok_username
	from users_mini_challenge uc
	left join users u on uc.uid = u.uid
	order by uc.inserted_at
	LIMIT $1 OFFSET $2
	`
	resp := []entity.AdminUserChallengeEntity{}
	err := r.db.SelectContext(ctx, &resp, q, limit, offset)
	return resp, err
}

func (r *Repository) GetUserChallengeWithUserDetail(ctx context.Context, userID string) (*entity.AdminUserChallengeEntity, error) {
	q := `select 
		uc.uid, u.username, u.email, u.full_name, 
		uc.ig_content_url, uc.ig_username, uc.tiktok_content_url, 
		uc.tiktok_username
	from users_mini_challenge uc
	left join users u on uc.uid = u.uid
	WHERE uc.uid = $1
	LIMIT 1
`
	resp := entity.AdminUserChallengeEntity{}
	err := r.db.GetContext(ctx, &resp, q, userID)
	return &resp, err
}
