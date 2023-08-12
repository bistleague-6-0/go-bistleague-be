package storage

import (
	"bistleague-be/model/config"
	"bistleague-be/services/utils/storageutils"
)

type Repository struct {
	cfg      *config.Config
	uploader *storageutils.ClientUploader
}

func New(cfg *config.Config, uploader *storageutils.ClientUploader) *Repository {
	return &Repository{
		cfg:      cfg,
		uploader: uploader,
	}
}
