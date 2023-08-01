package dto

type UpdateUserProfileRequestDTO struct {
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	Age         uint64 `json:"age"`
	PhoneNumber string `json:"phone_number"`
	Institution string `json:"institution"`
	Major       string `json:"major"`
	EntryYear   uint32 `json:"entry_year"`
	LinkedInURL string `json:"linkedin_url"`
	LineID      string `json:"line_id"`
}

type UserProfileResponseDTO struct {
	UID         string `json:"uid"`
	TeamID      string `json:"team_id"`
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	Age         uint64 `json:"age"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
	Institution string `json:"institution"`
	Major       string `json:"major"`
	EntryYear   uint32 `json:"entry_year"`
	LinkedInURL string `json:"linkedin_url"`
	LineID      string `json:"line_id"`
}
