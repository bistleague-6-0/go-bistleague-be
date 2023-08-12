package entity

var VerificationStatusMap = map[int8]string{
	-1: "rejected",
	0:  "no file",
	1:  "under review",
	2:  "accepted",
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
	UserID             string `db:"uid"`
	Username           string `db:"username"`
	FullName           string `db:"full_name"`
	StudentCard        string `db:"student_card_filename"`
	StudentCardStatus  int8   `db:"student_card_status"`
	SelfPortrait       string `db:"self_portrait_filename"`
	SelfPortraitStatus int8   `db:"self_portrait_status"`
	Twibbon            string `db:"twibbon_filename"`
	TwibbonStatus      int8   `db:"twibbon_status"`
	Enrollment         string `db:"enrollment_filename"`
	EnrollmentStatus   int8   `db:"enrollment_status"`
}
