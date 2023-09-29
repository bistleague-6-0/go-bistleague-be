package challenge

const QueryInsertChallenge = `INSERT INTO 
    users_mini_challenge(uid, ig_username, ig_content_url, tiktok_username, tiktok_content_url) 
VALUES ($1, $2, $3, $4, $5)`
