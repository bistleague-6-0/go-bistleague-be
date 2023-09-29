package challenge

import (
	"bistleague-be/model/dto"
	"bistleague-be/model/entity"
	"bistleague-be/services/repository/challenge"
	"context"
)

type Usecase struct {
	repo *challenge.Repository
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
