package application

import (
	"bistleague-be/model/config"
	"bistleague-be/services/repository/auth"
	"bistleague-be/services/repository/hello"
)

type CommonRepository struct {
	helloRepo *hello.Repository
	authRepo  *auth.Repository
}

func NewCommonRepository(cfg *config.Config, rsc *CommonResource) (*CommonRepository, error) {
	helloRepo := hello.New(cfg)
	authRepo := auth.New(cfg, rsc.Db, rsc.QBuilder)
	commonRepo := CommonRepository{
		helloRepo: helloRepo,
		authRepo:  authRepo,
	}
	return &commonRepo, nil
}
