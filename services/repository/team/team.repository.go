package team

import (
	"bistleague-be/model/config"
	"bistleague-be/model/entity"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
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
	q1 := `INSERT INTO teams (team_name, team_leader_id, team_member_mails)
			  VALUES ($1, $2, $3) returning team_id`
	err = tx.GetContext(ctx, &teamID, q1, newTeam.TeamName, newTeam.TeamLeaderID, newTeam.TeamMemberMails)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	fmt.Println(teamID)

	// create team token
	q3 := "INSERT INTO teams_code(team_id, code) VALUES ($1, $2)"
	_, err = tx.ExecContext(ctx, q3, teamID.Id, redeemToken)
	if err != nil {
		fmt.Println("err 54", err)
		tx.Rollback()
		return "", err
	}

	// create team docs table
	q1two := `INSERT INTO teams_docs(team_id) VALUES ($1)`
	_, err = tx.ExecContext(ctx, q1two, teamID.Id)
	if err != nil {
		fmt.Println("err 63", err)
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
		t.team_id, t.team_name, t.team_leader_id, tc.code,
		td.payment_filename, td.payment_url, td.payment_status, td.payment_rejection,
		u.uid, u.username, u.full_name,
		ud.student_card_filename, ud.student_card_url, ud.student_card_status, ud.student_card_rejection,
		ud.enrollment_filename, ud.enrollment_url, ud.enrollment_status, ud.enrollment_rejection,
		ud.self_portrait_filename, ud.self_portrait_url, ud.self_portrait_url, ud.self_portrait_rejection,
		ud.twibbon_filename, ud.twibbon_url, ud.twibbon_status, ud.twibbon_rejection,
		ud.is_doc_verified, u.is_profile_verified
	from users u
		left join users_docs ud on ud.uid = u.uid
		left join teams t on u.team_id = t.team_id
		left join teams_code tc on tc.team_id = t.team_id
		left join teams_docs td on td.team_id = t.team_id
	where u.team_id = $1 LIMIT 3
	`
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

func (r *Repository) InsertTeamDocument(ctx context.Context, filename string, fileURL string, teamID string) error {
	q := r.qb.Update("teams_docs").Set(goqu.Record{
		"payment_filename": filename,
		"payment_url":      fileURL,
		"payment_status":   1,
	}).Where(goqu.C("team_id").Eq(teamID))
	query, _, err := q.ToSQL()
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, query)
	return err
}

func (r *Repository) InsertTeamSubmission(ctx context.Context, filename string, fileURL string, teamID string, docType string) error {
	wib, _ := time.LoadLocation("Asia/Jakarta")
	currentTime := time.Now()
	currentTimeInWIB := currentTime.In(wib)

	var colFilename, colURL, colLastUpdate string

	if docType == "submission_1" {
		colFilename = "submission_1_filename"
		colURL = "submission_1_url"
		colLastUpdate = "submission_1_lastupdate"
	} else {
		colFilename = "submission_2_filename"
		colURL = "submission_2_url"
		colLastUpdate = "submission_2_lastupdate"
	}

	q := r.qb.Update("teams_docs").Set(goqu.Record{
		colFilename:   filename,
		colURL:        fileURL,
		colLastUpdate: pq.FormatTimestamp(currentTimeInWIB),
	}).Where(goqu.C("team_id").Eq(teamID))

	query, _, err := q.ToSQL()
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, query)
	return err
}

func (r *Repository) GetSubmission(ctx context.Context, teamID string) (*entity.TeamSubmission, error) {
	query := `
        SELECT
            team_id, submission_1_filename, submission_1_url, submission_1_lastupdate,
            submission_2_filename, submission_2_url, submission_2_lastupdate
        FROM teams_docs
        WHERE team_id = $1
        LIMIT 1
    `

	resp := entity.TeamSubmission{}
	err := r.db.GetContext(ctx, &resp, query, teamID)

	if err != nil {
		println("Error fetching submission:", err.Error())
	}

	return &resp, err
}

func (r *Repository) GetTeamCount(ctx context.Context) (int, error) {
	q := `SELECT COUNT(*) FROM teams`
	var count int
	err := r.db.GetContext(ctx, &count, q)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *Repository) GetPayments(ctx context.Context, page int, pageSize int) ([]entity.TeamPayment, error) {
	query := `
        SELECT
            t.team_id, t.team_name, t.team_member_mails, td.payment_filename, td.payment_url,
            td.payment_status, tc.code
        FROM teams t
		LEFT JOIN teams_docs td
		ON t.team_id = td.team_id
		LEFT JOIN teams_code tc
		ON t.team_id = tc.team_id
		ORDER BY t.team_name
		LIMIT $1 OFFSET $2
    `

	resp := []entity.TeamPayment{}
	offset := (page - 1) * pageSize
	err := r.db.SelectContext(ctx, &resp, query, pageSize, offset)

	return resp, err
}

func (r *Repository) UpdatePaymentStatus(ctx context.Context, teamID string, status int, rejection string) error {
	q := r.qb.Update("teams_docs").Set(goqu.Record{
		"payment_status":    1,
		"payment_rejection": rejection,
	}).Where(goqu.C("team_id").Eq(teamID))
	query, _, err := q.ToSQL()
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, query)
	return err
}
