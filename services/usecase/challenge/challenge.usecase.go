package challenge

import (
	"bistleague-be/model/config"
	"bistleague-be/model/dto"
	"bistleague-be/model/entity"
	"bistleague-be/services/repository/challenge"
	"context"
)

type Usecase struct {
	cfg  *config.Config
	repo *challenge.Repository
}

func New(cfg *config.Config, repo *challenge.Repository) *Usecase {
	return &Usecase{
		cfg:  cfg,
		repo: repo,
	}
}

func (u *Usecase) AddNewChallengeUsecase(ctx context.Context, req dto.InsertChallengeRequestDTO, userID string) (*dto.ChallengeResponseDTO, error) {
	uChallenge := entity.UserChallengeEntity{
		UID:              userID,
		IgUsername:       req.IgUsername,
		IgContentURl:     req.IgContentURl,
		TiktokUsername:   req.TiktokUsername,
		TiktokContentURl: req.TiktokContentURl,
	}
	err := u.repo.InsertUserChallenge(ctx, uChallenge)
	if err != nil {
		return nil, err
	}
	return &dto.ChallengeResponseDTO{
		UID: uChallenge.UID,
		InsertChallengeRequestDTO: dto.InsertChallengeRequestDTO{
			IgUsername:       uChallenge.IgUsername,
			IgContentURl:     uChallenge.IgContentURl,
			TiktokUsername:   uChallenge.TiktokUsername,
			TiktokContentURl: uChallenge.TiktokContentURl,
		},
	}, nil
}

func (u *Usecase) UpdateChallengeUsecase(ctx context.Context, req dto.UpdateChallengeRequestDTO, userID string) (*dto.ChallengeResponseDTO, error) {
	uChallenge := entity.UserChallengeEntity{
		UID:              userID,
		IgUsername:       req.IgUsername,
		IgContentURl:     req.IgContentURl,
		TiktokUsername:   req.TiktokUsername,
		TiktokContentURl: req.TiktokContentURl,
	}
	err := u.repo.UpdateUserChallenge(ctx, uChallenge)
	if err != nil {
		return nil, err
	}
	return &dto.ChallengeResponseDTO{
		UID: uChallenge.UID,
		InsertChallengeRequestDTO: dto.InsertChallengeRequestDTO{
			IgUsername:       uChallenge.IgUsername,
			IgContentURl:     uChallenge.IgContentURl,
			TiktokUsername:   uChallenge.TiktokUsername,
			TiktokContentURl: uChallenge.TiktokContentURl,
		},
	}, nil
}

func (u *Usecase) GetChallengeUsecase(ctx context.Context, userID string) (*dto.ChallengeResponseDTO, error) {
	resp, err := u.repo.GetUserChallenge(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &dto.ChallengeResponseDTO{
		UID: resp.UID,
		InsertChallengeRequestDTO: dto.InsertChallengeRequestDTO{
			IgUsername:       resp.IgUsername,
			IgContentURl:     resp.IgContentURl,
			TiktokUsername:   resp.TiktokUsername,
			TiktokContentURl: resp.TiktokContentURl,
		},
	}, nil
}
