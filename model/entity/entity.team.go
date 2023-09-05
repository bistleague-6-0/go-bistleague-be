package entity

import (
	"time"
)

var VerificationStatusMap = map[int8]string{
	-1: "rejected",
	0:  "no file",
	1:  "under review",
	2:  "accepted",
}

type TeamEntity struct {
	TeamID       string `db:"team_id"`
	TeamName     string `db:"team_name"`
	TeamLeaderID string `db:"team_leader_id"`

	PaymentFilename string `db:"payment_filename"`
	PaymentURL      string `db:"payment_url"`
	PaymentStatus   int8   `db:"payment_status"`

	TeamMemberMails []string `db:"team_member_mails"`
	IsActive        bool     `db:"is_active"`
}

type TeamRedeemCodeEntity struct {
	TeamID string `db:"team_id"`
	Code   string `db:"code"`
	Used   int8   `db:"used"`
}

type TeamWithUserEntity struct {
	TeamEntity
	UserID            string `db:"uid"`
	Username          string `db:"username"`
	FullName          string `db:"full_name"`
	IsDocVerified     bool   `db:"is_doc_verified"`
	IsProfileVerified bool   `db:"is_profile_verified"`

	StudentCard       string `db:"student_card_filename"`
	StudentCardStatus int8   `db:"student_card_status"`
	StudentCardURL    string `db:"student_card_url"`

	SelfPortrait       string `db:"self_portrait_filename"`
	SelfPortraitStatus int8   `db:"self_portrait_status"`
	SelfPortraitURL    string `db:"self_portrait_url"`

	Twibbon       string `db:"twibbon_filename"`
	TwibbonStatus int8   `db:"twibbon_status"`
	TwibbonURL    string `db:"twibbon_url"`

	Enrollment       string `db:"enrollment_filename"`
	EnrollmentStatus int8   `db:"enrollment_status"`
	EnrollmentURL    string `db:"enrollment_url"`

	RedeemCode string `db:"code"`
}

type TeamSubmission struct {
	TeamID                string    `db:"team_id"`
	Submission1Filename   string    `db:"submission_1_filename"`
	Submission1Url        string    `db:"submission_1_url"`
	Submission1LastUpdate time.Time `db:"submission_1_lastupdate"`
	Submission2Filename   string    `db:"submission_2_filename"`
	Submission2Url        string    `db:"submission_2_url"`
	Submission2LastUpdate time.Time `db:"submission_2_lastupdate"`
}

type TeamPayment struct {
	TeamID          string `db:"team_id"`
	TeamName        string `db:"team_name"`
	TeamMemberMails string `db:"team_member_mails"`
	PaymentFilename string `db:"payment_filename"`
	PaymentURL      string `db:"payment_url"`
	PaymentStatus   int8   `db:"payment_status"`
	Code            string `db:"code"`
}
