package application

import (
	"bistleague-be/model/config"
	"bistleague-be/services/repository/admin"
	"bistleague-be/services/repository/auth"
	"bistleague-be/services/repository/cache"
	"bistleague-be/services/repository/challenge"
	emailRepo "bistleague-be/services/repository/email"
	"bistleague-be/services/repository/hello"
	"bistleague-be/services/repository/profile"
	"bistleague-be/services/repository/storage"
	"bistleague-be/services/repository/team"
)

type CommonRepository struct {
	helloRepo     *hello.Repository
	authRepo      *auth.Repository
	teamRepo      *team.Repository
	profileRepo   *profile.Repository
	storageRepo   *storage.Repository
	adminRepo     *admin.Repository
	challengeRepo *challenge.Repository
	emailRepo     *emailRepo.Repository
	cacheRepo     *cache.Repository
}

func NewCommonRepository(cfg *config.Config, rsc *CommonResource) (*CommonRepository, error) {
	helloRepo := hello.New(cfg)
	authRepo := auth.New(cfg, rsc.Db, rsc.QBuilder)
	teamRepo := team.New(cfg, rsc.Db, rsc.QBuilder)
	profileRepo := profile.New(cfg, rsc.Db, rsc.QBuilder)
	storageRepo := storage.New(cfg, rsc.bucket)
	adminRepo := admin.New(cfg, rsc.Db, rsc.QBuilder)
	emailRepo := emailRepo.New(cfg)
	challengeRepo, err := challenge.New(cfg, rsc.Db)
	if err != nil {
		return nil, err
	}
	cacheRepo, err := cache.New(cfg, rsc.cache)
	if err != nil {
		return nil, err
	}
	commonRepo := CommonRepository{
		helloRepo:     helloRepo,
		authRepo:      authRepo,
		teamRepo:      teamRepo,
		profileRepo:   profileRepo,
		storageRepo:   storageRepo,
		adminRepo:     adminRepo,
		challengeRepo: challengeRepo,
		emailRepo:     emailRepo,
		cacheRepo:     cacheRepo,
	}
	return &commonRepo, nil
}
