package dto

type RedeemTeamCodeRequestDTO struct {
	RedeemCode string `json:"redeem_code" validate:"required"`
}

type CreateTeamRequestDTO struct {
	TeamName     string   `json:"team_name" validate:"required,min=8,max=20"`
	MemberEmails []string `json:"emails" validate:"required,listOfMail"`
	// PaymentProof string   `json:"payment_proof" validate:"base64"`
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
	VerificationStatus     string                         `json:"verification_status"`
	VerificationStatusCode int8                           `json:"verification_status_code"`
	Members                []GetTeamMemberInfoResponseDTO `json:"members"`
}

type GetTeamMemberInfoResponseDTO struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	IsLeader bool   `json:"is_leader"`
}
