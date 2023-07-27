CREATE TABLE IF NOT EXISTS users (
    uid uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(55) unique not null,
    password VARCHAR(155) not null,
    email VARCHAR(75) unique not null,
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