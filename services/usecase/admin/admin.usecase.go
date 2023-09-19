package admin

import (
	"bistleague-be/model/config"
	"bistleague-be/model/dto"
	"bistleague-be/model/entity"
	adminRepo "bistleague-be/services/repository/admin"
	"bistleague-be/services/repository/profile"
	"bistleague-be/services/repository/team"
	teamRepo "bistleague-be/services/repository/team"
	"bistleague-be/services/utils"
	"context"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Usecase struct {
	cfg         *config.Config
	repo        *adminRepo.Repository
	profileRepo *profile.Repository
	teamRepo    *teamRepo.Repository
}

func New(cfg *config.Config, repo *adminRepo.Repository, profileRepo *profile.Repository, teamRepo *team.Repository) *Usecase {
	return &Usecase{
		cfg:         cfg,
		repo:        repo,
		profileRepo: profileRepo,
		teamRepo:    teamRepo,
	}
}

func (u *Usecase) InsertNewAdmin(ctx context.Context, req dto.RegisterAdminRequestDTO) (*dto.AuthAdminResponseDTO, error) {
	newpw, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	admin := entity.AdminEntity{
		Password: string(newpw),
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
	token, err := utils.CreateJWTToken(u.cfg.Secret.AdminJWT, claims)
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
	token, err := utils.CreateJWTToken(u.cfg.Secret.AdminJWT, claims)
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
	data := []dto.GetTeamPaymentResponseDTO{}
	for _, payment := range resp {
		member_email_str := strings.Trim(payment.TeamMemberMails, "{}")
		member_email := strings.Split(member_email_str, ",")
		data = append(data, dto.GetTeamPaymentResponseDTO{
			TeamID:          payment.TeamID,
			TeamName:        payment.TeamName,
			TeamMemberMails: member_email,
			PaymentFilename: payment.PaymentFilename,
			PaymentURL:      payment.PaymentURL,
			PaymentStatus:   entity.VerificationStatusMap[payment.PaymentStatus],
			Code:            payment.Code,
		})
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
		Data:      data,
	}

	return &dtoResp, nil
}

func (u *Usecase) GetUserList(ctx context.Context, page int, pageSize int) (*dto.PaginationDTOWrapper, error) {
	resp, err := u.profileRepo.GetUserList(ctx, page, pageSize)
	if err != nil {
		return nil, err
	}
	data := []dto.UserDocsResponseDTO{}
	for _, user := range resp {
		data = append(data, dto.UserDocsResponseDTO{
			UID:                  user.UID,
			TeamName:             user.TeamName.String,
			FullName:             user.FullName,
			StudentCardFilename:  user.StudentCardFilename.String,
			StudentCardURL:       user.StudentCardURL.String,
			StudentCardStatus:    entity.VerificationStatusMap[user.StudentCardStatus],
			SelfPortraitFilename: user.SelfPortraitFilename.String,
			SelfPortraitURL:      user.SelfPortraitURL.String,
			SelfPortraitStatus:   entity.VerificationStatusMap[user.SelfPortraitStatus],
			TwibbonFilename:      user.TwibbonFilename.String,
			TwibbonURL:           user.TwibbonURL.String,
			TwibbonStatus:        entity.VerificationStatusMap[user.TwibbonStatus],
			EnrollmentFilename:   user.EnrollmentFilename.String,
			EnrollmentURL:        user.EnrollmentURL.String,
			EnrollmentStatus:     entity.VerificationStatusMap[user.EnrollmentStatus],
			IsProfileVerified:    user.IsProfileVerified,
		})
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
		Data:      data,
	}

	return &dtoResp, nil
}

func (u *Usecase) UpdateTeamPaymentStatus(ctx context.Context, teamID string, status int, rejection string) error {
	return u.teamRepo.UpdatePaymentStatus(ctx, teamID, status, rejection)
}

func (u *Usecase) UpdateUserDocumentStatus(ctx context.Context, userID string, doctype string, status int, rejection string) error {
	return u.profileRepo.UpdateUserDocumentStatus(ctx, userID, doctype, status, rejection)
}
