package dto

import (
	"time"
)

type RedeemTeamCodeRequestDTO struct {
	RedeemCode string `json:"redeem_code" validate:"required"`
}

type InsertTeamDocumentRequestDTO struct {
	Type         string `json:"doc_type" validate:"required,isTeamDoc"`
	DocumentName string `json:"document_name" validate:"required"`
	Document     string `json:"document" validate:"base64,required"`
}

type CreateTeamRequestDTO struct {
	TeamName     string   `json:"team_name" validate:"required,min=8,max=20"`
	MemberEmails []string `json:"emails" validate:"required,listOfMail"`
}
type CreateTeamResponseDTO struct {
	TeamRedeemToken string `json:"team_redeem_token"`
	JWTToken        string `json:"jwt_token"`
}

type GetTeamInfoResponseDTO struct {
	TeamID         string `json:"team_id"`
	TeamName       string `json:"team_name"`
	TeamRedeemCode string `json:"team_redeem_code"`
	// Is Team Participating
	IsActive bool `json:"is_active"`
	// Is Team payment is verified
	Payment           string `json:"payment_proof"`
	PaymentURL        string `json:"payment_proof_url"`
	PaymentStatus     string `json:"payment_status"`
	PaymentStatusCode int8   `json:"payment_status_code"`

	StudentCard           string `json:"student_card"`
	StudentCardURL        string `json:"student_card_url"`
	StudentCardStatus     string `json:"student_card_status"`
	StudentCardStatusCode int8   `json:"student_card_status_code"`

	SelfPortrait           string `json:"self_portrait"`
	SelfPortraitURL        string `json:"self_portrait_url"`
	SelfPortraitStatus     string `json:"self_portrait_status"`
	SelfPortraitStatusCode int8   `json:"self_portrait_status_code"`

	Twibbon           string `json:"twibbon"`
	TwibbonURL        string `json:"twibbon_url"`
	TwibbonStatus     string `json:"twibbon_status"`
	TwibbonStatusCode int8   `json:"twibbon_status_code"`

	Enrollment           string `json:"enrollment"`
	EnrollmentURL        string `json:"enrollment_url"`
	EnrollmentStatus     string `json:"enrollment_status"`
	EnrollmentStatusCode int8   `json:"enrollment_status_code"`

	Members []GetTeamMemberInfoResponseDTO `json:"members"`
}

type GetTeamMemberInfoResponseDTO struct {
	UserID            string `json:"user_id"`
	Username          string `json:"username"`
	Fullname          string `json:"fullname"`
	IsLeader          bool   `json:"is_leader"`
	IsDocVerified     bool   `json:"is_doc_verified"`
	IsProfileVerified bool   `json:"is_profile_verified"`
}

type InputTeamDocumentResponseDTO struct {
	DocumentType string `json:"doc_type"`
	DocumentName string `json:"doc_name"`
	DocumentURL  string `json:"doc_url"`
}

type GetSubmissionResponseDTO struct {
	TeamID               string    `json:"team_id"`
	DocumentType         string    `json:"doc_type"`
	SubmissionFilename   string    `json:"submission_filename"`
	SubmissionUrl        string    `json:"submission_url"`
	SubmissionLastUpdate time.Time `json:"submission_lastupdate"`
}
