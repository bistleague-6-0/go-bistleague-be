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
	VerificationStatus     string `json:"verification_status"`
	VerificationStatusCode int8   `json:"verification_status_code"`

	StudentCard           string `json:"student_card"`
	StudentCardStatus     string `json:"student_card_status"`
	StudentCardStatusCode int8   `json:"student_card_status_code"`

	SelfPortrait           string `json:"self_portrait"`
	SelfPortraitStatus     string `json:"self_portrait_status"`
	SelfPortraitStatusCode int8   `json:"self_portrait_status_code"`

	Twibbon           string `json:"twibbon"`
	TwibbonStatus     string `json:"twibbon_status"`
	TwibbonStatusCode int8   `json:"twibbon_status_code"`

	Enrollment           string `json:"enrollment"`
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
