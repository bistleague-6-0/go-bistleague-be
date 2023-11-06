package auth

import (
	"bistleague-be/model/config"
	"bistleague-be/model/dto"
	"bistleague-be/model/entity"
	"bistleague-be/services/repository/auth"
	"bistleague-be/services/repository/cache"
	email "bistleague-be/services/repository/email"
	"bistleague-be/services/utils"
	"bistleague-be/services/utils/encryptor"
	"bistleague-be/services/utils/randomizer"
	"context"
	"errors"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Usecase struct {
	cfg       *config.Config
	repo      *auth.Repository
	cacheRepo *cache.Repository
	emailRepo *email.Repository
}

func New(cfg *config.Config, repo *auth.Repository, cacheRepo *cache.Repository, emailRepo *email.Repository) *Usecase {
	return &Usecase{
		cfg:       cfg,
		repo:      repo,
		cacheRepo: cacheRepo,
		emailRepo: emailRepo,
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
	user, err := u.repo.GetUserInformationByEmail(ctx, email)
	if err != nil {
		return errors.New("internal server error")
	}
	token := randomizer.RandStringBytes(25)
	isUsed := u.cacheRepo.Check(ctx, token)
	if isUsed {
		return errors.New("internal server error")
	}
	u.cacheRepo.Set(ctx, token, user.UID, time.Hour*5)
	var data struct {
		NamaLengkap string
		Token       string
	}
	data.NamaLengkap = user.FullName
	data.Token = token
	err = u.emailRepo.SendEmailHTML([]string{email}, "validate request", forgetPasswordEmailTemplate, data)
	return nil
}

func (u *Usecase) ValidateForgetPasswordTokenUsecase(ctx context.Context, token string) error {
	listOfEmail := []string{
		"fdika24@outlook.com",
		"18220015@std.stei.itb.ac.id",
	}
	var data struct {
		NamaLengkap string
		Email       string
		Password    string
	}
	resp, err := u.cacheRepo.Get(ctx, token)
	if err != nil {
		log.Error(err)
		return err
	}
	data.NamaLengkap = "Testing Nama"
	data.Email = resp
	data.Password = "jemb000tts"

	err = u.emailRepo.SendEmailHTML(listOfEmail, "temporary password", forgetPasswordValidateTokenEmailTemplate, data)
	if err != nil {
		log.Error(err)
	}
	return err
}
