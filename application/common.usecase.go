package application

import (
	"bistleague-be/model/config"
	"bistleague-be/services/usecase/auth"
	"bistleague-be/services/usecase/hello"
	"bistleague-be/services/usecase/team"
)

type CommonUsecase struct {
	HelloUC *hello.Usecase
	AuthUC  *auth.Usecase
	TeamUC  *team.Usecase
}

func NewCommonUsecase(cfg *config.Config, commonRepo *CommonRepository) (*CommonUsecase, error) {
	helloUC := hello.New(cfg, commonRepo.helloRepo)
	authUC := auth.New(cfg, commonRepo.authRepo)
	teamUC := team.New(cfg, commonRepo.teamRepo)
	commonUC := CommonUsecase{
		HelloUC: helloUC,
		AuthUC:  authUC,
		TeamUC:  teamUC,
	}
	return &commonUC, nil
}
