package entity

import (
	"database/sql"
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

	PaymentFilename  string `db:"payment_filename"`
	PaymentURL       string `db:"payment_url"`
	PaymentStatus    int8   `db:"payment_status"`
	PaymentRejection string `db:"payment_rejection"`

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

	StudentCard          string `db:"student_card_filename"`
	StudentCardStatus    int8   `db:"student_card_status"`
	StudentCardURL       string `db:"student_card_url"`
	StudentCardRejection string `db:"student_card_rejection"`

	SelfPortrait          string `db:"self_portrait_filename"`
	SelfPortraitStatus    int    `db:"self_portrait_status"`
	SelfPortraitURL       string `db:"self_portrait_url"`
	SelfPortraitRejection string `db:"self_portrait_rejection"`

	Twibbon          string `db:"twibbon_filename"`
	TwibbonStatus    int8   `db:"twibbon_status"`
	TwibbonURL       string `db:"twibbon_url"`
	TwibbonRejection string `db:"twibbon_rejection"`

	Enrollment          string `db:"enrollment_filename"`
	EnrollmentStatus    int8   `db:"enrollment_status"`
	EnrollmentURL       string `db:"enrollment_url"`
	EnrollmentRejection string `db:"enrollment_rejection"`

	Submission1Url        *string    `db:"submission_1_url"`
	Submission2Url        *string    `db:"submission_2_url"`

	RedeemCode string `db:"code"`
}

type TeamSubmission struct {
	TeamID                string     `db:"team_id"`
	TeamName              string     `db:"team_name"`
	Submission1Filename   sql.NullString    `db:"submission_1_filename"`
	Submission1Url        sql.NullString    `db:"submission_1_url"`
	Submission1LastUpdate sql.NullTime `db:"submission_1_lastupdate"`
	Submission2Filename   sql.NullString    `db:"submission_2_filename"`
	Submission2Url        sql.NullString    `db:"submission_2_url"`
	Submission2LastUpdate sql.NullTime `db:"submission_2_lastupdate"`
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

type TeamVerification struct {
	UserID             string         `db:"uid"`
	TeamID             sql.NullString `db:"team_id"`
	TeamName           string         `db:"team_name"`
	FullName           string         `db:"full_name"`
	TeamLeaderID       string         `db:"team_leader_id"`
	Email              string         `db:"email"`
	Phone              sql.NullString `db:"phone_number"`
	PaymentStatus      int8           `db:"payment_status"`
	StudentCardStatus  int8           `db:"student_card_status"`
	SelfPortraitStatus int            `db:"self_portrait_status"`
	TwibbonStatus      int8           `db:"twibbon_status"`
	EnrollmentStatus   int8           `db:"enrollment_status"`
}
