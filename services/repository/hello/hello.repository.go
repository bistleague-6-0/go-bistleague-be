package hello

import (
	"bistleague-be/model/config"
	"github.com/gofiber/fiber/v2"
)

type Repository struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Repository {
	return &Repository{
		cfg: cfg,
	}
}

func (r *Repository) GetHello(ctx *fiber.Ctx) (string, error) {
	return "hello, world!", nil
}
