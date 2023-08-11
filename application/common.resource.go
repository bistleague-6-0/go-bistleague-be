package application

import (
	"bistleague-be/model/config"
	"context"
	_ "database/sql"
	"github.com/doug-martin/goqu/v9"
	"github.com/go-playground/validator/v10"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"regexp"
)

type CommonResource struct {
	Db       *sqlx.DB
	QBuilder *goqu.DialectWrapper
	Vld      *validator.Validate
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
	rsc := CommonResource{
		Db:       db,
		QBuilder: &dialect,
		Vld:      vld,
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
