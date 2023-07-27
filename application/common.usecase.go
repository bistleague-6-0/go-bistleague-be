package application

import (
	"bistleague-be/model/config"
	"bistleague-be/services/usecase/auth"
	"bistleague-be/services/usecase/hello"
)

type CommonUsecase struct {
	HelloUC *hello.Usecase
	AuthUC  *auth.Usecase
}

func NewCommonUsecase(cfg *config.Config, commonRepo *CommonRepository) (*CommonUsecase, error) {
	helloUC := hello.New(cfg, commonRepo.helloRepo)
	authUC := auth.New(cfg, commonRepo.authRepo)
	commonUC := CommonUsecase{
		HelloUC: helloUC,
		AuthUC:  authUC,
	}
	return &commonUC, nil
}
