package application

import (
	"bistleague-be/model/config"
	"bistleague-be/services/usecase/hello"
)

type CommonUsecase struct {
	HelloUC *hello.Usecase
}

func NewCommonUsecase(cfg *config.Config, commonRepo *CommonRepository) (*CommonUsecase, error) {
	helloUC := hello.New(cfg, commonRepo.helloRepo)
	commonUC := CommonUsecase{
		HelloUC: helloUC,
	}
	return &commonUC, nil
}
