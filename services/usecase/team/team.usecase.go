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
		//BuktiPembayaranURL: req.PaymentProof,
		TeamMemberMails: req.MemberEmails,
	}
	err := u.repo.CreateTeam(ctx, team)
	if err != nil {
		log.Println(err)
	}
	return err
}
