package application

import (
	"bistleague-be/model/config"
	"context"
)

type CommonResource struct {
}

func NewCommonResource(cfg *config.Config, ctx context.Context) (*CommonResource, error) {
	rsc := CommonResource{}
	return &rsc, nil
}
