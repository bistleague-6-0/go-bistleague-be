package hello

import (
	"bistleague-be/model/config"
	"firebase.google.com/go/auth"
	"github.com/gofiber/fiber/v2"
)

type Repository struct {
	cfg        *config.Config
	authClient *auth.Client
}

func New(cfg *config.Config, client *auth.Client) *Repository {
	return &Repository{
		cfg:        cfg,
		authClient: client,
	}
}

func (r *Repository) GetHello(ctx *fiber.Ctx) (string, error) {
	return "hello, world!", nil
}
