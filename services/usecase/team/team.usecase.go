package team

import (
	"bistleague-be/model/config"
	"bistleague-be/model/dto"
	"bistleague-be/model/entity"
	"bistleague-be/services/repository/team"
	"bistleague-be/services/utils"
	"context"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

type Usecase struct {
	cfg  *config.Config
	repo *team.Repository
}

func New(cfg *config.Config, repo *team.Repository) *Usecase {
	return &Usecase{
		cfg:  cfg,
		repo: repo,
	}
}

func (u *Usecase) CreateTeam(ctx context.Context, req dto.CreateTeamRequestDTO, teamLeaderID string) (string, error) {
	team := entity.TeamEntity{
		TeamName:     req.TeamName,
		TeamLeaderID: teamLeaderID,
		// MARK : PROCESS THIS FIRST
		//BuktiPembayaranURL: req.PaymentProof,
		TeamMemberMails: req.MemberEmails,
	}
	teamID, err := u.repo.CreateTeam(ctx, team)
	if err != nil {
		log.Println(err)
		return "", err
	}
	claims := entity.CustomClaim{
		TeamID: teamID,
		UserID: teamLeaderID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "rest",
			Subject:   "",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 5)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token, err := utils.CreateJWTToken(u.cfg.Secret.JWTSecret, claims)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *Usecase) GetTeamInformation(ctx context.Context, teamID string) (*dto.GetTeamInfoResponseDTO, error) {
	resp, err := u.repo.GetTeamInformation(ctx, teamID)
	if err != nil {
		return nil, err
	}
	result := dto.GetTeamInfoResponseDTO{}
	result.TeamID = teamID
	for _, team := range resp {
		result.TeamName = team.TeamName
		result.IsActive = team.IsActive
		result.VerificationStatusCode = team.VerificationStatus
		result.VerificationStatus = entity.VerificationStatusMap[team.VerificationStatus]
		result.Members = append(result.Members, dto.GetTeamMemberInfoResponseDTO{
			UserID:   team.UserID,
			Username: team.Username,
			Fullname: team.FullName,
			IsLeader: team.TeamLeaderID == team.UserID,
		})
	}
	return &result, nil
}
