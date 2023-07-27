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
	"golang.org/x/crypto/bcrypt"
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

func (u *Usecase) InsertNewUser(ctx context.Context, req dto.SignUpUserRequestDTO) (*dto.AuthUserResponseDTO, error) {
	newpw, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := entity.UserEntity{
		Username: req.Username,
		Password: string(newpw),
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
		TeamID: resp.TeamID.String,
		UserID: resp.UID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "rest",
			Subject:   "",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 5)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token, err := utils.CreateJWTToken(u.cfg.Secret.JWTSecret, claims)
	if err != nil {
		return nil, err
	}
	return &dto.AuthUserResponseDTO{
		User: dto.UserResponseDTO{
			UID:         resp.UID,
			TeamID:      resp.TeamID.String,
			Username:    resp.Username,
			Email:       resp.Email,
			FullName:    resp.FullName,
			Institution: resp.Institution.String,
			Major:       resp.Major.String,
			EntryYear:   resp.EntryYear,
			LinkedInURL: resp.LinkedInURL.String,
			LineID:      resp.LineID.String,
		},
		Token: token,
	}, nil
}

func (u *Usecase) SignInUser(ctx context.Context, req dto.SignInUserRequestDTO) (*dto.AuthUserResponseDTO, error) {
	user, err := u.repo.LoginUser(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	claims := entity.CustomClaim{
		TeamID: user.TeamID.String,
		UserID: user.UID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "rest",
			Subject:   "",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 5)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token, err := utils.CreateJWTToken(u.cfg.Secret.JWTSecret, claims)
	if err != nil {
		return nil, err
	}
	return &dto.AuthUserResponseDTO{
		User: dto.UserResponseDTO{
			UID:         user.UID,
			TeamID:      user.TeamID.String,
			Username:    user.Username,
			Email:       user.Email,
			FullName:    user.FullName,
			Institution: user.Institution.String,
			Major:       user.Major.String,
			EntryYear:   user.EntryYear,
			LinkedInURL: user.LinkedInURL.String,
			LineID:      user.LineID.String,
		},
		Token: token,
	}, nil
}
