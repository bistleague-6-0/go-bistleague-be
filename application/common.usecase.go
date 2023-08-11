package application

import (
	"bistleague-be/model/config"
	"bistleague-be/services/usecase/auth"
	"bistleague-be/services/usecase/hello"
	"bistleague-be/services/usecase/profile"
	"bistleague-be/services/usecase/team"
)

type CommonUsecase struct {
	HelloUC   *hello.Usecase
	AuthUC    *auth.Usecase
	TeamUC    *team.Usecase
	ProfileUC *profile.Usecase
}

func NewCommonUsecase(cfg *config.Config, commonRepo *CommonRepository) (*CommonUsecase, error) {
	helloUC := hello.New(cfg, commonRepo.helloRepo)
	authUC := auth.New(cfg, commonRepo.authRepo)
	teamUC := team.New(cfg, commonRepo.teamRepo)
	profileUC := profile.New(cfg, commonRepo.profileRepo)
	commonUC := CommonUsecase{
		HelloUC:   helloUC,
		AuthUC:    authUC,
		TeamUC:    teamUC,
		ProfileUC: profileUC,
	}
	return &commonUC, nil
}
