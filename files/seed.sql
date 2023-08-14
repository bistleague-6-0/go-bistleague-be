CREATE TABLE IF NOT EXISTS users (
    uid uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    username varchar(55) unique  not null ,
    email VARCHAR(75) unique not null,
    password VARCHAR(155) not null,
    full_name VARCHAR(155) not null,
    user_age int8 default 0,
    phone_number VARCHAR(155),
    institution VARCHAR(155),
    address TEXT DEFAULT '',
    major VARCHAR(55),
    entry_year int2 default 0,
    linkedin_url VARCHAR(100),
    line_id VARCHAR(55),
    team_id uuid,
    student_card_filename VARCHAR(155) DEFAULT '',
    student_card_status INT DEFAULT 0,
    self_portrait_filename VARCHAR(155) DEFAULT '',
    self_portrait_status INT DEFAULT 0,
    twibbon_filename VARCHAR(155) DEFAULT '',
    twibbon_status INT DEFAULT 0,
    enrollment_filename VARCHAR(155) DEFAULT '',
    enrollment_status INT DEFAULT 0,
    is_doc_verified boolean default false,
    is_profile_verified boolean default false,
    inserted_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp
);

CREATE TABLE IF NOT EXISTS teams(
    team_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    team_name VARCHAR(55) UNIQUE NOT NULL,
    team_leader_id UUID unique not null ,
    payment_filename VARCHAR(155) DEFAULT '',
    verification_status INT DEFAULT 0,
    team_member_mails TEXT[],
    is_active boolean default true
);

CREATE TABLE IF NOT EXISTS teams_code(
    team_id uuid PRIMARY KEY,
    code VARCHAR(55) UNIQUE,
    team_member_mails TEXT[],
    used int8 default 2
)