package application

import (
	"bistleague-be/model/config"
	"context"
	_ "database/sql"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type CommonResource struct {
	Db       *sqlx.DB
	QBuilder *goqu.DialectWrapper
}

func NewCommonResource(cfg *config.Config, ctx context.Context) (*CommonResource, error) {
	db, err := sqlx.Open(cfg.Database.DatabaseType, cfg.Database.Host)
	if err != nil {
		return nil, err
	}
	dialect := goqu.Dialect("postgres")
	rsc := CommonResource{
		Db:       db,
		QBuilder: &dialect,
	}
	return &rsc, nil
}
