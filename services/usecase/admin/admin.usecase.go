package admin

import (
	"bistleague-be/model/config"
	"bistleague-be/model/dto"
	"bistleague-be/model/entity"
	adminRepo "bistleague-be/services/repository/admin"
	profileRepo "bistleague-be/services/repository/profile"
	teamRepo "bistleague-be/services/repository/team"
	"bistleague-be/services/utils"
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Usecase struct {
	cfg         *config.Config
	repo        *adminRepo.Repository
	teamRepo    *teamRepo.Repository
	profileRepo *profileRepo.Repository
}

func New(cfg *config.Config, repo *adminRepo.Repository) *Usecase {
	return &Usecase{
		cfg:  cfg,
		repo: repo,
	}
}

func (u *Usecase) InsertNewAdmin(ctx context.Context, req dto.RegisterAdminRequestDTO) (*dto.AuthAdminResponseDTO, error) {
	newpw, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	admin := entity.AdminEntity{
		Password: string(newpw),
		Email:    req.Email,
		FullName: req.FullName,
		Username: req.Username,
	}
	resp, err := u.repo.RegisterNewAdmin(ctx, admin)
	if err != nil {
		return nil, err
	}
	claims := entity.CustomClaim{
		TeamID: "",
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
	return &dto.AuthAdminResponseDTO{
		Admin: dto.AuthAdminInfoResponse{
			AdminID:  resp.UID,
			Username: resp.Username,
		},
		Token: token,
	}, nil
}

func (u *Usecase) SignInAdmin(ctx context.Context, req dto.SignInAdminRequestDTO) (*dto.AuthAdminResponseDTO, error) {
	admin, err := u.repo.LoginAdmin(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	claims := entity.CustomClaim{
		TeamID: "",
		UserID: admin.UID,
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
	return &dto.AuthAdminResponseDTO{
		Admin: dto.AuthAdminInfoResponse{
			AdminID:  admin.UID,
			Username: admin.Username,
		},
		Token: token,
	}, nil
}

func (u *Usecase) GetTeamPayment(ctx context.Context, page int, pageSize int) (*dto.PaginationDTOWrapper, error) {
	resp, err := u.teamRepo.GetPayments(ctx, page, pageSize)
	if err != nil {
		return nil, err
	}

	totalTeam, err := u.teamRepo.GetTeamCount(ctx)
	if err != nil {
		return nil, err
	}

	var dtoResp dto.PaginationDTOWrapper

	totalPage := (totalTeam + pageSize - 1) / pageSize

	dtoResp = dto.PaginationDTOWrapper{
		PageSize:  pageSize,
		Page:      page,
		TotalPage: totalPage,
		Data:      resp,
	}

	return &dtoResp, nil
}

func (u *Usecase) GetUserList(ctx context.Context, page int, pageSize int) (*dto.PaginationDTOWrapper, error) {
	resp, err := u.profileRepo.GetUserList(ctx, page, pageSize)
	if err != nil {
		return nil, err
	}

	totalUser, err := u.profileRepo.GetUserCount(ctx)
	if err != nil {
		return nil, err
	}

	var dtoResp dto.PaginationDTOWrapper

	totalPage := (totalUser + pageSize - 1) / pageSize

	dtoResp = dto.PaginationDTOWrapper{
		PageSize:  pageSize,
		Page:      page,
		TotalPage: totalPage,
		Data:      resp,
	}

	return &dtoResp, nil
}
