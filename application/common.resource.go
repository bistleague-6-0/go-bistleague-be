package application

import (
	"bistleague-be/model/config"
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type CommonResource struct {
	AuthClient *auth.Client
}

func NewCommonResource(cfg *config.Config, ctx context.Context) (*CommonResource, error) {
	rsc := CommonResource{}
	fbase, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: cfg.Firebase.ProjectID,
	}, option.WithAPIKey(cfg.Firebase.ApiKey))
	if err != nil {
		return nil, err
	}
	authClient, err := fbase.Auth(ctx)
	if err != nil {
		return nil, err
	}
	rsc.AuthClient = authClient
	return &rsc, nil
}
