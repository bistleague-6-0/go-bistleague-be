package profile

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

func (r *Repository) GetUserProfile(ctx context.Context, userID string) (*entity.UserEntity, error) {
	q := `SELECT 
    uid, team_id, email, full_name, user_age,username, phone_number, address, institution, major, entry_year, linkedin_url, line_id 
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
    major =$7, entry_year =$8, linkedin_url =$9, line_id = $10, address = $11, is_profile_verified = $12, updated_at = now() WHERE uid = $1`
	_, err := r.db.ExecContext(ctx, q, req.UID, req.Email, req.FullName, req.Age, req.PhoneNumber.String, req.Institution.String, req.Major.String, req.EntryYear, req.LinkedInURL.String, req.LineID.String, req.Address, true)
	return err
}

func (r *Repository) UpdateUserDocument(ctx context.Context, userID string, filename string, fileurl string, doctype string) error {
	q := r.qb.Update("users_docs").Where(goqu.C("uid").Eq(userID))
	if doctype == "student_card" {
		q = q.Set(goqu.Record{
			"student_card_filename": filename,
			"student_card_url":      fileurl,
			"student_card_status":   1,
		})
	} else if doctype == "self_portrait" {
		q = q.Set(goqu.Record{
			"self_portrait_filename": filename,
			"self_portrait_url":      fileurl,
			"self_portrait_status":   1,
		})
	} else if doctype == "twibbon" {
		q = q.Set(goqu.Record{
			"twibbon_filename": filename,
			"twibbon_url":      fileurl,
			"twibbon_status":   1,
		})
	} else {
		q = q.Set(goqu.Record{
			"enrollment_filename": filename,
			"enrollment_url":      fileurl,
			"enrollment_status":   1,
		})
	}
	query, _, err := q.ToSQL()
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, query)
	return err
}
