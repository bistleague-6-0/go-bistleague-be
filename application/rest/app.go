package main

import (
	"bistleague-be/application"
	"bistleague-be/model/config"
	"bistleague-be/services/router/rest/auth"
	"bistleague-be/services/router/rest/hello"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func applicationDelegate(cfg *config.Config) (*fiber.App, error) {
	ctx := context.Background()
	app := fiber.New(fiber.Config{
		AppName: fmt.Sprintf("%s %s", cfg.Server.Name, cfg.Stage),
	})
	resource, err := application.NewCommonResource(cfg, ctx)
	if err != nil {
		return nil, err
	}
	repository, err := application.NewCommonRepository(cfg, resource)
	if err != nil {
		return nil, err
	}
	usecase, err := application.NewCommonUsecase(cfg, repository)
	if err != nil {
		return nil, err
	}

	//hello router
	helloRoute := hello.New(cfg, usecase.HelloUC)
	helloRoute.Register(app, resource.AuthClient)

	//auth route [Development only]
	if cfg.Stage == "staging" {
		authRoute := auth.New(cfg)
		authRoute.RegisterRoute(app)
	}
	return app, nil
}
