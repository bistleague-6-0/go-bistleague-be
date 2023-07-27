package dto

type SignUpUserRequestDTO struct {
	Username    string `json:"username" validate:"required,min=5,max=20"`
	Password    string `json:"password" validate:"required,min=8,max=20"`
	RePassword  string `json:"re_password" validate:"required,min=8,max=20"`
	Email       string `json:"email" validate:"required,email"`
	FullName    string `json:"full_name" validate:"required"`
	Institution string `json:"institution"`
	Major       string `json:"major"`
	LinkedInURL string `json:"linkedin_url"`
	LineID      string `json:"line_id"`
}

type SignInUserRequestDTO struct {
	Username   string `json:"username" validate:"required,min=5,max=20"`
	Password   string `json:"password" validate:"required,min=8,max=20"`
	RePassword string `json:"re_password" validate:"required,min=8,max=20"`
}

type AuthUserResponseDTO struct {
	User  UserResponseDTO `json:"user"`
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
