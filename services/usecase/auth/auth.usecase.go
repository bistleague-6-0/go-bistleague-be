package auth

import (
	"bistleague-be/model/config"
	"bistleague-be/model/dto"
	"bistleague-be/model/entity"
	"bistleague-be/services/repository/auth"
	"bistleague-be/services/utils"
	"context"
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
		Password: string(newpw),
		Email:    req.Email,
		FullName: req.FullName,
		Username: req.Username,
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
		User: dto.AuthUserInfoResponse{
			UserID:   resp.UID,
			Username: resp.Username,
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
		User: dto.AuthUserInfoResponse{
			UserID:   user.UID,
			Username: user.Username,
		},
		Token: token,
	}, nil
}
