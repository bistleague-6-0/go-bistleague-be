package dto

type InsertChallengeRequestDTO struct {
	IgUsername       string `json:"ig_username" validate:"required"`
	IgContentURl     string `json:"ig_content_url" validate:"required,url"`
	TiktokUsername   string `json:"tiktok_username"`
	TiktokContentURl string `json:"tiktok_content_url"`
}

type UpdateChallengeRequestDTO struct {
	UID string `json:"user_id"`
	InsertChallengeRequestDTO
}

type ChallengeResponseDTO struct {
	UID string `json:"uid"`
	InsertChallengeRequestDTO
}
