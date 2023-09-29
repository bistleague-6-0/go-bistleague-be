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
	UID                  string         `db:"uid"`
	TeamName             sql.NullString `db:"team_name"`
	FullName             string         `db:"full_name"`
	StudentCardFilename  sql.NullString `db:"student_card_filename"`
	StudentCardURL       sql.NullString `db:"student_card_url"`
	StudentCardStatus    int8           `db:"student_card_status"`
	SelfPortraitFilename sql.NullString `db:"self_portrait_filename"`
	SelfPortraitURL      sql.NullString `db:"self_portrait_url"`
	SelfPortraitStatus   int8           `db:"self_portrait_status"`
	TwibbonFilename      sql.NullString `db:"twibbon_filename"`
	TwibbonURL           sql.NullString `db:"twibbon_url"`
	TwibbonStatus        int8           `db:"twibbon_status"`
	EnrollmentFilename   sql.NullString `db:"enrollment_filename"`
	EnrollmentURL        sql.NullString `db:"enrollment_url"`
	EnrollmentStatus     int8           `db:"enrollment_status"`
	IsProfileVerified    bool           `db:"is_profile_verified"`
}

type UserChallengeEntity struct {
	UID              string `db:"uid"`
	IgUsername       string `db:"ig_username" validate:"required"`
	IgContentURl     string `db:"ig_content_url" validate:"required"`
	TiktokUsername   string `db:"tiktok_username"`
	TiktokContentURl string `db:"tiktok_content_url"`
}
