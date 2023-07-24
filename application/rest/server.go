package main

import (
	"bistleague-be/model/config"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func startServer(cfg *config.Config, app *fiber.App) error {
	return app.Listen(fmt.Sprintf(":%s", cfg.Server.Port))
}
