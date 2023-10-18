package profile

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
	fmt.Println(query)
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, query)
	return err
}

func (r *Repository) GetUserCount(ctx context.Context) (int, error) {
	q := `SELECT COUNT(*) FROM users`
	var count int
	err := r.db.GetContext(ctx, &count, q)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *Repository) GetUserList(ctx context.Context, page int, pageSize int) ([]entity.UserDocs, error) {
	q := `SELECT 
			u.uid, t.team_name, u.full_name, 
			ud.student_card_filename, ud.student_card_url, ud.student_card_status, 
			ud.self_portrait_filename, ud.self_portrait_url, ud.self_portrait_status,
			ud.twibbon_filename, ud.twibbon_url, ud.twibbon_status,
			ud.enrollment_filename, ud.enrollment_url, ud.enrollment_status,
			u.is_profile_verified
		FROM users u
		LEFT JOIN users_docs ud
		ON u.uid = ud.uid
		LEFT JOIN teams t
		ON u.team_id = t.team_id
		ORDER BY u.full_name
		LIMIT $1 OFFSET $2
	`
	resp := []entity.UserDocs{}
	offset := (page - 1) * pageSize
	err := r.db.SelectContext(ctx, &resp, q, pageSize, offset)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *Repository) UpdateUserDocumentStatus(ctx context.Context, userID string, doctype string, status int, rejection string) (string, error) {
	var teamID string
	err := r.db.GetContext(ctx, &teamID, "SELECT team_id FROM users WHERE uid = ?", userID)
	if err != nil {
		return "", err
	}

	q := r.qb.Update("users_docs").Where(goqu.C("uid").Eq(userID))
	if doctype == "student_card" {
		q = q.Set(goqu.Record{
			"student_card_status":    status,
			"student_card_rejection": rejection,
		})
	} else if doctype == "self_portrait" {
		q = q.Set(goqu.Record{
			"self_portrait_status":    status,
			"self_portrait_rejection": rejection,
		})
	} else if doctype == "twibbon" {
		q = q.Set(goqu.Record{
			"twibbon_status":    status,
			"twibbon_rejection": rejection,
		})
	} else {
		q = q.Set(goqu.Record{
			"enrollment_status":    status,
			"enrollment_rejection": rejection,
		})
	}
	query, _, err := q.ToSQL()
	if err != nil {
		return "", err
	}

	_, err = r.db.ExecContext(ctx, query)
	if err != nil {
		return "", err
	}

	return teamID, nil

}
