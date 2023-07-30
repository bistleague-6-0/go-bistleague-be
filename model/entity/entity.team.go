package entity

var VerificationStatusMap = map[int8]string{
	-1: "rejected",
	0:  "pending",
	1:  "accepted",
}

type TeamEntity struct {
	TeamID             string   `db:"team_id"`
	TeamName           string   `db:"team_name"`
	TeamLeaderID       string   `db:"team_leader_id"`
	BuktiPembayaranURL string   `db:"bukti_pembayaran_url"`
	IsVerified         bool     `db:"is_verified"`
	TeamMemberMails    []string `db:"team_member_mails"`
	IsActive           bool     `db:"is_active"`
	VerificationStatus int8     `db:"verification_status"`
}

type TeamRedeemCodeEntity struct {
	TeamID string `db:"team_id"`
	Code   string `db:"code"`
	Used   int8   `db:"used"`
}

type TeamWithUserEntity struct {
	TeamEntity
	UserID   string `db:"uid"`
	Username string `db:"username"`
	FullName string `db:"full_name"`
}
