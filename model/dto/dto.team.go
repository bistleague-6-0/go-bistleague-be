package dto

type RedeemTeamCodeRequestDTO struct {
	RedeemCode string `json:"redeem_code" validate:"required"`
}

type InsertTeamDocumentRequestDTO struct {
	Type     string `json:"doc_type" validate:"required,isTeamDoc"`
	Document string `json:"document" validate:"base64"`
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
	TeamID   string `json:"team_id"`
	TeamName string `json:"team_name"`
	// Is Team Participating
	IsActive bool `json:"is_active"`
	// Is Team payment is verified
	Payment                string `json:"payment_proof"`
	PaymentURL             string `json:"payment_proof_url"`
	VerificationStatus     string `json:"verification_status"`
	VerificationStatusCode int8   `json:"verification_status_code"`

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
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	IsLeader bool   `json:"is_leader"`
}

type InputTeamDocumentResponseDTO struct {
	DocumentType string `json:"doc_type"`
	DocumentName string `json:"doc_name"`
	DocumentURL  string `json:"doc_url"`
}
