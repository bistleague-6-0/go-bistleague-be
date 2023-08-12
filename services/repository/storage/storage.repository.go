package storage

import (
	"bistleague-be/model/config"
	"github.com/gofiber/fiber/v2"
)

type Repository struct {
	cfg      *config.Config
	Uploader *storageutils.ClientUploader
}

func New(cfg *config.Config, uploader *storageutils.ClientUploader) *Repository {
	return &Repository{
		cfg:      cfg,
		uploader: uploader,
	}
}
