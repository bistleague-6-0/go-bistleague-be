package document

import (
	"bistleague-be/model/config"
	"context"
)

type Repository struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Repository {
	return &Repository{
		cfg: cfg,
	}
}

func (r *Repository) UploadTeamDocument(ctx context.Context, doctype string, filename string, teamID string) error {
	return nil
}
