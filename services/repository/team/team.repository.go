package team

import (
	"bistleague-be/model/config"
	"bistleague-be/model/entity"
	"database/sql"
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

func (r *Repository) CreateTeam(ctx context.Context, newTeam entity.TeamEntity) (string, error) {
	var teamID struct {
		Id string `db:"team_id"`
	}
	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: 4,
		ReadOnly:  false,
	})
	if err != nil {
		return "", err
	}
	// create team
	query := `INSERT INTO teams (team_name, team_leader_id, team_member_mails,bukti_pembayaran_url)
			  VALUES ($1, $2, $3, $4) returning team_id`
	err = tx.GetContext(ctx, &teamID, query, newTeam.TeamName, newTeam.TeamLeaderID, newTeam.TeamMemberMails, newTeam.BuktiPembayaranURL)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	// update user's team id
	q2 := "UPDATE users SET team_id = $1 WHERE uid = $2"
	_, err = tx.ExecContext(ctx, q2, teamID.Id, newTeam.TeamLeaderID)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	err = tx.Commit()
	if err != nil {
		return "", err
	}
	return teamID.Id, nil
}

func (r *Repository) GetTeamInformation(ctx context.Context, teamID string) ([]entity.TeamWithUserEntity, error) {
	query := `select 
    u.uid, t.team_leader_id, u.username, u.full_name, t.team_id, t.team_name, t.bukti_pembayaran_url, t.verification_status
		from users u
			left join teams t
				on u.team_id = t.team_id
		where u.team_id = $1 LIMIT 3`
	resp := []entity.TeamWithUserEntity{}
	err := r.db.SelectContext(ctx, &resp, query, teamID)
	return resp, err
}
