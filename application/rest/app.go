package main

import (
	"bistleague-be/application"
	"bistleague-be/model/config"
	"bistleague-be/model/dto"
	"bistleague-be/services/router/rest/auth"
	"bistleague-be/services/router/rest/hello"
	"bistleague-be/services/router/rest/profile"
	"bistleague-be/services/router/rest/team"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"net/http"
)

func applicationDelegate(cfg *config.Config) (*fiber.App, error) {
	ctx := context.Background()
	app := fiber.New(fiber.Config{
		AppName: fmt.Sprintf("%s %s", cfg.Server.Name, cfg.Stage),
	})
	// setup gzip
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	// setup cors
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

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

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(dto.NoBodyDTOResponseWrapper{
			Status:  http.StatusOK,
			Message: "Hello, Hacker!",
		})
	})
	//hello route
	helloRoute := hello.New(cfg, usecase.HelloUC)
	helloRoute.Register(app)

	//auth route
	authRoute := auth.New(cfg, usecase.AuthUC, resource.Vld)
	authRoute.RegisterRoute(app)

	//profile route
	profileRoute := profile.New(cfg, usecase.ProfileUC, resource.Vld)
	profileRoute.Register(app)

	//team route
	teamRoute := team.New(cfg, resource.Vld, usecase.TeamUC)
	teamRoute.Register(app)

	// admin group
	adminGroup := app.Group("/admin")
	adminGroup.Get("", func(ctx *fiber.Ctx) error {
		return ctx.JSON(dto.NoBodyDTOResponseWrapper{
			Status:  http.StatusOK,
			Message: "Fuck You Hacker!",
		})
	})

	return app, nil
}
