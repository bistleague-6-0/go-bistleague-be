package application

import (
	"bistleague-be/model/config"
	"bistleague-be/services/usecase/admin"
	"bistleague-be/services/usecase/auth"
	"bistleague-be/services/usecase/challenge"
	"bistleague-be/services/usecase/hello"
	"bistleague-be/services/usecase/profile"
	"bistleague-be/services/usecase/team"
)

type CommonUsecase struct {
	HelloUC     *hello.Usecase
	AuthUC      *auth.Usecase
	TeamUC      *team.Usecase
	ProfileUC   *profile.Usecase
	AdminUC     *admin.Usecase
	ChallengeUC *challenge.Usecase
}

func NewCommonUsecase(cfg *config.Config, commonRepo *CommonRepository) (*CommonUsecase, error) {
	helloUC := hello.New(cfg, commonRepo.helloRepo)
	authUC := auth.New(cfg, commonRepo.authRepo, commonRepo.cacheRepo, commonRepo.emailRepo)
	teamUC := team.New(cfg, commonRepo.teamRepo, commonRepo.storageRepo, commonRepo.profileRepo)
	profileUC := profile.New(cfg, commonRepo.profileRepo)
	adminUC := admin.New(cfg, commonRepo.adminRepo, commonRepo.profileRepo, commonRepo.teamRepo, commonRepo.challengeRepo, commonRepo.emailRepo)
	challengeUC := challenge.New(cfg, commonRepo.challengeRepo)
	commonUC := CommonUsecase{
		HelloUC:     helloUC,
		AuthUC:      authUC,
		TeamUC:      teamUC,
		ProfileUC:   profileUC,
		AdminUC:     adminUC,
		ChallengeUC: challengeUC,
	}
	return &commonUC, nil
}
