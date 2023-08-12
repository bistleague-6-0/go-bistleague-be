package storage

import (
	"bistleague-be/model/config"
	"bistleague-be/services/utils/storageutils"
	"cloud.google.com/go/storage"
	"context"
	"fmt"
)

type Repository struct {
	cfg    *config.Config
	bucket *storage.BucketHandle
}

func New(cfg *config.Config, bucket *storage.BucketHandle) *Repository {
	return &Repository{
		cfg:    cfg,
		bucket: bucket,
	}
}

func (r *Repository) UploadDocument(ctx context.Context, file *storageutils.Base64File) (string, error) {
	filename := fmt.Sprintf("%s%s", file.Name, file.Ext)
	wc := r.bucket.Object(filename).NewWriter(ctx)
	if _, err := wc.Write(file.Contents); err != nil {
		return "", fmt.Errorf("Write: %v", err)
	}
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %v", err)
	}
	fileURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s", r.cfg.Storage.BucketName, filename)
	return fileURL, nil
}
