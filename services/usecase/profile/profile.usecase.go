package profile

import (
	"bistleague-be/model/config"
	"bistleague-be/model/dto"
	"bistleague-be/services/repository/profile"
	"context"
)

type Usecase struct {
	cfg  *config.Config
	repo *profile.Repository
}

func New(cfg *config.Config, repo *profile.Repository) *Usecase {
	return &Usecase{
		cfg:  cfg,
		repo: repo,
	}
}

func (u *Usecase) GetUserProfile(ctx context.Context, userID string) (*dto.UserProfileResponseDTO, error) {
	resp, err := u.repo.GetUserProfile(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &dto.UserProfileResponseDTO{
		UID:         userID,
		TeamID:      resp.TeamID.String,
		Email:       resp.Email,
		FullName:    resp.FullName,
		Username:    resp.Username,
		PhoneNumber: resp.PhoneNumber,
		Institution: resp.Institution.String,
		Major:       resp.Institution.String,
		EntryYear:   resp.EntryYear,
		LinkedInURL: resp.LinkedInURL.String,
		LineID:      resp.LineID.String,
	}, nil
}

func (u *Usecase) UpdateUserProfile(ctx context.Context) (*dto.UserProfileResponseDTO, error) {
	return nil, nil
}
