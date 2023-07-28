package team

import (
	"bistleague-be/model/config"
	"bistleague-be/model/entity"
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type Repository struct {
	cfg *config.Config
	db  *sqlx.DB
	qb  *goqu.DialectWrapper
}

func New(cfg *config.Config, db *sqlx.DB, qb *goqu.DialectWrapper) *Repository {
	return &Repository{
		cfg: cfg,
		db:  db,
		qb:  qb,
	}
}

func (r *Repository) CreateTeam(ctx context.Context, newTeam entity.TeamEntity) error {
	query := `INSERT INTO teams (team_name, team_leader_id, team_member_mails,bukti_pembayaran_url)
			  VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, query, newTeam.TeamName, newTeam.TeamLeaderID, newTeam.TeamMemberMails, newTeam.BuktiPembayaranURL)
	return err
}
