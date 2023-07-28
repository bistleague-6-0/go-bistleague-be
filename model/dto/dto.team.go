package dto

type CreateTeamRequestDTO struct {
	TeamName     string   `json:"team_name" validate:"required,min=8,max=20"`
	MemberEmails []string `json:"emails" validate:"required,listOfMail"`
	PaymentProof string   `json:"payment_proof" validate:"base64"`
}

type GetTeamInfoResponseDTO struct {
	TeamID   string `json:"team_id"`
	TeamName string `json:"team_name"`
	// Is Team Participating
	IsActive bool `json:"is_active"`
	// Is Team payment is verified
	IsVerified bool                           `json:"is_verified"`
	Members    []GetTeamMemberInfoResponseDTO `json:"members"`
}

type GetTeamMemberInfoResponseDTO struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	IsLeader bool   `json:"is_leader"`
}
