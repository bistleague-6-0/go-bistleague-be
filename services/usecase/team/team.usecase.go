package team

import (
	"bistleague-be/model/config"
	"bistleague-be/model/dto"
	"bistleague-be/model/entity"
	"bistleague-be/services/repository/team"
	"context"
	"log"
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

func (u *Usecase) CreateTeam(ctx context.Context, req dto.CreateTeamRequestDTO, teamLeaderID string) error {
	team := entity.TeamEntity{
		TeamName:     req.TeamName,
		TeamLeaderID: teamLeaderID,
		// MARK : PROCESS THIS FIRST
		//BuktiPembayaranURL: req.PaymentProof,
		TeamMemberMails: req.MemberEmails,
	}
	err := u.repo.CreateTeam(ctx, team)
	if err != nil {
		log.Println(err)
	}
	return err
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
		result.IsVerified = team.IsVerified
		result.Members = append(result.Members, dto.GetTeamMemberInfoResponseDTO{
			UserID:   team.UserID,
			Username: team.Username,
			Fullname: team.FullName,
			IsLeader: team.TeamLeaderID == team.UserID,
		})
	}
	return &result, nil
}
