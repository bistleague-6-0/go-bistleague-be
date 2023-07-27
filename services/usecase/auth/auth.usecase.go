package auth

import (
	"bistleague-be/model/config"
	"bistleague-be/services/repository/auth"
	"context"
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

func (u *Usecase) InsertNewUser(ctx context.Context) (string, error) {
	return u.repo.RegisterNewUser(ctx)
}
