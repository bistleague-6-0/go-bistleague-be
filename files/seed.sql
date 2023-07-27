CREATE TABLE IF NOT EXISTS users (
    uid uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    username varchar(55) unique  not null ,
    email VARCHAR(75) unique not null,
    password VARCHAR(155) not null,
    full_name VARCHAR(155) not null,
    institution VARCHAR(155),
    major VARCHAR(55),
    entry_year int2 default 0,
    linkedin_url VARCHAR(100),
    line_id VARCHAR(55),
    team_id uuid,
    inserted_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp
);

CREATE TABLE IF NOT EXISTS teams(
    team_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    team_name VARCHAR(55) UNIQUE NOT NULL,
    bukti_pembayaran_url VARCHAR(155) NOT NULL
);

CREATE TABLE IF NOT EXISTS teams_member_email(
    team_id uuid PRIMARY KEY,
    email VARCHAR(155) not null
);

CREATE TABLE IF NOT EXISTS teams_code(
    team_id uuid PRIMARY KEY,
    code VARCHAR(55) UNIQUE,
    used int4 DEFAULT 3
)