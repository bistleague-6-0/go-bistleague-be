package application

import (
	"bistleague-be/model/config"
	"bistleague-be/services/repository/hello"
)

type CommonRepository struct {
	helloRepo *hello.Repository
}

func NewCommonRepository(cfg *config.Config, rsc *CommonResource) (*CommonRepository, error) {
	helloRepo := hello.New(cfg, rsc.AuthClient)
	commonRepo := CommonRepository{
		helloRepo: helloRepo,
	}
	return &commonRepo, nil
}
