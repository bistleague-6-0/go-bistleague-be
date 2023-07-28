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

func (r *Repository) GetTeamInformation(ctx context.Context, teamID string) ([]entity.TeamWithUserEntity, error) {
	query := `select 
    u.uid, t.team_leader_id, u.username, u.full_name, t.team_id, t.team_name, t.bukti_pembayaran_url, t.is_verified
		from users u
			left join teams t
				on u.team_id = t.team_id
		where u.team_id = $1 LIMIT 3`
	resp := []entity.TeamWithUserEntity{}
	err := r.db.SelectContext(ctx, &resp, query, teamID)
	return resp, err
}
