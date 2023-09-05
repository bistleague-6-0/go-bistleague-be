package dto

type UpdateUserProfileRequestDTO struct {
	Email       string `json:"email" validate:"required,email"`
	FullName    string `json:"full_name" validate:"required"`
	Age         uint64 `json:"age" validate:"required"`
	Address     string `json:"address" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Institution string `json:"institution" validate:"required"`
	Major       string `json:"major" validate:"required"`
	EntryYear   uint32 `json:"entry_year" validate:"required"`
	LinkedInURL string `json:"linkedin_url" validate:"required,url"`
	LineID      string `json:"line_id" validate:"required"`
}

type UserProfileResponseDTO struct {
	UID         string `json:"uid"`
	TeamID      string `json:"team_id"`
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	Age         uint64 `json:"age"`
	Username    string `json:"username"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Institution string `json:"institution"`
	Major       string `json:"major"`
	EntryYear   uint32 `json:"entry_year"`
	LinkedInURL string `json:"linkedin_url"`
	LineID      string `json:"line_id"`
}

type UserDocsResponseDTO struct {
	UID                  string `json:"uid"`
	TeamName             string `json:"team_name"`
	FullName             string `json:"full_name"`
	StudentCardFilename  string `json:"student_card_filename"`
	StudentCardURL       string `json:"student_card_url"`
	StudentCardStatus    string `json:"student_card_status"`
	SelfPortraitFilename string `json:"self_portrait_filename"`
	SelfPortraitURL      string `json:"self_portrait_url"`
	SelfPortraitStatus   string `json:"self_portrait_status"`
	TwibbonFilename      string `json:"twibbon_filename"`
	TwibbonURL           string `json:"twibbon_url"`
	TwibbonStatus        string `json:"twibbon_status"`
	EnrollmentFilename   string `json:"enrollment_filename"`
	EnrollmentURL        string `json:"enrollment_url"`
	EnrollmentStatus     string `json:"enrollment_status"`
	IsProfileVerified    bool   `db:"is_profile_verified"`
}
