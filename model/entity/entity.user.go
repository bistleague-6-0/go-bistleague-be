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
	Address     string         `db:"address"`
	PhoneNumber sql.NullString `db:"phone_number"`
	Institution sql.NullString `db:"institution"`
	Major       sql.NullString `db:"major"`
	EntryYear   uint32         `db:"entry_year"`
	LinkedInURL sql.NullString `db:"linkedin_url"`
	LineID      sql.NullString `db:"line_id"`
	InsertedAt  time.Time      `db:"inserted_at"`
	UpdatedAt   sql.NullTime   `db:"updated_at,omitempty"`
}

type UserDocs struct {
	UID                  string `db:"uid"`
	TeamName             string `db:"team_name"`
	FullName             string `db:"full_name"`
	StudentCardFilename  string `db:"student_card_filename"`
	StudentCardURL       string `db:"student_card_url"`
	StudentCardStatus    string `db:"student_card_status"`
	SelfPortraitFilename string `db:"self_portrait_filename"`
	SelfPortraitURL      string `db:"self_portrait_url"`
	SelfPortraitStatus   string `db:"self_portrait_status"`
	TwibbonFilename      string `db:"twibbon_filename"`
	TwibbonURL           string `db:"twibbon_url"`
	TwibbonStatus        string `db:"twibbon_status"`
	EnrollmentFilename   string `db:"enrollment_filename"`
	EnrollmentURL        string `db:"enrollment_url"`
	EnrollmentStatus     string `db:"enrollment_status"`
	IsProfileVerified    bool   `db:"is_profile_verified"`
}
