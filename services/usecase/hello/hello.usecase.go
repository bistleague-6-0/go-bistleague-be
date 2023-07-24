package hello

import (
	"bistleague-be/model/config"
	"bistleague-be/services/repository/hello"
	"github.com/gofiber/fiber/v2"
)

type Usecase struct {
	cfg       *config.Config
	HelloRepo *hello.Repository
}

func New(cfg *config.Config, helloRepo *hello.Repository) *Usecase {
	return &Usecase{
		cfg:       cfg,
		HelloRepo: helloRepo,
	}
}

func (u *Usecase) GetHello(
	ctx *fiber.Ctx,
) (string, error) {
	resp, err := u.HelloRepo.GetHello(ctx)
	if err != nil {
		return "", err
	}
	return resp, nil
}
