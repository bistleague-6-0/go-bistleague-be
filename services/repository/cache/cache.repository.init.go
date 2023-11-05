package cache

import (
	"bistleague-be/model/config"
	"github.com/jellydator/ttlcache/v3"
)

type Repository struct {
	cfg   *config.Config
	cache *ttlcache.Cache[string, string]
}

func New(cfg *config.Config, cache *ttlcache.Cache[string, string]) (*Repository, error) {
	return &Repository{
		cache: cache,
	}, nil
}
