package team

import (
	"bistleague-be/model/config"
	"bistleague-be/model/entity"
	"database/sql"
	"errors"
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

func (r *Repository) CreateTeam(ctx context.Context, newTeam entity.TeamEntity, redeemToken string) (string, error) {
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
	query := `INSERT INTO teams (team_name, team_leader_id, team_member_mails)
			  VALUES ($1, $2, $3) returning team_id`
	err = tx.GetContext(ctx, &teamID, query, newTeam.TeamName, newTeam.TeamLeaderID, newTeam.TeamMemberMails)
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

	// create team token
	q3 := "INSERT INTO teams_code(team_id, code) VALUES ($1, $2)"
	_, err = tx.ExecContext(ctx, q3, teamID.Id, redeemToken)
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
    u.uid, t.team_leader_id, u.username, u.full_name, t.team_id, t.team_name, t.payment_filename, t.verification_status,
    u.student_card_filename, u.student_card_status , u.self_portrait_filename, u.self_portrait_status , u.twibbon_filename, u.twibbon_status,
    u.enrollment_filename, u.enrollment_status
		from users u
			left join teams t
				on u.team_id = t.team_id
		where u.team_id = $1 LIMIT 3`
	resp := []entity.TeamWithUserEntity{}
	err := r.db.SelectContext(ctx, &resp, query, teamID)
	return resp, err
}

func (r *Repository) RedeemTeamCode(ctx context.Context, userID string, code string) (*entity.TeamRedeemCodeEntity, error) {
	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: 4,
		ReadOnly:  false,
	})
	if err != nil {
		return nil, err
	}

	// read code from db
	q1 := "SELECT team_id, used FROM teams_code WHERE code = $1 LIMIT 1"
	tc := entity.TeamRedeemCodeEntity{}

	// get token detail
	err = tx.GetContext(ctx, &tc, q1, code)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if tc.TeamID == "" || tc.Used < 1 {
		tx.Rollback()
		return nil, errors.New("code cannot be redeemed")
	}

	// assign team to user
	q2 := "UPDATE users SET team_id = $1 WHERE uid = $2"
	_, err = tx.ExecContext(ctx, q2, tc.TeamID, userID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// update token usage
	q3 := "UPDATE teams_code SET used = $1 WHERE team_id = $2"
	_, err = tx.ExecContext(ctx, q3, tc.Used-1, tc.TeamID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &tc, nil
}

func (r *Repository) InsertTeamDocument(ctx context.Context, filename string, teamID string) error {
	q := r.qb.Update("teams").Set(goqu.Record{
		"payment_filename":    filename,
		"verification_status": 1,
	}).Where(goqu.C("team_id").Eq(teamID))
	query, _, err := q.ToSQL()
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, query)
	return err
}
