package application

import (
	"bistleague-be/model/config"
	"bistleague-be/services/utils/storageutils"
	"cloud.google.com/go/storage"
	"context"
	_ "database/sql"
	"encoding/json"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/go-playground/validator/v10"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"google.golang.org/api/option"
	"regexp"
)

type CommonResource struct {
	Db       *sqlx.DB
	QBuilder *goqu.DialectWrapper
	Vld      *validator.Validate
	Uploader *storageutils.ClientUploader
}

func NewCommonResource(cfg *config.Config, ctx context.Context) (*CommonResource, error) {
	db, err := sqlx.Open(cfg.Database.DatabaseType, cfg.Database.Host)
	if err != nil {
		return nil, err
	}
	dialect := goqu.Dialect("postgres")
	vld := validator.New()
	vld.RegisterValidation("listOfMail", isListOfEmail)

	jsonCreds, err := json.Marshal(cfg.ServiceAccount)
	if err != nil {
		return nil, err
	}
	stg, err := storage.NewClient(ctx, option.WithCredentialsJSON(jsonCreds))
	if err != nil {
		fmt.Println(cfg.ServiceAccount)
		fmt.Println("error kontol", err)
		return nil, err
	}

	uploader := storageutils.ClientUploader{
		Cl:         stg,
		ProjectID:  cfg.Storage.BucketName,
		BucketName: cfg.Storage.ProjectID,
	}

	rsc := CommonResource{
		Db:       db,
		QBuilder: &dialect,
		Vld:      vld,
		Uploader: &uploader,
	}
	return &rsc, nil
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
