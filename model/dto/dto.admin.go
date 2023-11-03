package dto

import (
	"time"
)


type RegisterAdminRequestDTO struct {
	Username string `json:"username" validate:"required,min=8,max=20"`
	Password string `json:"password" validate:"required,min=8,max=20"`
	FullName string `json:"full_name" validate:"required"`
}

type SignInAdminRequestDTO struct {
	Username string `json:"username" validate:"required,min=8,max=20"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}

type AuthAdminResponseDTO struct {
	Admin AuthAdminInfoResponse `json:"admin"`
	Token string                `json:"jwt_token"`
}

type AuthAdminInfoResponse struct {
	AdminID  string `json:"admin_id"`
	Username string `json:"username"`
}

type UpdateTeamPaymentStatus struct {
	Status    int    `json:"status"`
	Rejection string `json:"rejection"`
}

type UpdateUserDocumentStatus struct {
	DocumentType string `json:"doc_type"`
	Status       int    `json:"status"`
	Rejection    string `json:"rejection"`
}

type GetAllSubmissionResponseDTO struct {
	TeamID               string     `json:"team_id"`
	Submission1Filename   *string    `json:"submission1_filename"`
	Submission1Url        *string    `json:"submission1_url"`
	Submission1LastUpdate *time.Time `json:"submission1_lastupdate"`
	Submission2Filename   *string    `json:"submission2_filename"`
	Submission2Url        *string    `json:"submission2_url"`
	Submission2LastUpdate *time.Time `json:"submission2_lastupdate"`
}