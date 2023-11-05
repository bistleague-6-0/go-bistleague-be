package cache

import (
	"context"
	"errors"
	"time"
)

func (r *Repository) Check(ctx context.Context, key string) bool {
	return r.cache.Has(key)
}

func (r *Repository) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	r.cache.Set(key, value, ttl)
	return nil
}

func (r *Repository) Get(ctx context.Context, key string) (string, error) {
	isAvail := r.cache.Has(key)
	if !isAvail {
		return "", errors.New("item not found")
	}
	item := r.cache.Get(key)
	if item.IsExpired() {
		return "", errors.New("item is expired")
	}
	value := item.Value()
	if value == "" {
		return "", errors.New("item is empty")
	}
	return value, nil
}

func (r *Repository) Delete(ctx context.Context, key string) {
	r.cache.Delete(key)
}
