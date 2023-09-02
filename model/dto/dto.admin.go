package dto

type RegisterAdminRequestDTO struct {
	Email    string `json:"email" validate:"required,email"`
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
