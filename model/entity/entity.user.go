package entity

import (
	"database/sql"
	"time"
)

type UserEntity struct {
	UID         string         `db:"uid"`
	TeamID      sql.NullString `db:"team_id"`
	Password    string         `db:"password"`
	Email       string         `db:"email"`
	Username    string         `db:"username"`
	FullName    string         `db:"full_name"`
	Age         uint64         `db:"user_age"`
	PhoneNumber sql.NullString `db:"phone_number"`
	Institution sql.NullString `db:"institution"`
	Major       sql.NullString `db:"major"`
	EntryYear   uint32         `db:"entry_year"`
	LinkedInURL sql.NullString `db:"linkedin_url"`
	LineID      sql.NullString `db:"line_id"`
	InsertedAt  time.Time      `db:"inserted_at"`
	UpdatedAt   sql.NullTime   `db:"updated_at,omitempty"`
}
