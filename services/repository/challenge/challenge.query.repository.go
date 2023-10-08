package challenge

const (
	queryInsertChallenge = `INSERT INTO 
    users_mini_challenge(uid, ig_username, ig_content_url, tiktok_username, tiktok_content_url) 
VALUES ($1, $2, $3, $4, $5)`
	queryUpdateChallenge = `UPDATE 
    users_mini_challenge 
	SET ig_username=$1, ig_content_url=$2, tiktok_username=$3, tiktok_content_url=$4
	WHERE uid = $5`
	queryGetChallenge = `SELECT 
    ig_username, ig_content_url, tiktok_username, tiktok_content_url 
	FROM users_mini_challenge
	WHERE uid =$1`
)
