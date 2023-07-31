package profile

import (
	"bistleague-be/model/config"
	"bistleague-be/model/entity"
	"context"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	cfg *config.Config
	db  *sqlx.DB
}

func New(cfg *config.Config, db *sqlx.DB) *Repository {
	return &Repository{
		cfg: cfg,
		db:  db,
	}
}

func (r *Repository) GetUserProfile(ctx context.Context, userID string) (*entity.UserEntity, error) {
	q := `SELECT 
    uid, team_id, email, full_name, username, institution, major, entry_year, linkedin_url, line_id 
	FROM users WHERE uid = $1 LIMIT 1`
	resp := entity.UserEntity{}
	err := r.db.GetContext(ctx, &resp, q, userID)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
