package auth

import (
	"bistleague-be/model/config"
	"bistleague-be/model/dto"
	"bistleague-be/model/entity"
	"bistleague-be/services/repository/auth"
	"bistleague-be/services/repository/cache"
	"bistleague-be/services/utils"
	"bistleague-be/services/utils/encryptor"
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Usecase struct {
	cfg       *config.Config
	repo      *auth.Repository
	cacheRepo *cache.Repository
}

func New(cfg *config.Config, repo *auth.Repository, cacheRepo *cache.Repository) *Usecase {
	return &Usecase{
		cfg:       cfg,
		repo:      repo,
		cacheRepo: cacheRepo,
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
	token, err := utils.GenerateJWTToken(u.cfg.Secret.JWTSecret, resp.UID, resp.TeamID.String)
	if err != nil {
		return nil, err
	}
	refreshKey, err := encryptor.EncryptRefreshKey(encryptor.RefreshKey{
		Uid:       user.UID,
		TeamID:    user.TeamID.String,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 3 * 24)),
	}, u.cfg.Secret.JWTSecret)
	if err != nil {
		return nil, err
	}
	return &dto.AuthUserResponseDTO{
		User: dto.AuthUserInfoResponse{
			UserID:   resp.UID,
			Username: resp.Username,
		},
		Token:        token,
		RefreshToken: refreshKey,
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
	token, err := utils.GenerateJWTToken(u.cfg.Secret.JWTSecret, user.UID, user.TeamID.String)
	if err != nil {
		return nil, err
	}
	refreshKey, err := encryptor.EncryptRefreshKey(encryptor.RefreshKey{
		Uid:       user.UID,
		TeamID:    user.TeamID.String,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 3)),
	}, u.cfg.Secret.JWTSecret)
	if err != nil {
		return nil, err
	}
	return &dto.AuthUserResponseDTO{
		User: dto.AuthUserInfoResponse{
			UserID:   user.UID,
			Username: user.Username,
		},
		Token:        token,
		RefreshToken: refreshKey,
	}, nil
}

func (u *Usecase) RefreshToken(ctx context.Context, req dto.RefreshTokenRequestDTO) (*dto.RefreshTokenResponseDTO, error) {
	refreshVal, err := encryptor.DecryptRefreshKey(req.RefreshKey, u.cfg.Secret.JWTSecret)
	if err != nil {
		return nil, err
	}
	if refreshVal.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("expired")
	}
	resp, err := u.repo.GetUserInformation(ctx, refreshVal.Uid)
	if err != nil {
		return nil, err
	}
	if resp.UID != req.UserID {
		return nil, errors.New("user id does not match")
	}
	accessKey, err := utils.GenerateJWTToken(u.cfg.Secret.JWTSecret, resp.UID, resp.TeamID.String)
	if err != nil {
		return nil, err
	}
	result := dto.RefreshTokenResponseDTO{
		AccessToken: accessKey,
		IsUpdated:   resp.TeamID.String != "",
	}
	return &result, nil
}

func (u *Usecase) ForgetPasswordUsecase(ctx context.Context, email string) error {
	// selalu return success mau itu dia dpt emailnya atau engga.
	// soalnya, klo misalnya email ga ketemu di return email ga ketemu
	// bisa di abuse sama hacker
	// return error hanya untuk server error seperti db conn error, parse error, dll
	token := "random token dong"
	isUsed := u.cacheRepo.Check(ctx, token)
	if isUsed {
		return errors.New("udah kepake bwang")
	}
	u.cacheRepo.Set(ctx, token, email, time.Hour*5)
	resp, err := u.cacheRepo.Get(ctx, email)
	if err != nil {
		log.Error(err)
		return err
	}
	fmt.Println("cache value", resp)
	return nil
}

func (u *Usecase) ValidateForgetPasswordTokenUsecase(ctx context.Context, token string) error {
	return nil
}
