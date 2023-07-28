package entity

type TeamEntity struct {
	TeamID             string   `db:"team_id"`
	TeamName           string   `db:"team_name"`
	TeamLeaderID       string   `db:"team_leader_id"`
	BuktiPembayaranURL string   `db:"bukti_pembayaran_url"`
	IsVerified         bool     `db:"is_verified"`
	TeamMemberMails    []string `db:"team_member_mails"`
	IsActive           bool     `db:"is_active"`
}

type TeamWithUserEntity struct {
	TeamEntity
	UserID   string `db:"uid"`
	Username string `db:"username"`
	FullName string `db:"full_name"`
}
