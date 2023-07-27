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

type AuthUserResponseDTO struct {
	User  AuthUserInfoResponse `json:"user"`
	Token string               `json:"jwt_token"`
}

type AuthUserInfoResponse struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}
