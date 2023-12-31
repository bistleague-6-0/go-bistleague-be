package application

import (
	"bistleague-be/model/config"
	"cloud.google.com/go/storage"
	"github.com/gofiber/fiber/v2/log"
	"context"
	_ "database/sql"
	"encoding/json"
	"github.com/doug-martin/goqu/v9"
	"github.com/go-playground/validator/v10"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jellydator/ttlcache/v3"
	"github.com/jmoiron/sqlx"
	"google.golang.org/api/option"
	"regexp"
	"time"
)

type CommonResource struct {
	Db       *sqlx.DB
	QBuilder *goqu.DialectWrapper
	Vld      *validator.Validate
	bucket   *storage.BucketHandle
	cache    *ttlcache.Cache[string, string]
}

func NewCommonResource(cfg *config.Config, ctx context.Context) (*CommonResource, error) {
	db, err := sqlx.Open(cfg.Database.DatabaseType, cfg.Database.Host)
	if err != nil {
		return nil, err
	}
	dialect := goqu.Dialect("postgres")
	vld := validator.New()
	vld.RegisterValidation("listOfMail", isListOfEmail)
	vld.RegisterValidation("isTeamDoc", isTeamDocs)

	jsonCreds, err := json.Marshal(cfg.ServiceAccount)
	if err != nil {
		return nil, err
	}
	storageCli, err := storage.NewClient(ctx, option.WithCredentialsJSON(jsonCreds))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	bucket := storageCli.Bucket(cfg.Storage.BucketName)
	cache := ttlcache.New[string, string](
		ttlcache.WithTTL[string, string](12 * time.Hour),
	)
	go cache.Start()
	go func() {
		for {
			time.Sleep(4 * time.Hour)
			cache.DeleteExpired()
		}
	}()
	rsc := CommonResource{
		Db:       db,
		QBuilder: &dialect,
		Vld:      vld,
		bucket:   bucket,
		cache:    cache,
	}
	return &rsc, nil
}

func isTeamDocs(fl validator.FieldLevel) bool {
	doctypes := map[string]int8{
		"payment":       0,
		"student_card":  1,
		"self_portrait": 2,
		"twibbon":       3,
		"enrollment":    4,
		"submission_1":  5,
		"submission_2":  6,
	}
	input, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	_, ok = doctypes[input]
	return ok
}

func isListOfEmail(fl validator.FieldLevel) bool {
	input, ok := fl.Field().Interface().([]string)
	if !ok {
		return false
	}
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(emailRegex)
	for _, mail := range input {
		if !regex.MatchString(mail) {
			return false
		}
	}
	return true
}
