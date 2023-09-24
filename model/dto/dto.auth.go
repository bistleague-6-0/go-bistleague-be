package dto

type SignUpUserRequestDTO struct {
	Email      string `json:"email" validate:"required,email"`
	Username   string `json:"username" validate:"required,min=8,max=20"`
	Password   string `json:"password" validate:"required,min=8,max=20"`
	RePassword string `json:"re_password" validate:"required,min=8,max=20"`
	FullName   string `json:"full_name" validate:"required"`
}

type SignInUserRequestDTO struct {
	Username string `json:"username" validate:"required,min=8,max=20"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}

type RefreshTokenRequestDTO struct {
	RefreshKey string `json:"refresh_token" validate:"required"`
	UserID     string `json:"user_id" validate:"required"`
}

type RefreshTokenResponseDTO struct {
	AccessToken string `json:"jwt_token"`
	IsUpdated   bool   `json:"is_updated"`
}

type AuthUserResponseDTO struct {
	User         AuthUserInfoResponse `json:"user"`
	Token        string               `json:"jwt_token"`
	RefreshToken string               `json:"refresh_token"`
}

type AuthUserInfoResponse struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}
