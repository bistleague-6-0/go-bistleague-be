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
    uid, team_id, email, full_name, user_age,username, phone_number, institution, major, entry_year, linkedin_url, line_id 
	FROM users WHERE uid = $1 LIMIT 1`
	resp := entity.UserEntity{}
	err := r.db.GetContext(ctx, &resp, q, userID)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (r *Repository) UpdateUserProfile(ctx context.Context, req entity.UserEntity) error {
	q := `UPDATE users 
SET email = $2, full_name = $3, user_age = $4, phone_number = $5, institution =$6, 
    major =$7, entry_year =$8, linkedin_url =$9, line_id = $10 WHERE uid = $1`
	_, err := r.db.ExecContext(ctx, q, req.UID, req.Email, req.FullName, req.Age, req.PhoneNumber.String, req.Institution.String, req.Major.String, req.EntryYear, req.LinkedInURL.String, req.LineID.String)
	return err
}