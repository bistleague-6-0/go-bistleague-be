package application

import (
	"bistleague-be/model/config"
	"bistleague-be/services/repository/auth"
	"bistleague-be/services/repository/hello"
	"bistleague-be/services/repository/profile"
	"bistleague-be/services/repository/team"
)

type CommonRepository struct {
	helloRepo   *hello.Repository
	authRepo    *auth.Repository
	teamRepo    *team.Repository
	profileRepo *profile.Repository
}

func NewCommonRepository(cfg *config.Config, rsc *CommonResource) (*CommonRepository, error) {
	helloRepo := hello.New(cfg)
	authRepo := auth.New(cfg, rsc.Db, rsc.QBuilder)
	teamRepo := team.New(cfg, rsc.Db, rsc.QBuilder)
	profileRepo := profile.New(cfg, rsc.Db)
	commonRepo := CommonRepository{
		helloRepo:   helloRepo,
		authRepo:    authRepo,
		teamRepo:    teamRepo,
		profileRepo: profileRepo,
	}
	return &commonRepo, nil
}
