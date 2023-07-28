package dto

type CreateTeamRequestDTO struct {
	TeamName     string   `json:"team_name" validate:"required,min=8,max=20"`
	MemberEmails []string `json:"emails" validate:"required,listOfMail"`
	PaymentProof string   `json:"payment_proof" validate:"base64"`
}
