package dto

type CreateUserRequestDTO struct {
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
	RePassword  string `json:"re_password" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	FullName    string `json:"full_name" validate:"required"`
	Institution string `json:"institution"`
	Major       string `json:"major"`
	LinkedInURL string `json:"linkedin_url"`
	LineID      string `json:"line_id"`
}

type AuthUserResponseDTO struct {
	Info  UserResponseDTO `json:"info"`
	Token string          `json:"token"`
}

type UserResponseDTO struct {
	UID         string `json:"uid"`
	TeamID      string `json:"team_id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	Institution string `json:"institution"`
	Major       string `json:"major"`
	LinkedInURL string `json:"linkedin_url"`
	LineID      string `json:"line_id"`
}
