package application

import (
	"bistleague-be/model/config"
	"bistleague-be/services/repository/auth"
	"bistleague-be/services/repository/document"
	"bistleague-be/services/repository/hello"
	"bistleague-be/services/repository/profile"
	"bistleague-be/services/repository/storage"
	"bistleague-be/services/repository/team"
)

type CommonRepository struct {
	helloRepo   *hello.Repository
	authRepo    *auth.Repository
	teamRepo    *team.Repository
	docsRepo    *document.Repository
	profileRepo *profile.Repository
	storageRepo *storage.Repository
}

func NewCommonRepository(cfg *config.Config, rsc *CommonResource) (*CommonRepository, error) {
	helloRepo := hello.New(cfg)
	authRepo := auth.New(cfg, rsc.Db, rsc.QBuilder)
	teamRepo := team.New(cfg, rsc.Db, rsc.QBuilder)
	docsRepo := document.New(cfg)
	profileRepo := profile.New(cfg, rsc.Db)
	storageRepo := storage.New(cfg, rsc.Uploader)
	commonRepo := CommonRepository{
		helloRepo:   helloRepo,
		authRepo:    authRepo,
		teamRepo:    teamRepo,
		profileRepo: profileRepo,
		storageRepo: storageRepo,
		docsRepo:    docsRepo,
	}
	return &commonRepo, nil
}
