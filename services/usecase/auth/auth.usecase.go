package auth

import (
	"bistleague-be/model/config"
	"bistleague-be/model/dto"
	"bistleague-be/model/entity"
	"bistleague-be/services/repository/auth"
	"bistleague-be/services/utils"
	"context"
	"database/sql"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Usecase struct {
	cfg  *config.Config
	repo *auth.Repository
}

func New(cfg *config.Config, repo *auth.Repository) *Usecase {
	return &Usecase{
		cfg:  cfg,
		repo: repo,
	}
}

func (u *Usecase) InsertNewUser(ctx context.Context, req dto.CreateUserRequestDTO) (*dto.AuthUserResponseDTO, error) {
	user := entity.UserEntity{
		Username: req.Username,
		Password: req.RePassword,
		Email:    req.Email,
		FullName: req.FullName,
		Institution: sql.NullString{
			String: req.Institution,
		},
		Major: sql.NullString{
			String: req.Major,
		},
		LinkedInURL: sql.NullString{
			String: req.LinkedInURL,
		},
		LineID: sql.NullString{
			String: req.LineID,
		},
	}
	resp, err := u.repo.RegisterNewUser(ctx, user)
	if err != nil {
		return nil, err
	}
	claims := entity.CustomClaim{
		TeamID: resp.TeamID,
		UserID: resp.UID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "rest",
			Subject:   "",
			ExpiresAt: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token, err := utils.CreateJWTToken(u.cfg.Secret.JWTSecret, claims)
	if err != nil {
		return nil, err
	}
	return &dto.AuthUserResponseDTO{
		Info: dto.UserResponseDTO{
			UID:         resp.UID,
			TeamID:      resp.TeamID,
			Username:    resp.Username,
			Email:       resp.Email,
			FullName:    resp.FullName,
			Institution: resp.Institution.String,
			Major:       resp.Major.String,
			LinkedInURL: resp.LinkedInURL.String,
			LineID:      resp.LineID.String,
		},
		Token: token,
	}, nil
}
