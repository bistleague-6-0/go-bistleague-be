package dto

type UserResponseDTO struct {
	UID         string `json:"uid"`
	TeamID      string `json:"team_id"`
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	Username    string `json:"username"`
	Institution string `json:"institution"`
	Major       string `json:"major"`
	EntryYear   uint32 `json:"entry_year"`
	LinkedInURL string `json:"linkedin_url"`
	LineID      string `json:"line_id"`
}
