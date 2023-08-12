package profile

import (
	"bistleague-be/model/config"
	"bistleague-be/model/dto"
	"bistleague-be/model/entity"
	"bistleague-be/services/repository/profile"
	"context"
	"database/sql"
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
		Age:         resp.Age,
		Address:     resp.Address,
		PhoneNumber: resp.PhoneNumber.String,
		Institution: resp.Institution.String,
		Major:       resp.Major.String,
		EntryYear:   resp.EntryYear,
		LinkedInURL: resp.LinkedInURL.String,
		LineID:      resp.LineID.String,
	}, nil
}

func (u *Usecase) UpdateUserProfile(ctx context.Context, req dto.UpdateUserProfileRequestDTO, userID string) error {
	ety := entity.UserEntity{
		UID:      userID,
		Email:    req.Email,
		FullName: req.FullName,
		Age:      req.Age,
		Address:  req.Address,
		PhoneNumber: sql.NullString{
			String: req.PhoneNumber,
		},
		Institution: sql.NullString{
			String: req.Institution,
		},
		Major: sql.NullString{
			String: req.Major,
		},
		EntryYear: req.EntryYear,
		LinkedInURL: sql.NullString{
			String: req.LinkedInURL,
		},
		LineID: sql.NullString{
			String: req.LineID,
		},
	}
	return u.repo.UpdateUserProfile(ctx, ety)
}
