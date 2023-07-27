package entity

import (
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	"time"
)

type UserEntity struct {
	UID         string         `db:"uid"`
	TeamID      string         `db:"team_id"`
	Username    string         `db:"username"`
	Password    string         `db:"password"`
	Email       string         `db:"email"`
	FullName    string         `db:"full_name"`
	Institution sql.NullString `db:"institution"`
	Major       sql.NullString `db:"major"`
	LinkedInURL sql.NullString `db:"linkedin_url"`
	LineID      sql.NullString `db:"line_id"`
	InsertedAt  time.Time      `db:"inserted_at"`
	UpdatedAt   sql.NullTime   `db:"updated_at,omitempty"`
}

func (e *UserEntity) GetRecord() goqu.Record {
	return goqu.Record{
		"username":     e.Username,
		"password":     e.Password,
		"email":        e.Email,
		"full_name":    e.FullName,
		"institution":  e.Institution.String,
		"major":        e.Major.String,
		"linkedin_url": e.LinkedInURL.String,
		"line_id":      e.LineID.String,
		"updated_at":   e.UpdatedAt.Time,
	}
}
